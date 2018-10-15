#!/usr/bin/env bash

set -e

PROJECT="nettaton"
PROJECT_ROOT="$(dirname "$(readlink -f ${BASH_SOURCE})")/.."
BUCKET_REGION="eu-west-1"
ENVIRONMENT="${2}"
CMD_DIR="${PROJECT_ROOT}/cmd"
BUILD_LIST=$(find ${CMD_DIR}/* -type d)
ARTEFACT_NAME="handler"

cd ${PROJECT_ROOT}

#### Test ####################################################################

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

#### Build ###################################################################

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

  echo "---- ---- ---- ----"
  echo " Building Complete"
  echo "---- ---- ---- ----"
}

#### Deploy ##################################################################

deploy() {
  chkenv

  echo "---- ---- ---- ----"
  echo "  Deploy Init"
  echo "---- ---- ---- ----"

  cd terraform

  terraform init \
    -backend-config="bucket=${PROJECT}-${ENVIRONMENT}-tfstate" \
    -backend-config="key=terraform.tfstate" \
    -backend-config="region=${BUCKET_REGION}"

  terraform apply \
    --var-file="./vars/${ENVIRONMENT}.tfvars" \
    -auto-approve

  rm -rf .terraform

  cd - >/dev/null

  echo "---- ---- ---- ----"
  echo " Deploy Complete"
  echo "---- ---- ---- ----"
}

#### Smoketest ###############################################################

smoketest() {
  chkenv

  echo "---- ---- ---- ----"
  echo "  Smoketest Test Init"
  echo "---- ---- ---- ----"

  go run cmd/smoketest/main.go --env ${ENVIRONMENT}

  echo "---- ---- ---- ----"
  echo " Smoketest Test Complete"
  echo "---- ---- ---- ----"
}

#### Misc ####################################################################

chkenv() {
  if [[ -z $ENVIRONMENT ]]; then
    echo 'environment must be set'
    exit 1
  fi
}


$1
exit 0


main