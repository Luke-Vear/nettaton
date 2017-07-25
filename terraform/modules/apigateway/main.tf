# apigateway/main.tf

resource "aws_api_gateway_rest_api" "api" {
  name = "${var.api_name}"
}

output "api_id" {
  value = "${aws_api_gateway_rest_api.api.id}"
}

resource "aws_api_gateway_resource" "question" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  parent_id   = "${aws_api_gateway_rest_api.api.root_resource_id}"
  path_part   = "question"
}

output "question_resource_id" {
  value = "${aws_api_gateway_resource.question.id}"
}

output "question_resource_path" {
  value = "${aws_api_gateway_resource.question.path}"
}

resource "aws_api_gateway_resource" "user" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  parent_id   = "${aws_api_gateway_rest_api.api.root_resource_id}"
  path_part   = "user"
}

output "user_resource_id" {
  value = "${aws_api_gateway_resource.user.id}"
}

resource "aws_api_gateway_resource" "user_id" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  parent_id   = "${aws_api_gateway_resource.user.id}"
  path_part   = "{id}"
}

resource "aws_api_gateway_resource" "score" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  parent_id   = "${aws_api_gateway_resource.user_id.id}"
  path_part   = "score"
}

output "score_resource_id" {
  value = "${aws_api_gateway_resource.score.id}"
}

resource "aws_api_gateway_resource" "token" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  parent_id   = "${aws_api_gateway_resource.user_id.id}"
  path_part   = "token"
}

output "token_resource_id" {
  value = "${aws_api_gateway_resource.token.id}"
}
