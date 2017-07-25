#!/usr/bin/env bash

set -eu

PROJECT_ROOT=$(dirname "$(readlink -f ${BASH_SOURCE})")
CMD_DIR="${PROJECT_ROOT}/cmd"
BUILD_LIST=$(find ${CMD_DIR}/* -type d)

cd ${PROJECT_ROOT}
for dir in ${BUILD_LIST[@]}; do
  cd $dir
  echo "Building ${dir##*/}"
  make #&
  cd - >/dev/null
done

echo "---- ---- ---- ----"
echo "Building Complete"
echo "---- ---- ---- ----"