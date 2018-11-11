locals {
  endpoint = "${var.env == "prod" ? "nettaton.com" : "${var.env}.nettaton.com"}"
}

resource "aws_route53_record" "web" {
  zone_id = "${var.r53_zone_id}"
  name    = "${local.endpoint}"
  type    = "A"

  alias {
    name                   = "${aws_cloudfront_distribution.web.domain_name}"
    zone_id                = "${aws_cloudfront_distribution.web.hosted_zone_id}"
    evaluate_target_health = false
  }
}

resource "aws_s3_bucket" "web_frontend" {
  bucket = "${local.endpoint}"
  acl    = "public-read"

  website {
    index_document = "index.html"
  }
}

resource "aws_acm_certificate" "web" {
  provider          = "aws.us-east-1"
  domain_name       = "${local.endpoint}"
  validation_method = "DNS"
}

resource "aws_route53_record" "web_cert_validation" {
  provider = "aws.us-east-1"
  name     = "${aws_acm_certificate.web.domain_validation_options.0.resource_record_name}"
  type     = "${aws_acm_certificate.web.domain_validation_options.0.resource_record_type}"
  zone_id  = "${var.r53_zone_id}"
  records  = ["${aws_acm_certificate.web.domain_validation_options.0.resource_record_value}"]
  ttl      = 60
}

resource "aws_acm_certificate_validation" "web_cert_validation" {
  provider                = "aws.us-east-1"
  certificate_arn         = "${aws_acm_certificate.web.arn}"
  validation_record_fqdns = ["${aws_route53_record.web_cert_validation.fqdn}"]
}

resource "aws_cloudfront_distribution" "web" {
  aliases = ["${local.endpoint}"]

  default_cache_behavior {
    allowed_methods = ["GET", "HEAD"]
    cached_methods  = ["GET", "HEAD"]
    compress        = true

    forwarded_values {
      cookies {
        forward = "none"
      }

      query_string = false
    }

    target_origin_id       = "${local.endpoint}"
    viewer_protocol_policy = "redirect-to-https"
  }

  default_root_object = "${aws_s3_bucket_object.index.id}"
  enabled             = true

  origin {
    s3_origin_config {
      origin_access_identity = "${aws_cloudfront_origin_access_identity.web_cloudfront.cloudfront_access_identity_path}"
    }

    domain_name = "${aws_s3_bucket.web_frontend.bucket_domain_name}" //
    origin_id   = "${local.endpoint}"
  }

  price_class = "PriceClass_100"

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    acm_certificate_arn = "${aws_acm_certificate.web.arn}"

    minimum_protocol_version = "TLSv1.2_2018"
    ssl_support_method       = "sni-only"
  }
}

resource "aws_cloudfront_origin_access_identity" "web_cloudfront" {}

resource "aws_s3_bucket_policy" "web_cloudfront" {
  bucket = "${aws_s3_bucket.web_frontend.id}"
  policy = "${data.aws_iam_policy_document.web_cloudfront.json}"
}

data "aws_iam_policy_document" "web_cloudfront" {
  statement {
    sid     = ""
    actions = ["s3:GetObject"]

    principals {
      type        = "AWS"
      identifiers = ["${aws_cloudfront_origin_access_identity.web_cloudfront.iam_arn}"]
    }

    resources = ["${aws_s3_bucket.web_frontend.arn}/*"]
  }

  statement {
    sid     = ""
    actions = ["s3:ListBucket"]

    principals {
      type        = "AWS"
      identifiers = ["${aws_cloudfront_origin_access_identity.web_cloudfront.iam_arn}"]
    }

    resources = ["${aws_s3_bucket.web_frontend.arn}"]
  }
}
