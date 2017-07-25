# iam/main.tf

resource "aws_iam_role" "lambda_role" {
  name = "${var.role_name}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

output "role_arn" {
  value = "${aws_iam_role.lambda_role.arn}"
}

resource "aws_iam_policy" "dynamo_policy" {
  name = "${var.role_name}-userdb-rw"

  policy = <<EOF
{
  "Version":"2012-10-17",
  "Statement":[
    {
      "Sid":"GetPutUpdateDeleteDynamoDB", 
      "Effect":"Allow",
      "Action":[
        "dynamodb:GetItem",
        "dynamodb:PutItem",
        "dynamodb:UpdateItem",
        "dynamodb:DeleteItem"
      ],
      "Resource":"${var.dynamodb_arn}"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "attach" {
  role       = "${aws_iam_role.lambda_role.name}"
  policy_arn = "${aws_iam_policy.dynamo_policy.arn}"
}
