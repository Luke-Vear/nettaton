# lambda/main.tf

resource "aws_lambda_function" "function" {
  function_name = "${var.function_name}"
  handler       = "handler.Handle"
  runtime       = "python2.7"
  filename      = "${var.code_zip_path}"
  role          = "${var.role_name}"

  environment {
    variables = "${var.lambda_env_variables}"
  }
}

output "lambda_arn" {
  value = "${aws_lambda_function.function.arn}"
}

resource "aws_api_gateway_method" "method" {
  rest_api_id   = "${var.api_id}"
  resource_id   = "${var.resource_id}"
  http_method   = "${var.http_method}"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "integration" {
  rest_api_id             = "${var.api_id}"
  resource_id             = "${var.resource_id}"
  http_method             = "${aws_api_gateway_method.method.http_method}"
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:${var.region}:lambda:path/2015-03-31/functions/${aws_lambda_function.function.arn}/invocations"
}

resource "aws_lambda_permission" "permission" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.function.function_name}"
  principal     = "apigateway.amazonaws.com"
  source_arn    = "arn:aws:execute-api:${var.region}:${var.account_id}:${var.api_id}/*/${var.http_method}${var.resource_path}"
}
