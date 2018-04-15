resource "aws_iam_role" "lambda_role" {
  name = "${title(var.name)}${title(var.env)}LambdaRole"

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

resource "aws_iam_policy" "dynamo_policy" {
  name = "${title(var.name)}${title(var.env)}QuestionTableCRUD"

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
      "Resource":"${aws_dynamodb_table.questionstore.arn}"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "attach" {
  role       = "${aws_iam_role.lambda_role.name}"
  policy_arn = "${aws_iam_policy.dynamo_policy.arn}"
}
