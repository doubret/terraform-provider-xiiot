package xiiot

import (
	api_client "github.com/doubret/terraform-provider-xiiot/client"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/doubret/terraform-provider-xiiot/xiiot/resources"
	"github.com/doubret/terraform-provider-xiiot/xiiot/utils"
	swagger_client "github.com/go-openapi/runtime/client"
	swagger_strfmt "github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:        providerSchema(),
		ResourcesMap:  providerResources(),
		ConfigureFunc: providerConfigure,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "User to login within the XiIoT REST api",
			DefaultFunc: schema.EnvDefaultFunc("XI_USER", nil),
		},
		"password": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Password to login within the XiIoT REST api",
			DefaultFunc: schema.EnvDefaultFunc("XI_PASSWORD", nil),
		},
		"endpoint": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The hostname of the XiIoT REST api",
			DefaultFunc: schema.EnvDefaultFunc("XI_ENDPOINT", "iot.nutanix.com"),
		},
	}
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"xiiot_application":       resources.XiIoTApplication(),
		"xiiot_category":          resources.XiIoTCategory(),
		"xiiot_cloudcreds_aws":    resources.XiIoTCloudCredsAws(),
		"xiiot_cloudcreds_gcp":    resources.XiIoTCloudCredsGcp(),
		"xiiot_containerregistry": resources.XiIoTContainerRegistry(),
		"xiiot_dockerprofile":     resources.XiIoTDockerProfile(),
		"xiiot_edge":              resources.XiIoTEdge(),
		"xiiot_project":           resources.XiIoTProject(),
		"xiiot_sensor":            resources.XiIoTSensor(),
		"xiiot_user":              resources.XiIoTUser(),
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	endpoint := d.Get("endpoint").(string)

	client := api_client.New(swagger_client.New(endpoint, "/v1", []string{"https"}), swagger_strfmt.Default)
	bearer := utils.GetBearerToken(client, username, password)

	config := configuration.Configuration{
		client,
		bearer,
	}

	return config, nil
}
