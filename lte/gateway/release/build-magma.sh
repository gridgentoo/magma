#!/bin/bash
#
# Copyright (c) 2016-present, Facebook, Inc.
# All rights reserved.
#
# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree. An additional grant
# of patent rights can be found in the PATENTS file in the same directory.

# This script builds Magma based on the current state of your repo. It needs to
# be run inside the VM.

set -e
shopt -s extglob

# Please update the version number accordingly for beta/stable builds
# Test builds are versioned automatically by fabfile.py
VERSION=1.0.2 # magma version number
SCTPD_MIN_VERSION=1.0.2 # earliest version of sctpd with which this version is compatible

# RelWithDebInfo or Debug
BUILD_TYPE=Debug

# Cmdline options that overwrite the version configs above
COMMIT_HASH=""  # hash of top magma commit (hg log $MAGMA_PATH)
CERT_FILE="$MAGMA_ROOT/.cache/test_certs/rootCA.pem"
CONTROL_PROXY_FILE="$MAGMA_ROOT/lte/gateway/configs/control_proxy.yml"

while [[ $# -gt 0 ]]
do
key="$1"
case $key in
    -v|--version)
    VERSION="$2"
    shift  # pass argument or value
    ;;
    -h|--hash)
    COMMIT_HASH="$2"
    shift
    ;;
    -t|--type)
    BUILD_TYPE="$2"
    shift  # pass argument or value
    ;;
    -c|--cert)
    CERT_FILE="$2"
    shift
    ;;
    -p|--proxy)
    CONTROL_PROXY_FILE="$2"
    shift
    ;;
    *)
    echo "Error: unknown cmdline option:" $key
    echo "Usage: $0 [-v|--version V] [-i|--iteration I] [-h|--hash HASH]
    [-t|--type Debug|RelWithDebInfo] [-c|--cert <path to cert .pem file>]
    [-p|--proxy <path to control_proxy config .yml file]>"
    exit 1
    ;;
esac
shift  # past argument or value
done

case $BUILD_TYPE in
    Debug)
    ;;
    RelWithDebInfo)
    ;;
    *)
    echo "Error: unknown type option:" $BUILD_TYPE
    echo "Usage: [-t|--type Debug|RelWithDebInfo]"
    exit 1
    ;;
esac


# Default options
BUILD_DATE=`date -u +"%Y%m%d%H%M%S"`
ARCH=amd64
PKGFMT=deb
PKGNAME=magma
SCTPD_PKGNAME=magma-sctpd

# Magma system dependencies: anything that we depend on at the top level, add
# here.
MAGMA_DEPS=(
    "grpc-dev >= 1.15.0"
    "libprotobuf10 >= 3.0.0"
    "lighttpd >= 1.4.45"
    "libxslt1.1"
    "nghttp2-proxy >= 1.18.1"
    "python3-protobuf >= 3.0.0"
    "redis-server >= 3.2.0"
    "sudo"
    "dnsmasq >= 2.72"
    "net-tools" # for ifconfig
    "python3-pip"
    "python3-apt" # The version in pypi is abandoned and broken on stretch
    "python3-aioeventlet" # The version in pypi got deleted
    "libsystemd-dev"
    "libyaml-cpp-dev" # install yaml parser
    "libgoogle-glog-dev"
    "nlohmann-json-dev" # c++ json parser
    "python-redis"
    "magma-cpp-redis"
    "libfolly-dev" # required for C++ services
    "libdouble-conversion-dev" # required for folly
    "libboost-chrono-dev" # required for folly
    "td-agent-bit >= 1.3.2" # fluent-bit
    "ntpdate" # required for eventd time synchronization
    )

# OAI runtime dependencies
OAI_DEPS=(
    "libasan3"
    "libconfig9"
    "oai-asn1c >= 0~20160721+c3~r43c4a295"
    "oai-freediameter >= 1.2.0-1"
    "oai-gnutls >= 3.1.23"
    "oai-nettle >= 1.0.1"
    "prometheus-cpp-dev >= 1.0.2"
    "liblfds710"
    "magma-sctpd >= ${SCTPD_MIN_VERSION}"
    )

