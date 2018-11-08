output "website_endpoint" {
    value = "${aws_s3_bucket.web_frontend.website_endpoint}"
}

resource "aws_s3_bucket" "web_frontend" {
  bucket = "${var.env == "prod" ? "nettaton.com" : "${var.env}.nettaton.com"}"
  acl    = "public-read"

  website {
    index_document = "index.html"
  }
}

resource "aws_s3_bucket_object" "index" {
  bucket = "${aws_s3_bucket.web_frontend.id}"
  key    = "index.html"
  source = "${var.web_path}/index.html"
  etag   = "${md5(file("${var.web_path}/index.html"))}"
  content_type = "text/html"
}

resource "aws_s3_bucket_object" "web_js" {
  bucket = "${aws_s3_bucket.web_frontend.id}"
  key    = "static/js/${var.web_js}"
  source = "${var.web_path}/static/js/${var.web_js}"
  etag   = "${md5(file("${var.web_path}/static/js/${var.web_js}"))}"
  content_type = "text/javascript"
}


resource "aws_s3_bucket_policy" "allow_public_s3" {
  bucket = "${aws_s3_bucket.web_frontend.id}"
  policy = "${data.aws_iam_policy_document.allow_public_s3.json}"
}

data "aws_iam_policy_document" "allow_public_s3" {
  statement {
    actions = ["s3:GetObject"]

    principals {
      type        = "*"
      identifiers = ["*"]
    }

    resources = [
      //"${aws_s3_bucket.web_frontend.arn}",
      "${aws_s3_bucket.web_frontend.arn}/*",
    ]
  }
}

# resource "aws_s3_bucket_policy" "allow_cloudfront_s3" {
#   bucket = "${aws_s3_bucket.web_frontend.id}"
#   policy = "${data.aws_iam_policy_document.bucket_policy.json}"
# }

# data "aws_iam_policy_document" "allow_cloudfront_s3" {
#   statement {
#     actions = ["s3:GetObject"]

#     principals {
#       type        = "AWS"
#       identifiers = ["${aws_cloudfront_origin_access_identity.origin_access_identity.iam_arn}"]
#     }

#     resources = [
#       "${aws_s3_bucket.web_frontend.arn}",
#       "${aws_s3_bucket.web_frontend.arn}/*",
#     ]
#   }

#   statement {
#     actions = ["s3:ListBucket"]

#     principals {
#       type        = "AWS"
#       identifiers = ["${aws_cloudfront_origin_access_identity.origin_access_identity.iam_arn}"]
#     }

#     resources = [
#       "${aws_s3_bucket.web_frontend.arn}",
#       "${aws_s3_bucket.web_frontend.arn}/*",
#     ]
#   }
#}