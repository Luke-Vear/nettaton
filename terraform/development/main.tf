# development/main.tf

provider "aws" {
  region = "${var.region}"
}

module "account" {
  source = "../modules/account"
}

module "iam" {
  source    = "../modules/iam"
  role_name = "${var.env}"
}

module "apig" {
  source   = "../modules/apigateway"
  api_name = "${var.env}"
}

module "question" {
  source               = "../modules/lambda"
  region               = "${var.region}"
  account_id           = "${module.account.account_id}"
  function_name        = "${var.env}-question"
  code_zip_path        = "../../cmd/question/handler.zip"
  role_name            = "${module.iam.role_arn}"
  api_id               = "${module.apig.api_id}"
  resource_id          = "${module.apig.question_resource_id}"
  resource_path        = "${module.apig.question_resource_path}"
  http_method          = "GET"
  lambda_env_variables = "${var.lambda_env_variables}"
}

module "answer" {
  source               = "../modules/lambda"
  region               = "${var.region}"
  account_id           = "${module.account.account_id}"
  function_name        = "${var.env}-answer"
  code_zip_path        = "../../cmd/answer/handler.zip"
  role_name            = "${module.iam.role_arn}"
  api_id               = "${module.apig.api_id}"
  resource_id          = "${module.apig.question_resource_id}"
  resource_path        = "${module.apig.question_resource_path}"
  http_method          = "POST"
  lambda_env_variables = "${var.lambda_env_variables}"
}
