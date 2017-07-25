variable "region" {
  default = "eu-west-1"
}

variable "db_capacity" {
  default = "3"
}

variable "env" {}

variable "name" {
  default = "nettaton"
}

variable "lambda_env_variables" {
  type = "map"
}
