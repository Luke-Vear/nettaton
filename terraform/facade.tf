# only in prod need dns for quiz api
# only in prod need dns for gui bucket
# resource "aws_security_group_rule" "rabbitmq_tcp_5672_google" {
#   count = "${var.env == "prod" ? 1 : 0}"
#   type      = "ingress"
#   from_port = 5672
#   to_port   = 5672
#   protocol  = "tcp"
#   cidr_blocks = [
#     "${var.google_vpc_cidr}",
#   ]
#   security_group_id = "${aws_security_group.queue.id}"
# }

