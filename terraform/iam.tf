data "aws_iam_policy_document" "lambda_role" {
  statement {
    actions = [
      "sts:AssumeRole",
    ]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "lambda_role" {
  name               = "${title(var.name)}${title(var.env)}LambdaRole"
  path               = "/${var.name}-${var.env}/"
  assume_role_policy = "${data.aws_iam_policy_document.lambda_role.json}"

  lifecycle {
    create_before_destroy = true
  }
}

data "aws_iam_policy_document" "dynamo_crud" {
  statement {
    actions = [
      "dynamodb:GetItem",
      "dynamodb:PutItem",
      "dynamodb:UpdateItem",
      "dynamodb:DeleteItem",
    ]

    resources = [
      "${aws_dynamodb_table.questionstore.arn}",
    ]
  }
}

resource "aws_iam_policy" "dynamo_crud" {
  name   = "${title(var.name)}${title(var.env)}QuestionTableCRUD"
  path   = "/${var.name}-${var.env}/"
  policy = "${data.aws_iam_policy_document.dynamo_crud.json}"
}

resource "aws_iam_role_policy_attachment" "dynamo_crud" {
  role       = "${aws_iam_role.lambda_role.name}"
  policy_arn = "${aws_iam_policy.dynamo_crud.arn}"
}

data "aws_iam_policy_document" "lambda_logging" {
  statement {
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = [
      "arn:aws:logs:*:*:*",
    ]
  }
}

resource "aws_iam_policy" "lambda_logging" {
  name   = "${title(var.name)}${title(var.env)}LambdaLogging"
  path   = "/${var.name}-${var.env}/"
  policy = "${data.aws_iam_policy_document.lambda_logging.json}"
}

resource "aws_iam_role_policy_attachment" "lambda_logging" {
  role       = "${aws_iam_role.lambda_role.name}"
  policy_arn = "${aws_iam_policy.lambda_logging.arn}"
}
