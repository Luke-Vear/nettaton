# dynamodb/main.tf

resource "aws_dynamodb_table" "table" {
  name           = "${var.db_name}"
  read_capacity  = "${var.db_capacity}"
  write_capacity = "${var.db_capacity}"
  hash_key       = "id"

  attribute {
    name = "id"
    type = "S"
  }
}

output "db_arn" {
  value = "${aws_dynamodb_table.table.arn}"
}
