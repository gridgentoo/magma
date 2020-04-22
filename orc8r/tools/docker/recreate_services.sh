#!/bin/bash
#
# Copyright (c) 2016-present, Facebook, Inc.
# All rights reserved.
#
# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree. An additional grant
# of patent rights can be found in the PATENTS file in the same directory.

# This script is intended to finish a docker-based upgrade by re-creating
# the docker containers with recently downloaded images

RUNNING_TAG=$(docker ps --filter name=magmad --format "{{.Image}}" | cut -d ":" -f 2)

source /var/opt/magma/docker/.env

# If tag running is equal to .env, then do nothing
if [ "$RUNNING_TAG" == "$IMAGE_VERSION" ]; then
  exit
fi

# Otherwise recreate containers with the new image
cd /var/opt/magma/docker || exit
/usr/local/bin/docker-compose up --force-recreate -d --remove-orphans

# Remove all stopped containers and dangling images
docker system prune -af
