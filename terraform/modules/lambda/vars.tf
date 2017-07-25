# lambda/vars.tf

variable "region" {}

variable "account_id" {}

variable "function_name" {}

variable "code_zip_path" {}

variable "role_name" {}

variable "api_id" {}

variable "resource_id" {}

variable "resource_path" {}

variable "http_method" {}

variable "lambda_env_variables" {
  type = "map"
}
