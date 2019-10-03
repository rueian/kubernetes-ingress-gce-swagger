package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/spec"
	backendconfigv1beta1 "k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1"
	frontendconfigv1beta1 "k8s.io/ingress-gce/pkg/apis/frontendconfig/v1beta1"
)

func main() {
	swagger := spec.Swagger{
		SwaggerProps: spec.SwaggerProps{
			Definitions: map[string]spec.Schema{},
		},
	}

	backendConfig := backendconfigv1beta1.GetOpenAPIDefinitions(spec.MustCreateRef)
	frontendConfig := frontendconfigv1beta1.GetOpenAPIDefinitions(spec.MustCreateRef)

	for k, v := range backendConfig {
		swagger.Definitions[k] = v.Schema
	}

	for k, v := range frontendConfig {
		swagger.Definitions[k] = v.Schema
	}

	bs, err := json.MarshalIndent(swagger, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}
