provider "aws" {
  region = "eu-west-1"
}

# resource "aws_instance" "example" {
#   count    = 5
#   function = "${element(var.functions, count.index)}"

#   tags {
#     Name = "example-${count.index}"
#   }
# }

module "answer" {
  source        = "../modules/lambda"
  function_name = "nettaton-test-answer"
  code_zip_path = "../../cmd/answer/handler.zip"
  role_name     = "${module.iam.role_arn}"
}

module "question" {
  source        = "../modules/lambda"
  function_name = "nettaton-test-question"
  code_zip_path = "../../cmd/answer/handler.zip"
  role_name     = "${module.iam.role_arn}"
}

module "iam" {
  source    = "../modules/iam"
  role_name = "nettaton-test"
}
