package main

import (
	"encoding/json"
	"fmt"
	"strings"

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

	backendConfig := backendconfigv1beta1.GetOpenAPIDefinitions(ref)
	frontendConfig := frontendconfigv1beta1.GetOpenAPIDefinitions(ref)

	for k, v := range backendConfig {
		swagger.Definitions[slashToDot(k)] = v.Schema
	}

	for k, v := range frontendConfig {
		swagger.Definitions[slashToDot(k)] = v.Schema
	}

	bs, err := json.MarshalIndent(swagger, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}

func slashToDot(in string) (out string) {
	return strings.ReplaceAll(in, "/", ".")
}

func ref(uri string) spec.Ref {
	return spec.MustCreateRef("#/definitions/" + slashToDot(uri))
}
