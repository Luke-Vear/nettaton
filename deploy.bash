#!/usr/bin/env bash

PROJECT="nettaton"
ENVIRONMENT="${1}"
COMPONENT="${PWD##*/}"


if [[ -z $ENVIRONMENT ]]; then
  echo 'Need $1 set to environment'
  exit 1
fi

ROLE_NAME="${PROJECT}-${ENVIRONMENT}"
FUNC_NAME="${PROJECT}-${ENVIRONMENT}-${COMPONENT}"

LAMBDA_ROLE_ARN=$(aws iam get-role \
  --role-name ${ROLE_NAME} \
  --query 'Role.Arn' \
  --output text)

echo "Building ${COMPONENT}"
make || exit 1

echo "Updating Lambda function"
aws lambda update-function-code \
  --function-name ${FUNC_NAME} \
  --zip-file fileb://handler.zip

echo "Removing Files"
rm -v handler.{so,zip}
