provider "aws" {
  region = "eu-west-1"
}

resource "aws_api_gateway_rest_api" "nettatonapi" {
  name        = "nettatonapi"
  description = "This is my API for demonstration purposes"
}

resource "aws_api_gateway_resource" "question" {
  rest_api_id = "${aws_api_gateway_rest_api.nettatonapi.id}"
  parent_id   = "${aws_api_gateway_rest_api.nettatonapi.root_resource_id}"
  path_part   = "question"
}

resource "aws_api_gateway_method" "question_get" {
  rest_api_id   = "${aws_api_gateway_rest_api.nettatonapi.id}"
  resource_id   = "${aws_api_gateway_resource.question.id}"
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_method" "question_post" {
  rest_api_id   = "${aws_api_gateway_rest_api.nettatonapi.id}"
  resource_id   = "${aws_api_gateway_resource.question.id}"
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_resource" "user" {
  rest_api_id = "${aws_api_gateway_rest_api.nettatonapi.id}"
  parent_id   = "${aws_api_gateway_rest_api.nettatonapi.root_resource_id}"
  path_part   = "user"
}

resource "aws_api_gateway_method" "user_post" {
  rest_api_id   = "${aws_api_gateway_rest_api.nettatonapi.id}"
  resource_id   = "${aws_api_gateway_resource.user.id}"
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_resource" "user_id" {
  rest_api_id = "${aws_api_gateway_rest_api.nettatonapi.id}"
  parent_id   = "${aws_api_gateway_resource.user.id}"
  path_part   = "{id}"
}

resource "aws_api_gateway_resource" "score" {
  rest_api_id = "${aws_api_gateway_rest_api.nettatonapi.id}"
  parent_id   = "${aws_api_gateway_resource.user_id.id}"
  path_part   = "score"
}

resource "aws_api_gateway_method" "score_get" {
  rest_api_id   = "${aws_api_gateway_rest_api.nettatonapi.id}"
  resource_id   = "${aws_api_gateway_resource.score.id}"
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_resource" "token" {
  rest_api_id = "${aws_api_gateway_rest_api.nettatonapi.id}"
  parent_id   = "${aws_api_gateway_resource.user_id.id}"
  path_part   = "token"
}

resource "aws_api_gateway_method" "token_post" {
  rest_api_id   = "${aws_api_gateway_rest_api.nettatonapi.id}"
  resource_id   = "${aws_api_gateway_resource.token.id}"
  http_method   = "POST"
  authorization = "NONE"
}
