resource "aws_s3_bucket_object" "index" {
  bucket       = "${aws_s3_bucket.web_frontend.id}"
  key          = "index.html"
  source       = "${var.web_path}/index.html"
  etag         = "${md5(file("${var.web_path}/index.html"))}"
  content_type = "text/html"
}

resource "aws_s3_bucket_object" "web_js" {
  bucket        = "${aws_s3_bucket.web_frontend.id}"
  key           = "${var.web_js}"
  source        = "${var.web_path}/${var.web_js}"
  etag          = "${md5(file("${var.web_path}/${var.web_js}"))}"
  cache_control = "max-age=15552000"                              // 180 days
  content_type  = "text/javascript"
}

resource "aws_s3_bucket_object" "web_css" {
  bucket        = "${aws_s3_bucket.web_frontend.id}"
  key           = "${var.web_css}"
  source        = "${var.web_path}/${var.web_css}"
  etag          = "${md5(file("${var.web_path}/${var.web_css}"))}"
  cache_control = "max-age=15552000"                               // 180 days
  content_type  = "text/css"
}
