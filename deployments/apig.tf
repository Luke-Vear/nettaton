resource "aws_api_gateway_rest_api" "api" {
  name = "${var.name}-${var.env}"
}

resource "aws_api_gateway_resource" "question" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  parent_id   = "${aws_api_gateway_rest_api.api.root_resource_id}"
  path_part   = "question"
}

resource "aws_api_gateway_resource" "question_id" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  parent_id   = "${aws_api_gateway_resource.question.id}"
  path_part   = "{id}"
}

resource "aws_api_gateway_resource" "question_id_answer" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  parent_id   = "${aws_api_gateway_resource.question_id.id}"
  path_part   = "answer"
}

resource "aws_api_gateway_deployment" "quiz_api" {
  rest_api_id = "${aws_api_gateway_rest_api.api.id}"
  stage_name  = "monostage"
}

resource "aws_api_gateway_domain_name" "quiz_api" {
  domain_name     = "${aws_acm_certificate.quiz_api.domain_name}"
  certificate_arn = "${aws_acm_certificate.quiz_api.arn}"
}

resource "aws_api_gateway_base_path_mapping" "quiz_api" {
  api_id      = "${aws_api_gateway_rest_api.api.id}"
  stage_name  = "${aws_api_gateway_deployment.quiz_api.stage_name}"
  domain_name = "${aws_api_gateway_domain_name.quiz_api.domain_name}"
}

resource "aws_acm_certificate" "quiz_api" {
  provider          = "aws.us-east-1"
  domain_name       = "${var.env == "prod" ? "api.nettaton.com" : "api.${var.env}.nettaton.com"}"
  validation_method = "DNS"
}

resource "aws_route53_record" "quiz_api" {
  zone_id = "${var.r53_zone_id}"
  name    = "${aws_api_gateway_domain_name.quiz_api.domain_name}"
  type    = "A"

  alias {
    name                   = "${aws_api_gateway_domain_name.quiz_api.cloudfront_domain_name}"
    zone_id                = "${aws_api_gateway_domain_name.quiz_api.cloudfront_zone_id}"
    evaluate_target_health = false
  }
}

resource "aws_route53_record" "cert_validation" {
  provider = "aws.us-east-1"
  name     = "${aws_acm_certificate.quiz_api.domain_validation_options.0.resource_record_name}"
  type     = "${aws_acm_certificate.quiz_api.domain_validation_options.0.resource_record_type}"
  zone_id  = "${var.r53_zone_id}"
  records  = ["${aws_acm_certificate.quiz_api.domain_validation_options.0.resource_record_value}"]
  ttl      = 60
}

resource "aws_acm_certificate_validation" "cert" {
  provider                = "aws.us-east-1"
  certificate_arn         = "${aws_acm_certificate.quiz_api.arn}"
  validation_record_fqdns = ["${aws_route53_record.cert_validation.fqdn}"]
}
