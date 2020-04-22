"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

from distutils.util import strtobool

import sys
from fabric.api import cd, env, execute, local, run, settings
from fabric.operations import get

sys.path.append('../../orc8r')
import tools.fab.dev_utils as dev_utils
import tools.fab.pkg as pkg
from tools.fab.hosts import ansible_setup, split_hoststring, vagrant_setup
from tools.fab.vagrant import setup_env_vagrant

"""
Magma Gateway packaging tool:

Magma packages released to different channels have different version schemes.

    - dev: used for development.

    - test: used for Continuous Integration (CI). Packages in the `test`
            channel should be built and released automatically.

# HOWTO build magma.deb

1. From your laptop, update the magma version number in `release/build-magma.sh`

2. From the dev VM, build the magma package. Dependency packages are recorded
in `release/magma.lockfile`

    fab dev package
    # optionally upload to aws (if you are configured for it)
    fab dev package upload_to_aws
"""

AGW_ROOT = "$MAGMA_ROOT/lte/gateway"
AGW_PYTHON_ROOT = "$MAGMA_ROOT/lte/gateway/python"
ORC8R_AGW_PYTHON_ROOT = "$MAGMA_ROOT/orc8r/gateway/python"
AGW_INTEG_ROOT = "$MAGMA_ROOT/lte/gateway/python/integ_tests"
DEFAULT_CERT = "$MAGMA_ROOT/.cache/test_certs/rootCA.pem"
DEFAULT_PROXY = "$MAGMA_ROOT/lte/gateway/configs/control_proxy.yml"

# Look for keys as specified in our ~/.ssh/config
env.use_ssh_config = True
# Disable ssh known hosts to resolve key errors
# with multiple vagrant boxes in use.
env.disable_known_hosts = True


def dev():
    env.debug_mode = True


def test():
    env.debug_mode = False


def package(vcs='hg', all_deps="False",
            cert_file=DEFAULT_CERT, proxy_config=DEFAULT_PROXY,
            destroy_vm='False'):
    """ Builds the magma package """
    all_deps = False if all_deps == "False" else True
    destroy_vm = bool(strtobool(destroy_vm))

    # If a host list isn't specified, default to the magma vagrant vm
    if not env.hosts:
        vagrant_setup('magma', destroy_vm=destroy_vm)

    if not hasattr(env, 'debug_mode'):
        print("Error: The Deploy target isn't specified. Specify one with\n\n"
              "\tfab [dev|test] package")
        exit(1)

    hash = pkg.get_commit_hash(vcs)

    with cd('~/magma/lte/gateway'):
        print("Building magma package, picking up commit %s..." % hash)
        run('make clean')
        build_type = "Debug" if env.debug_mode else "RelWithDebInfo"
        run('./release/build-magma.sh -h "%s" -t %s --cert %s --proxy %s' %
            (hash, build_type, cert_file, proxy_config))

        # Generate magma dependency packages
        run('mkdir -p ~/magma-deps')
        print("Generating lte/setup.py magma dependency packages")
        run('./release/pydep finddep -b --build-output ~/magma-deps python/setup.py')

        print("Generating orc8r/setup.py magma dependency packages")
        run('./release/pydep finddep -b --build-output ~/magma-deps %s/setup.py' % ORC8R_AGW_PYTHON_ROOT)

        run('rm -rf ~/magma-packages')
        run('mkdir -p ~/magma-packages')
        try:
            run('cp -f ~/magma-deps/*.deb ~/magma-packages')
        except Exception:
            # might be a problem if no deps found, but don't handle here
            pass
        run('mv *.deb ~/magma-packages')

        with cd('release'):
            run('cat mirrored_packages | '
                'xargs -I% sudo aptitude download -q2 %')
            run('cp *.deb ~/magma-packages')
            run('sudo rm -f *.deb')

        if all_deps:
            pkg.download_all_pkgs()
            run('cp /var/cache/apt/archives/*.deb ~/magma-packages')


def depclean():
    '''Remove all generated packaged for dependencies'''
    # If a host list isn't specified, default to the magma vagrant vm
    if not env.hosts:
        setup_env_vagrant()
    run('rm -rf ~/magma-deps')


def upload_to_aws():
    # If a host list isn't specified, default to the magma vagrant vm
    if not env.hosts:
        setup_env_vagrant()

    pkg.upload_pkgs_to_aws()


def copy_packages():
    if not env.hosts:
        setup_env_vagrant()
    pkg.copy_packages()


