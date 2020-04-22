# Cwf Secrets

Cwf Secrets is used to apply a set of secrets for a cwf gateway. These secrets
provide permanent gateway identification so that gateway pods will be recreated
with the same hwid and challenge key.

## TL;DR;

```bash
# Copy secrets into subchart root
$ mkdir charts/secrets/.secrets && \
    cp -r <secrets>/* charts/secrets/.secrets/
$ ls charts/secrets/.secrets
snowflake gw_challenge.key

# Apply secrets
helm template charts/secrets \
    --name <cwf release name> \
    --namespace magma | kubectl -n magma apply -f -
```

## Overview

This chart installs a secret that serves as identifiers for the gateway. 
The secrets are expected to be provided as files and placed under
secrets subchart root.
```bash
$ ls charts/secrets/.secrets
snowflake  gw_challenge.key
$ pwd
magma/cwf/gateway/helm/cwf
```

## Creating Gateway Info
If creating a gateway for the first time, you'll need to create a snowflake
and challenge key before installing the gateway. To do so:

```
$ docker login <DOCKER REGISTRY>
$ docker pull <DOCKER REGISTRY>/gateway_python:<IMAGE_VERSION>
$ docker run -d <DOCKER_REGISTRY>/gateway_python:<IMAGE_VERSION> python3.5 -m magma.magmad.main

This will output a container ID such as
f3bc383a95db16f2e448fdf67cac133a5f9019375720b59477aebc96bacd05a9

Run the following, substituting your container ID here
$ docker cp <container ID>:/etc/snowflake charts/secrets/.secrets
$ docker cp <container ID>:/var/opt/magma/certs/gw_challenge.key /charts/secrets/.secrets
```

Otherwise if redploying a gateway with permanent gwinfo, copy the existing 
snowflake from `etc/snowflake` and challenge key at 
`/var/opt/magma/certs/gw_challenge.key`

## Configuration

The following table lists the configurable secret locations, 
docker credentials and their default values.

| Parameter        | Description     | Default   |
| ---              | ---             | ---       |
| `create` | Set to ``true`` to create cwf secrets. | `false` |
| `secret.gwinfo` | Root relative secrets directory. | `.secrets` |
