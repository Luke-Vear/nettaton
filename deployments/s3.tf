resource "aws_s3_bucket" "artefact" {
  bucket = "${var.name}-${var.env}-artefact"
  acl    = "private"
}
