#!/usr/bin/env bash

set -e

PROJECT="nettaton"
PROJECT_ROOT=$(dirname "$(readlink -f ${BASH_SOURCE})")
BUCKET_REGION="eu-west-1"
ENVIRONMENT="${1}"
CMD_DIR="${PROJECT_ROOT}/cmd"
BUILD_LIST=$(find ${CMD_DIR}/* -type d)
ARTEFACT_NAME="handler"

if [[ -z $ENVIRONMENT ]]; then
  echo 'Need $1 set to environment'
  exit 1
fi

#### Test ####################################################################

echo "---- ---- ---- ----"
echo "  Unit Test Init"
echo "---- ---- ---- ----"

for t in $(go list ./... | grep -v -E 'vendor|sandbox'); do
  go test -cover ${t} 
done

echo "---- ---- ---- ----"
echo " Unit Test Complete"
echo "---- ---- ---- ----"

#### Build ###################################################################

echo "---- ---- ---- ----"
echo "   Building Init"
echo "---- ---- ---- ----"

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


#### Terraform ###############################################################

echo "---- ---- ---- ----"
echo "  Terraform Init"
echo "---- ---- ---- ----"

cd terraform

terraform init \
  -backend-config="bucket=${PROJECT}-${ENVIRONMENT}-tfstate" \
  -backend-config="key=terraform.tfstate" \
  -backend-config="region=${BUCKET_REGION}"

terraform get -update

terraform apply \
  --var-file="./vars/${ENVIRONMENT}.tfvars" \
  -auto-approve

rm -rf .terraform

cd - >/dev/null

echo "---- ---- ---- ----"
echo " Terraform Complete"
echo "---- ---- ---- ----"

exit 0