def connect_gateway_to_cloud(control_proxy_setting_path=None,
                             cert_path=DEFAULT_CERT):
    """
    Setup the gateway VM to connects to the cloud
    Path to control_proxy.yml and rootCA.pem could be specified to use
    non-default control proxy setting and certificates
    """
    setup_env_vagrant()
    dev_utils.connect_gateway_to_cloud(control_proxy_setting_path, cert_path)


def s1ap_setup_cloud():
    """ Prepare VMs for s1ap tests touching the cloud. """
    # Use the local cloud for integ tests
    connect_gateway_to_cloud()

    # Update the gateway's streamer timeout and restart services
    run("sudo mkdir -p /var/opt/magma/configs")
    _set_service_config_var('streamer', 'reconnect_sec', 3)

    # Update the gateway's metricsd collect/sync intervals
    _set_service_config_var('metricsd', 'collect_interval', 5)
    _set_service_config_var('metricsd', 'sync_interval', 5)

    run("sudo systemctl stop magma@*")
    run("sudo systemctl restart magma@magmad")


def integ_test(gateway_host=None, test_host=None, trf_host=None,
               destroy_vm="True"):
    """
    Run the integration tests. This defaults to running on local vagrant
    machines, but can also be pointed to an arbitrary host (e.g. amazon) by
    passing "address:port" as arguments

    gateway_host: The ssh address string of the machine to run the gateway
        services on. Formatted as "host:port". If not specified, defaults to
        the `magma` vagrant box.

    test_host: The ssh address string of the machine to run the tests on
        on. Formatted as "host:port". If not specified, defaults to the
        `magma_test` vagrant box.

    trf_host: The ssh address string of the machine to run the TrafficServer
        on. Formatted as "host:port". If not specified, defaults to the
        `magma_trfserver` vagrant box.
    """

    destroy_vm = bool(strtobool(destroy_vm))

    # Setup the gateway: use the provided gateway if given, else default to the
    # vagrant machine
    gateway_ip = '192.168.60.142'
    if not gateway_host:
        gateway_host = vagrant_setup("magma", destroy_vm)
    else:
        ansible_setup(gateway_host, "dev", "magma_dev.yml")
        gateway_ip = gateway_host.split('@')[1].split(':')[0]

    execute(_dist_upgrade)
    execute(_build_magma)
    execute(_run_unit_tests)
    execute(_python_coverage)
    execute(_start_gateway)

    # Run suite of integ tests that are required to be run on the access gateway
    # instead of the test VM
    # TODO: fix the integration test T38069907
    # execute(_run_local_integ_tests)

    # Setup the trfserver: use the provided trfserver if given, else default to the
    # vagrant machine
    if not trf_host:
        trf_host = vagrant_setup("magma_trfserver", destroy_vm)
    else:
        ansible_setup(trf_host, "trfserver", "magma_trfserver.yml")
    execute(_start_trfserver)

    # Run the tests: use the provided test machine if given, else default to
    # the vagrant machine
    if not test_host:
        test_host = vagrant_setup("magma_test", destroy_vm)
    else:
        ansible_setup(test_host, "test", "magma_test.yml")

    execute(_make_integ_tests)
    execute(_run_integ_tests, gateway_ip)

    if not gateway_host:
        setup_env_vagrant()
    else:
        env.hosts = [gateway_host]
    execute(_oai_coverage)


def get_test_logs(gateway_host=None, test_host=None, trf_host=None):
    '''
    Downloads the relevant magma logs from the given gateway and test machines.
    Places the logs in /tmp/build_logs.zip

    gateway_host: The ssh address string of the gateway machine formatted as
        "host:port". If not specified, defaults to the `magma` vagrant box.

    test_host: The ssh address string of the test machine formatted as
        "host:port". If not specified, defaults to the `magma_test` vagrant
        box.

    trf_host: The ssh address string of the machine to run the TrafficServer
        on. Formatted as "host:port". If not specified, defaults to the
        `magma_trfserver` vagrant box.

    '''
    # Grab the build logs from the machines and bring them to the host
    local('rm -rf /tmp/build_logs')
    local('mkdir /tmp/build_logs')
    local('mkdir /tmp/build_logs/dev')
    local('mkdir /tmp/build_logs/test')
    local('mkdir /tmp/build_logs/trfserver')
    dev_files = ['/var/log/mme.log', '/var/log/syslog']
    test_files = ['/var/log/syslog', '/tmp/fw/']
    trf_files = ['/home/admin/nohup.out']

    # Set up to enter the gateway host
    env.host_string = gateway_host
    if not gateway_host:
        setup_env_vagrant("magma")
        gateway_host = env.hosts[0]
    (env.user, _, _) = split_hoststring(gateway_host)

    # Don't fail if the logs don't exists
    for p in dev_files:
        with settings(warn_only=True):
            get(remote_path=p, local_path='/tmp/build_logs/dev/',
                use_sudo=True)

    # Set up to enter the trfserver host
    env.host_string = trf_host
    if not trf_host:
        setup_env_vagrant("magma_trfserver")
        trf_host = env.hosts[0]
    (env.user, _, _) = split_hoststring(trf_host)

    # Don't fail if the logs don't exists
    for p in trf_files:
        with settings(warn_only=True):
            get(remote_path=p, local_path='/tmp/build_logs/trfserver/',
                use_sudo=True)

    # Set up to enter the test host
    env.host_string = test_host
    if not test_host:
        setup_env_vagrant("magma_test")
        test_host = env.hosts[0]
    (env.user, _, _) = split_hoststring(test_host)

    # Fix the permissions on the fw directory -- it has permissions 000
    # otherwise
    with settings(warn_only=True):
        run('sudo chmod 755 /tmp/fw')

    # Don't fail if the logs don't exists
    for p in test_files:
        with settings(warn_only=True):
            get(remote_path=p, local_path='/tmp/build_logs/test/',
                use_sudo=True)

    local('rm -rf /tmp/build_logs.zip')
    local('zip -r /tmp/build_logs.zip /tmp/build_logs')
    local('rm -rf /tmp/build_logs')


