#!/usr/bin/env bash

set -e

source hack/common.sh
source hack/docker-common.sh

${IMAGE_BUILD_CMD} build -t "${OPERATOR_FULL_IMAGE_NAME}" . \
    --build-arg="LDFLAGS=${LDFLAGS}" --platform="${TARGET_OS}"/"${TARGET_ARCH}"
