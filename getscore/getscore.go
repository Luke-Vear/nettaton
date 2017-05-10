package main

import "github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"

// Handle is invoked by the shim.
func Handle(evt interface{}, ctx *runtime.Context) (string, error) {
	return "Score!", nil
}

func main() {}
