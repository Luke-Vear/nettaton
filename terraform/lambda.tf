data "aws_caller_identity" "current" {}

module "createquestion" {
  source            = "./modules/lambda"
  name              = "${var.name}"
  env               = "${var.env}"
  app               = "createquestion"
  region            = "${var.region}"
  lambda_role_arn   = "${aws_iam_role.lambda_role.arn}"
  account_id        = "${data.aws_caller_identity.current.account_id}"
  api_id            = "${aws_api_gateway_rest_api.api.id}"
  api_resource_id   = "${aws_api_gateway_resource.question.id}"
  api_resource_path = "${aws_api_gateway_resource.question.path}"
  api_http_method   = "POST"
  artefact_bucket   = "${aws_s3_bucket.artefact.id}"
  artefact_dir      = "${var.cmd_path}/createquestion"
  dynamo_table      = "${aws_dynamodb_table.questionstore.name}"
}

module "readquestion" {
  source            = "./modules/lambda"
  name              = "${var.name}"
  env               = "${var.env}"
  app               = "readquestion"
  region            = "${var.region}"
  lambda_role_arn   = "${aws_iam_role.lambda_role.arn}"
  account_id        = "${data.aws_caller_identity.current.account_id}"
  api_id            = "${aws_api_gateway_rest_api.api.id}"
  api_resource_id   = "${aws_api_gateway_resource.question_id.id}"
  api_resource_path = "${aws_api_gateway_resource.question_id.path}"
  api_http_method   = "GET"
  artefact_bucket   = "${aws_s3_bucket.artefact.id}"
  artefact_dir      = "${var.cmd_path}/readquestion"
  dynamo_table      = "${aws_dynamodb_table.questionstore.name}"
}

module "answerquestion" {
  source            = "./modules/lambda"
  name              = "${var.name}"
  env               = "${var.env}"
  app               = "answerquestion"
  region            = "${var.region}"
  lambda_role_arn   = "${aws_iam_role.lambda_role.arn}"
  account_id        = "${data.aws_caller_identity.current.account_id}"
  api_id            = "${aws_api_gateway_rest_api.api.id}"
  api_resource_id   = "${aws_api_gateway_resource.question_id_answer.id}"
  api_resource_path = "${aws_api_gateway_resource.question_id_answer.path}"
  api_http_method   = "POST"
  artefact_bucket   = "${aws_s3_bucket.artefact.id}"
  artefact_dir      = "${var.cmd_path}/answerquestion"
  dynamo_table      = "${aws_dynamodb_table.questionstore.name}"
}
