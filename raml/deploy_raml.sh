#!/usr/bin/env bash

echo "script to deploy raml to s3 bucket"
exit 1

PROJECT="nettaton"
ENVIRONMENT="${1}"
FUNC_LIST=(answer login question userscore register)


if [[ -z $ENVIRONMENT ]]; then
  echo 'Need $1 set to environment'
  exit 1
fi

#### DynamoDB ################################################################
TABLE_NAME="${PROJECT}-${ENVIRONMENT}-userdb"

echo "Checking for ${TABLE_NAME}"
DYNAMO_ARN=$(aws dynamodb describe-table \
  --table-name ${TABLE_NAME} \
  --query 'Table.TableArn' \
  --output text 2>/dev/null)

if (( $? != 0 )); then
  echo "No scoredb, creating"
  DYNAMO_ARN=$(aws dynamodb create-table \
    --table-name ${TABLE_NAME} \
    --attribute-definitions AttributeName=id,AttributeType=S \
    --key-schema AttributeName=id,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=3,WriteCapacityUnits=3 \
    --query 'TableDescription.TableArn' \
    --output text)
fi

aws dynamodb wait table-exists \
  --table-name ${TABLE_NAME}

#### IAM #####################################################################
echo "Checking for ${ENVIRONMENT} lambda role"

# check roles
LAMBDA_ROLE_ARN=$(aws iam get-role \
  --role-name ${PROJECT}-${ENVIRONMENT} \
  --query 'Role.Arn' \
  --output text 2>/dev/null)

if (( $? != 0 )); then
  echo "No ${ENVIRONMENT} lambda role, creating"
  POL_DOC=$(cat <<EOF
{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":["dynamodb:BatchGetItem","dynamodb:BatchWriteItem","dynamodb:DeleteItem","dynamodb:GetItem","dynamodb:GetRecords","dynamodb:PutItem","dynamodb:Query","dynamodb:Scan","dynamodb:UpdateItem"],"Resource":["${DYNAMO_ARN}"]}]}
EOF
)

  # create-policy
  POLICY_ARN=$(aws iam create-policy \
    --policy-name "${TABLE_NAME}" \
    --policy-document "${POL_DOC}" \
    --query 'Policy.Arn' \
    --output text)

  # create-role
  LAMBDA_ROLE_ARN=$(aws iam create-role \
    --role-name ${PROJECT}-${ENVIRONMENT} \
    --assume-role-policy-document '{
      "Statement": [{
        "Effect": "Allow",
        "Principal": {
          "Service": "lambda.amazonaws.com"
        },
        "Action": "sts:AssumeRole"
      }]
    }' \
    --query 'Role.Arn' \
    --output text 2>/dev/null)

  # attach to role, the policy
  aws iam attach-role-policy \
    --role-name ${PROJECT}-${ENVIRONMENT} \
    --policy-arn ${POLICY_ARN} \
    || exit 1
fi

#### Lambda ##################################################################
for component in ${FUNC_LIST[@]}; do
  echo "Creating ${component} function"
  cd ./cmd/"${component}"

  LAMBDA_NAME="${PROJECT}-${ENVIRONMENT}-${component}"

  echo "Building ${component}"
  make || exit 1

  echo "Deploying ${component}"
  LAMBDA_FUNC_ARN=$(aws lambda create-function \
    --function-name "${LAMBDA_NAME}" \
    --zip-file fileb://handler.zip \
    --role ${LAMBDA_ROLE_ARN} \
    --runtime python2.7 \
    --handler handler.Handle \
    --query 'FunctionArn' \
    --output text 2>/dev/null)
  
  rm -vf handler.{so,zip}
  cd -
done