# OVS runtime dependencies
OVS_DEPS=(
    "magma-libfluid >= 0.1.0.4"
    )

# generate string for FPM
SYSTEM_DEPS=""
for dep in "${MAGMA_DEPS[@]}"
do
    SYSTEM_DEPS=${SYSTEM_DEPS}" -d '"${dep}"'"
done
for dep in "${OAI_DEPS[@]}"
do
    SYSTEM_DEPS=${SYSTEM_DEPS}" -d '"${dep}"'"
done
for dep in "${OVS_DEPS[@]}"
do
    SYSTEM_DEPS=${SYSTEM_DEPS}" -d '"${dep}"'"
done

RELEASE_DIR=${MAGMA_ROOT}/lte/gateway/release
POSTINST=${RELEASE_DIR}/magma-postinst

# python environment
PY_VERSION=python3.5
PY_PKG_LOC=dist-packages
PY_DEST=/usr/local/lib/${PY_VERSION}/${PY_PKG_LOC}
PY_PROTOS=${PYTHON_BUILD}/gen/
PY_LTE=${MAGMA_ROOT}/lte/gateway/python
PY_ORC8R=${MAGMA_ROOT}/orc8r/gateway/python
PY_TMP_BUILD=/tmp/build-${PKGNAME}
PY_TMP_BUILD_SUFFIX=/usr/lib/python3/${PY_PKG_LOC}

PWD=`pwd`

glob_files () {
    # Given a list of files represented by the pattern in $1, and a package
    # output location in $2, generate a string of file locations that can be
    # passed to FPM. If $1 is a glob, you MUST surround it with quotes!
    #
    # For example, if you wanted to have all the files maching
    # foo/bar/*.yml end up in /etc/magma/, you would call:
    #
    # glob_files "foo/bar/*.yml" /etc/magma
    #
    # which would return:
    #
    # foo/bar/baz.yml=/etc/magma foo/bar/qux.yml=/etc/magma
    #
    # This is useful because fpm only accepts individual files or entire
    # directories for the dir package source type.
    RES=""
    for f in $1
    do
        RES="$RES $f=$2"
    done

    echo $RES
}

# The resulting package is placed in $OUTPUT_DIR
# or in the cwd.
if [ -z "$1" ]; then
  OUTPUT_DIR=${PWD}
else
  OUTPUT_DIR=$1
  if [ ! -d "$OUTPUT_DIR" ]; then
    echo "error: $OUTPUT_DIR is not a valid directory. Exiting..."
    exit 1
  fi
fi

# Build OAI and sessiond C/C++ services
cd "${MAGMA_ROOT}/lte/gateway"
OAI_BUILD="${C_BUILD}/oai"
SESSIOND_BUILD="${C_BUILD}/session_manager"
SCTPD_BUILD="${C_BUILD}/sctpd"

make build_oai BUILD_TYPE="${BUILD_TYPE}"
make build_session_manager BUILD_TYPE="${BUILD_TYPE}"
make build_sctpd BUILD_TYPE="${BUILD_TYPE}"

# Next, gather up the python files and put them into a build path.
#
# Note: Debian-based distributions install packages by default into a
# dist-packages directory, which is different than other distros, which drop
# packages into the site-packages directory.

# clean python build dir
if [ -d ${PY_TMP_BUILD} ]; then
    rm -r ${PY_TMP_BUILD}
fi

FULL_VERSION=${VERSION}-$(date +%s)-${COMMIT_HASH}

