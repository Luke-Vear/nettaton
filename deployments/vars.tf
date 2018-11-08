# general
variable "name" {
  default = "nettaton"
}

variable "env" {}

variable "region" {
  default = "eu-west-1"
}

variable "cmd_path" {
  default = "../cmd"
}

# dynamo
variable "db_capacity" {
  default = 3
}

# r53
variable "r53_zone_id" {}

# web
variable "web_path" {
  default = "../web/build"
}

variable "web_js" {}
