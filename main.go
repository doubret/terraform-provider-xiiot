package main

import (
	"github.com/doubret/terraform-provider-xiiot/xiiot"
	apiclient "github.com/doubret/terraform-provider-xiiot/xiiot/client/client"
	"github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	apimodels "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/plugin"
	"log"
	"os"
)

func login(client *apiclient.XiClient, user string, password string) runtime.ClientAuthInfoWriter {
	credential := &apimodels.Credential{
		Email:    &user,
		Password: &password,
	}

	resp, err := client.Operations.LoginCall(operations.NewLoginCallParams().WithRequest(credential))

	if err != nil {
		log.Fatal(err)
	}

	return httptransport.BearerToken(*resp.Payload.Token)
}

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: xiiot.Provider,
	}

	client := apiclient.New(httptransport.New("iot.nutanix.com", "/v1", []string{"https"}), strfmt.Default)

	bearerTokenAuth := login(client, os.Getenv("XI_USER"), os.Getenv("XI_PASSWORD"))

	resp, err := client.Operations.ApplicationList(operations.NewApplicationListParams(), bearerTokenAuth)

	if err != nil {
		log.Fatal(err)
	}

	plugin.Serve(&opts)
}
