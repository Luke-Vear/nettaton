#!/usr/bin/env bash

set -e

PROJECT="nettaton"
R53_ZONE_ID="ZYJ1F75JP3JG"
BUCKET_REGION="eu-west-1"

PROJECT_ROOT="$(dirname "$(readlink -f ${BASH_SOURCE})")/.."
CMD_DIR="${PROJECT_ROOT}/cmd"
BUILD_LIST=$(find ${CMD_DIR}/* -type d)
ARTEFACT_NAME="handler"

WEB_DEPLOY_DIR="${PROJECT_ROOT}/web/deploy"

cd ${PROJECT_ROOT}

#### Test
unit() {
  for t in $(go list ./... | grep -v -E 'vendor|sandbox'); do
    go test -cover ${t} 
  done
}

#### Build
build_backend() {
  for dir in ${BUILD_LIST[@]}; do
    cd ${dir}
    component="${dir##*/}"
    echo "Building ${component}"

    go build -o "${ARTEFACT_NAME}"
    touch --date=@0 "${ARTEFACT_NAME}"
    zip -X -j "${ARTEFACT_NAME}".zip "${ARTEFACT_NAME}"

    rm -vf "${ARTEFACT_NAME}"
    cd - >/dev/null
  done
}

build_frontend() {
  npm run build --prefix web
}

#### TF
tf() {
  chkenv

  ifrm ${WEB_DEPLOY_DIR}
  cp -rpf "${PROJECT_ROOT}/web/dist" ${WEB_DEPLOY_DIR}

  sed -i "s/f3dc7042-bc46-42d2-9f8f-41417d48ca4d/$ENV/" $(find  ${WEB_DEPLOY_DIR} -name "*.js")

  tf_args=(
    "--var env=${ENV}"
    "--var r53_zone_id=${R53_ZONE_ID}"
    "--var web_deploy_dir=${WEB_DEPLOY_DIR}"
    "--var web_js=$(find ${WEB_DEPLOY_DIR} -name "*.js" -exec basename {} \+)"
    "--var web_css=$(find ${WEB_DEPLOY_DIR} -name "*.css" -exec basename {} \+)"
  )

  if [[ ${1} != plan ]]; then
    tf_args+=("--auto-approve")
  fi

  cd deployments

  terraform init \
    --backend-config="bucket=${PROJECT}-${ENV}-tfstate" \
    --backend-config="key=terraform.tfstate" \
    --backend-config="region=${BUCKET_REGION}"

  terraform ${1} ${tf_args[@]}
  clean_tf

  cd - >/dev/null
}

plan() {
  tf plan
}

deploy() {
  tf apply
}

destroy() {
  tf destroy
}

#### Serve
serve() {
  npm run start --prefix web
}

#### Smoketest
smoke() {
  chkenv
  go run test/smoketest.go --env ${ENV}
}

#### Clean
clean_tf() {
  rm_targets=(
    "${PROJECT_ROOT}/deployments/.terraform/modules"
    "${PROJECT_ROOT}/deployments/.terraform/terraform.tfstate"
  )

  for rmt in ${rm_targets[@]}; do  
    ifrm ${rmt}
  done
}
clean() {
  clean_tf

  find "${PROJECT_ROOT}" -name "*.zip" -exec rm {} \+

  rm_targets=(
    "${PROJECT_ROOT}/web/dist"
    ${WEB_DEPLOY_DIR}
  )
  
  for rmt in ${rm_targets[@]}; do  
    ifrm ${rmt}
  done
}

#### Misc
chkenv() {
  if [[ -z ${ENV} ]]; then
    echo 'ENV must be set'
    exit 1
  fi

  export ${ENV}
}

ifrm() {
  if [[ -a ${1} ]]; then 
    rm -rf ${1}
  fi
}

$1
exit 0