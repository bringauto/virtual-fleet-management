#!/usr/bin/env bash

set -e

SCRIPT_DIR="$(dirname "$0")"
cd ${SCRIPT_DIR}

VERSION=$(sed -E -n 's/version=([^=]+)/\1/p' < version.txt)
MACHINE=$(uname -m | sed -E 's/_/-/')

INSTALL_DIR="./virtual-fleet-management_${VERSION}_${MACHINE}-linux"
INSTALL_DIR_SCENARIOS="./virtual-fleet-management-scenarios_${VERSION}_${MACHINE}-linux"

if [[ -d ${INSTALL_DIR} ]]; then
  echo "${INSTALL_DIR} already exist. Delete it please" >&2
  exit 1
fi

if [[ -d ${INSTALL_DIR_SCENARIOS} ]]; then
  echo "${INSTALL_DIR_SCENARIOS} already exist. Delete it please" >&2
  exit 1
fi

go get virtual-fleet-management
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w'

mkdir -p "${INSTALL_DIR}"
mkdir -p "${INSTALL_DIR_SCENARIOS}"

cp virtual-fleet-management       "${INSTALL_DIR}/"
cp README.md                      "${INSTALL_DIR}/"
#cp LICENSE                        "${INSTALL_DIR}/" // TODO: Uncomment this line after adding LICENSE file
cp -r scenarios                   "${INSTALL_DIR_SCENARIOS}/"

zip -r "virtual-fleet-management_v${VERSION}_${MACHINE}-linux.zip" ${INSTALL_DIR}/
zip -r "virtual-fleet-management-scenarios_v${VERSION}_${MACHINE}-linux.zip" ${INSTALL_DIR_SCENARIOS}/

rm -fr "${INSTALL_DIR}"
rm -fr "${INSTALL_DIR_SCENARIOS}"
