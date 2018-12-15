package configuration

import (
	api_client "github.com/doubret/terraform-provider-xiiot/xiiot/client/client"
	swagger_runtime "github.com/go-openapi/runtime"
)

type Configuration struct {
	Client *api_client.XiClient
	Auth   swagger_runtime.ClientAuthInfoWriter
}
