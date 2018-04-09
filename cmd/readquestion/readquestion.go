package main

import (
	"github.com/Luke-Vear/nettaton/internal/nettaton"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(nettaton.ReadQuestion)
}
