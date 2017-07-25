# development/main.tf

provider "aws" {
  region = "${var.region}"
}

module "account" {
  source = "./modules/account"
}

module "iam" {
  source       = "./modules/iam"
  role_name    = "${var.name}-${var.env}"
  dynamodb_arn = "${module.dynamodb.db_arn}"
  db_crud_sid  = "${var.name}${var.env}PutUpdateDelete"
}

module "dynamodb" {
  source      = "./modules/dynamodb"
  db_capacity = "${var.db_capacity}"
  db_name     = "${var.lambda_env_variables["TABLE"]}"
}

module "apig" {
  source   = "./modules/apigateway"
  api_name = "${var.name}-${var.env}"
}

module "question" {
  source               = "./modules/lambda"
  region               = "${var.region}"
  account_id           = "${module.account.account_id}"
  function_name        = "${var.name}-${var.env}-question"
  code_zip_path        = "../cmd/question/handler.zip"
  role_name            = "${module.iam.role_arn}"
  api_id               = "${module.apig.api_id}"
  resource_id          = "${module.apig.question_resource_id}"
  resource_path        = "${module.apig.question_resource_path}"
  http_method          = "GET"
  lambda_env_variables = "${var.lambda_env_variables}"
}

module "answer" {
  source               = "./modules/lambda"
  region               = "${var.region}"
  account_id           = "${module.account.account_id}"
  function_name        = "${var.name}-${var.env}-answer"
  code_zip_path        = "../cmd/answer/handler.zip"
  role_name            = "${module.iam.role_arn}"
  api_id               = "${module.apig.api_id}"
  resource_id          = "${module.apig.question_resource_id}"
  resource_path        = "${module.apig.question_resource_path}"
  http_method          = "POST"
  lambda_env_variables = "${var.lambda_env_variables}"
}

module "register" {
  source               = "./modules/lambda"
  region               = "${var.region}"
  account_id           = "${module.account.account_id}"
  function_name        = "${var.name}-${var.env}-register"
  code_zip_path        = "../cmd/register/handler.zip"
  role_name            = "${module.iam.role_arn}"
  api_id               = "${module.apig.api_id}"
  resource_id          = "${module.apig.user_resource_id}"
  resource_path        = "${module.apig.user_resource_path}"
  http_method          = "POST"
  lambda_env_variables = "${var.lambda_env_variables}"
}

module "userscore" {
  source               = "./modules/lambda"
  region               = "${var.region}"
  account_id           = "${module.account.account_id}"
  function_name        = "${var.name}-${var.env}-userscore"
  code_zip_path        = "../cmd/userscore/handler.zip"
  role_name            = "${module.iam.role_arn}"
  api_id               = "${module.apig.api_id}"
  resource_id          = "${module.apig.score_resource_id}"
  resource_path        = "${module.apig.score_resource_path}"
  http_method          = "GET"
  lambda_env_variables = "${var.lambda_env_variables}"
}

module "login" {
  source               = "./modules/lambda"
  region               = "${var.region}"
  account_id           = "${module.account.account_id}"
  function_name        = "${var.name}-${var.env}-login"
  code_zip_path        = "../cmd/login/handler.zip"
  role_name            = "${module.iam.role_arn}"
  api_id               = "${module.apig.api_id}"
  resource_id          = "${module.apig.token_resource_id}"
  resource_path        = "${module.apig.token_resource_path}"
  http_method          = "POST"
  lambda_env_variables = "${var.lambda_env_variables}"
}
