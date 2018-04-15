#!/usr/bin/env bash

set -e

echo "---- ---- ---- ----"
echo "   Building Init"
echo "---- ---- ---- ----"

PROJECT_ROOT=$(dirname "$(readlink -f ${BASH_SOURCE})")
CMD_DIR="${PROJECT_ROOT}/cmd"
BUILD_LIST=$(find ${CMD_DIR}/* -type d)
ARTEFACT_NAME="handler"

cd ${PROJECT_ROOT}
for dir in ${BUILD_LIST[@]}; do
  cd $dir
  component="${dir##*/}"
  echo "Building ${component}"

  go build -o "${ARTEFACT_NAME}"
  touch --date=@0 "${ARTEFACT_NAME}"
  zip -X -j "${ARTEFACT_NAME}".zip "${ARTEFACT_NAME}"

  rm -vf "${ARTEFACT_NAME}"
  cd - >/dev/null
done

echo "---- ---- ---- ----"
echo " Building Complete"
echo "---- ---- ---- ----"