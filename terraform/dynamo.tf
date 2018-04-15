resource "aws_dynamodb_table" "questionstore" {
  name           = "${var.name}-${var.env}-questionstore"
  read_capacity  = "${var.db_capacity}"
  write_capacity = "${var.db_capacity}"
  hash_key       = "id"

  attribute {
    name = "id"
    type = "S"
  }
}
