#!/usr/bin/env bash

set -e

PROJECT="nettaton"
R53_ZONE_ID="ZYJ1F75JP3JG"
BUCKET_REGION="eu-west-1"

PROJECT_ROOT="$(dirname "$(readlink -f ${BASH_SOURCE})")/.."
CMD_DIR="${PROJECT_ROOT}/cmd"
BUILD_LIST=$(find ${CMD_DIR}/* -type d)
ARTEFACT_NAME="handler"

cd ${PROJECT_ROOT}

#### Test
unit() {
  echo "---- ---- ---- ----"
  echo "  Unit Test Init"
  echo "---- ---- ---- ----"

  for t in $(go list ./... | grep -v -E 'vendor|sandbox'); do
    go test -cover ${t} 
  done

  echo "---- ---- ---- ----"
  echo " Unit Test Complete"
  echo "---- ---- ---- ----"
}

#### Build
build() {
  echo "---- ---- ---- ----"
  echo "   Building Init"
  echo "---- ---- ---- ----"

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

  cd web
  GOOS=js GOARCH=wasm go build -o main.wasm
  cd ..
  echo "some zip thing"

  echo "---- ---- ---- ----"
  echo " Building Complete"
  echo "---- ---- ---- ----"
}

#### TF
tf() {
  chkenv

  echo "---- ---- ---- ----"
  echo "  TF $1 Init"
  echo "---- ---- ---- ----"

  cd deployments

  terraform init \
    -backend-config="bucket=${PROJECT}-${ENV}-tfstate" \
    -backend-config="key=terraform.tfstate" \
    -backend-config="region=${BUCKET_REGION}"

  terraform ${1} \
    --var env="${ENV}" \
    --var r53_zone_id="${R53_ZONE_ID}" \
    -input=false

  rm -rf .terraform

  cd - >/dev/null

  echo "---- ---- ---- ----"
  echo " TF $1 Complete"
  echo "---- ---- ---- ----"
}

plan() {
  tf plan
}

deploy() {
  tf apply
}

#### Smoketest
smoke() {
  chkenv

  echo "---- ---- ---- ----"
  echo "  Smoketest Init"
  echo "---- ---- ---- ----"

  go run test/smoketest.go --env ${ENV}

  echo "---- ---- ---- ----"
  echo " Smoketest Complete"
  echo "---- ---- ---- ----"
}

#### Misc
chkenv() {
  if [[ -z $ENV ]]; then
    echo 'ENV must be set'
    exit 1
  fi
}

$1
exit 0