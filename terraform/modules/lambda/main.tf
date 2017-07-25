resource "aws_lambda_function" "demo_lambda" {
  function_name = "${var.function_name}"
  handler       = "handler.Handle"
  runtime       = "python2.7"
  filename      = "${var.code_zip_path}"
  role          = "${var.role_name}"
}
