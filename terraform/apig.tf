resource "aws_api_gateway_rest_api" "api" {
  name = "${var.name}-${var.env}"
}

resource "aws_api_gateway_resource" "question" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  parent_id   = "${aws_api_gateway_rest_api.api.root_resource_id}"
  path_part   = "question"
}

resource "aws_api_gateway_resource" "question_id" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  parent_id   = "${aws_api_gateway_resource.question.id}"
  path_part   = "{id}"
}

resource "aws_api_gateway_resource" "question_id_answer" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  parent_id   = "${aws_api_gateway_resource.question_id.id}"
  path_part   = "answer"
}