# first do python protos and then build the python packages.
# library will be dropped in $PY_TMP_BUILD/usr/lib/python3/dist-packages
# scripts will be dropped in $PY_TMP_BUILD/usr/bin.
# Use pydep to generate the lockfile and python deps
cd ${PY_ORC8R}
make protos
PKG_VERSION=${FULL_VERSION} ${PY_VERSION} setup.py install --root ${PY_TMP_BUILD} --install-layout deb \
    --no-compile --single-version-externally-managed
${RELEASE_DIR}/pydep finddep -l ${RELEASE_DIR}/magma.lockfile setup.py
ORC8R_PY_DEPS=`${RELEASE_DIR}/pydep lockfile ${RELEASE_DIR}/magma.lockfile`

cd ${PY_LTE}
make protos
PKG_VERSION=${FULL_VERSION} ${PY_VERSION} setup.py install --root ${PY_TMP_BUILD} --install-layout deb \
    --no-compile --single-version-externally-managed
${RELEASE_DIR}/pydep finddep -l ${RELEASE_DIR}/magma.lockfile setup.py
LTE_PY_DEPS=`${RELEASE_DIR}/pydep lockfile ${RELEASE_DIR}/magma.lockfile`

# now the binaries are built, we can package up everything else and build the
# magma package.
PKGFILE=${PKGNAME}_${FULL_VERSION}_${ARCH}.${PKGFMT}
BUILD_PATH=${OUTPUT_DIR}/${PKGFILE}

cd $PWD
# remove old packages
if [ -f ${BUILD_PATH} ]; then
  rm ${BUILD_PATH}
fi

SERVICE_DIR="/etc/systemd/system/"
ANSIBLE_FILES="${MAGMA_ROOT}/lte/gateway/deploy/roles/magma/files"

SCTPD_VERSION_FILE=$(mktemp)
SCTPD_MIN_VERSION_FILE=$(mktemp)

# files to be removed should be safely named (no special chars from mktemp)
# use current value (see https://github.com/koalaman/shellcheck/wiki/SC2064)
# shellcheck disable=SC2064
trap "rm -f '${SCTPD_VERSION_FILE}' '${SCTPD_MIN_VERSION_FILE}'" EXIT

echo "${FULL_VERSION}" > "${SCTPD_VERSION_FILE}"
echo "${SCTPD_MIN_VERSION}" > "${SCTPD_MIN_VERSION_FILE}"

BUILDCMD="fpm \
-s dir \
-t ${PKGFMT} \
-a ${ARCH} \
-n ${SCTPD_PKGNAME} \
-v ${FULL_VERSION} \
--provides ${SCTPD_PKGNAME} \
--replaces ${SCTPD_PKGNAME} \
--package ${OUTPUT_DIR}/${SCTPD_PKGNAME}_${FULL_VERSION}_${ARCH}.${PKGFMT} \
--description 'Magma SCTPD' \
--exclude '*/.ignoreme' \
${SCTPD_BUILD}/sctpd=/usr/local/sbin/ \
${SCTPD_VERSION_FILE}=/usr/local/share/sctpd/version \
$(glob_files "${SERVICE_DIR}/sctpd.service" /etc/systemd/system/sctpd.service)"

eval "$BUILDCMD"

