package main

import (
	"github.com/Luke-Vear/nettaton/internal"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	config := internal.LoadConfig()
	resolver := internal.NewDependencyResolver(config)

	nettatonNexus := resolver.ResolveNettatonNexus()

	lambda.Start(nettatonNexus.ReadQuestion)
}
