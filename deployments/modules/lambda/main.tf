# lambda/main.tf

resource "aws_lambda_function" "go_func" {
  function_name = "${var.name}-${var.env}-${var.app}"
  runtime       = "go1.x"

  s3_bucket        = "${var.artefact_bucket}"
  s3_key           = "${var.app}/handler.zip"
  source_code_hash = "${base64sha256(file("${var.artefact_dir}/handler.zip"))}"
  handler          = "handler"

  role = "${var.lambda_role_arn}"

  environment {
    variables = {
      "NETTATON_TABLE" = "${var.dynamo_table}"
    }
  }
}

resource "aws_api_gateway_method" "method" {
  rest_api_id   = "${var.api_id}"
  resource_id   = "${var.api_resource_id}"
  http_method   = "${var.api_http_method}"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "integration" {
  rest_api_id             = "${var.api_id}"
  resource_id             = "${var.api_resource_id}"
  http_method             = "${aws_api_gateway_method.method.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:${var.region}:lambda:path/2015-03-31/functions/${aws_lambda_function.go_func.arn}/invocations"
}

resource "aws_lambda_permission" "invoke" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.go_func.function_name}"
  principal     = "apigateway.amazonaws.com"
  source_arn    = "arn:aws:execute-api:${var.region}:${var.account_id}:${var.api_id}/*/${var.api_http_method}${var.api_resource_path}"
}

resource "aws_s3_bucket_object" "createquestion" {
  bucket = "${var.artefact_bucket}"
  key    = "${var.app}/handler.zip"
  source = "${var.artefact_dir}/handler.zip"
  etag   = "${md5(file("${var.artefact_dir}/handler.zip"))}"
}