def _dist_upgrade():
    """ Upgrades OS packages on dev box """
    run('sudo apt-get update')
    run('sudo apt-get -y dist-upgrade')


def _build_magma():
    """ Builds magma """

    with cd(AGW_ROOT):
        run('make')


def _oai_coverage():
    """ Get the code coverage statistic for OAI """

    with cd(AGW_ROOT):
        run('make coverage_oai')


def _run_unit_tests():
    """ Run the magma unit tests """
    with cd(AGW_ROOT):
        # Run the unit tests
        run('make test')


def _python_coverage():
    with cd(AGW_PYTHON_ROOT):
        run('make coverage')


def _start_gateway():
    """ Starts the gateway """

    with cd(AGW_ROOT):
        run('make run')


def _run_local_integ_tests():
    """ Execute integ tests that run on magma access gateway """
    with cd(AGW_INTEG_ROOT):
        run('make local_integ_test')


def _set_service_config_var(service, var_name, value):
    """ Sets variable in config file by value """
    run("echo '%s: %s' | sudo tee -a /var/opt/magma/configs/%s.yml" % (
        var_name, str(value), service))


def _start_trfserver():
    """ Starts the traffic gen server"""
    # disable-tcp-checksumming
    # trfgen-server daemon
    host = env.hosts[0].split(':')[0]
    port = env.hosts[0].split(':')[1]
    key = env.key_filename
    local('ssh -i %s -o UserKnownHostsFile=/dev/null'
          ' -o StrictHostKeyChecking=no -tt %s -p %s'
          ' \'sudo ethtool --offload eth1 rx off tx off; sudo ethtool --offload eth2 rx off tx off;'
          ' nohup sudo /usr/local/bin/traffic_server.py -d 192.168.60.144 62462\''
          % (key, host, port))


def _make_integ_tests():
    """ Build and run the integration tests """

    with cd(AGW_PYTHON_ROOT):
        run('make')
    with cd(AGW_INTEG_ROOT):
        run('make')


def _run_integ_tests(gateway_ip='192.168.60.142'):
    """ Run the integration tests

    For now, just run a single basic test
    """

    host = env.hosts[0].split(':')[0]
    port = env.hosts[0].split(':')[1]
    key = env.key_filename
    """
    NOTE: the s1aptester produces a bunch of output which the python ssh
    library, and thus fab, has trouble processing quickly. Instead, we manually
    ssh into machine and run the tests

    ssh switch reference:
        -i: identity file
        -tt: (really) force a psuedo tty -- The tests can't initialize logging
            without this
        -p: the port to connect to
        -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no: have ssh
         never prompt to confirm the host fingerprints
    """
    local('ssh -i %s -o UserKnownHostsFile=/dev/null'
          ' -o StrictHostKeyChecking=no -tt %s -p %s'
          ' \'cd $MAGMA_ROOT/lte/gateway/python/integ_tests; '
          # We don't have a proper shell, so the `magtivate` alias isn't
          # available. We instead directly source the activate file
          ' sudo ethtool --offload eth1 rx off tx off; sudo ethtool --offload eth2 rx off tx off;'
          ' source ~/build/python/bin/activate;'
          ' export GATEWAY_IP=%s;'
          ' make integ_test\''
          % (key, host, port, gateway_ip))