BUILDCMD="fpm \
-s dir \
-t ${PKGFMT} \
-a ${ARCH} \
-n ${PKGNAME} \
-v ${FULL_VERSION} \
--provides ${PKGNAME} \
--replaces ${PKGNAME} \
--package ${BUILD_PATH} \
--description 'Magma Access Gateway' \
--after-install ${POSTINST} \
--exclude '*/.ignoreme' \
--config-files /etc/sysctl.d/99-magma.conf \
${ORC8R_PY_DEPS} \
${LTE_PY_DEPS} \
${SYSTEM_DEPS} \
${OAI_BUILD}/oai_mme/mme=/usr/local/bin/ \
${SESSIOND_BUILD}/sessiond=/usr/local/bin/ \
${SCTPD_MIN_VERSION_FILE}=/usr/local/share/magma/sctpd_min_version \
$(glob_files "${SERVICE_DIR}/magma@.service" /etc/systemd/system/magma@.service) \
$(glob_files "${SERVICE_DIR}/magma@control_proxy.service" /etc/systemd/system/magma@control_proxy.service) \
$(glob_files "${SERVICE_DIR}/magma@magmad.service" /etc/systemd/system/magma@magmad.service) \
$(glob_files "${SERVICE_DIR}/magma@mme.service" /etc/systemd/system/magma@mme.service) \
$(glob_files "${SERVICE_DIR}/magma@sessiond.service" /etc/systemd/system/magma@sessiond.service) \
$(glob_files "${SERVICE_DIR}/magma@mobilityd.service" /etc/systemd/system/magma@mobilityd.service) \
$(glob_files "${SERVICE_DIR}/magma@pipelined.service" /etc/systemd/system/magma@pipelined.service) \
$(glob_files "${SERVICE_DIR}/magma@redirectd.service" /etc/systemd/system/magma@redirectd.service) \
$(glob_files "${SERVICE_DIR}/magma@dnsd.service" /etc/systemd/system/magma@dnsd.service) \
$(glob_files "${SERVICE_DIR}/magma@lighttpd.service" /etc/systemd/system/magma@lighttpd.service) \
$(glob_files "${SERVICE_DIR}/magma@redis.service" /etc/systemd/system/magma@redis.service) \
$(glob_files "${SERVICE_DIR}/magma@td-agent-bit.service" /etc/systemd/system/magma@td-agent-bit.service) \
${CERT_FILE}=/var/opt/magma/certs/rootCA.pem \
$(glob_files "${MAGMA_ROOT}/lte/gateway/configs/!(control_proxy.yml|pipelined.yml|sessiond.yml)" /etc/magma/) \
$(glob_files "${MAGMA_ROOT}/lte/gateway/configs/pipelined.yml_prod" /etc/magma/pipelined.yml) \
$(glob_files "${MAGMA_ROOT}/lte/gateway/configs/sessiond.yml_prod" /etc/magma/sessiond.yml) \
$(glob_files "${MAGMA_ROOT}/lte/gateway/configs/templates/*" /etc/magma/templates/) \
$(glob_files "${MAGMA_ROOT}/orc8r/gateway/configs/templates/*" /etc/magma/templates/) \
${CONTROL_PROXY_FILE}=/etc/magma/ \
$(glob_files "${ANSIBLE_FILES}/magma_modules_load" /etc/modules-load.d/magma.conf) \
$(glob_files "${ANSIBLE_FILES}/logrotate_oai.conf" /etc/logrotate.d/oai) \
$(glob_files "${ANSIBLE_FILES}/local-cdn/*" /var/www/local-cdn/) \
${ANSIBLE_FILES}/99-magma.conf=/etc/sysctl.d/ \
${ANSIBLE_FILES}/magma_ifaces_gtp=/etc/network/interfaces.d/gtp \
${ANSIBLE_FILES}/20auto-upgrades=/etc/apt/apt.conf.d/20auto-upgrades \
${ANSIBLE_FILES}/coredump=/usr/local/bin/ \
${MAGMA_ROOT}/lte/gateway/release/stretch_snapshot=/usr/local/share/magma/ \
${MAGMA_ROOT}/orc8r/tools/ansible/roles/fluent_bit/files/60-fluent-bit.conf=/etc/rsyslog.d/60-fluent-bit.conf \
${PY_PROTOS}=${PY_DEST} \
$(glob_files "${PY_TMP_BUILD}/${PY_TMP_BUILD_SUFFIX}/${PKGNAME}*" ${PY_DEST}) \
$(glob_files "${PY_TMP_BUILD}/${PY_TMP_BUILD_SUFFIX}/*.egg-info" ${PY_DEST}) \
$(glob_files "${PY_TMP_BUILD}/usr/bin/*" /usr/local/bin/)"

eval "$BUILDCMD"
