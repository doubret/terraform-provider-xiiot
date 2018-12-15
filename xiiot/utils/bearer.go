package utils

import (
	api_client "github.com/doubret/terraform-provider-xiiot/xiiot/client/client"
	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	swagger_runtime "github.com/go-openapi/runtime"
	swagger_client "github.com/go-openapi/runtime/client"
	"log"
)

func GetBearerToken(client *api_client.XiClient, user string, password string) swagger_runtime.ClientAuthInfoWriter {
	credential := &api_models.Credential{
		Email:    &user,
		Password: &password,
	}

	resp, err := client.Operations.LoginCall(api_operations.NewLoginCallParams().WithRequest(credential))

	if err != nil {
		log.Fatal(err)
	}

	return swagger_client.BearerToken(*resp.Payload.Token)
}
