variable "region" {
  default = "eu-west-1"
}

variable "env" {}

variable "lambda_env_variables" {
  type = "map"
}
