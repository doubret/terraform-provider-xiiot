package xiiot

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	//	"github.com/doubret/terraform-provider-xiiot/xiiot/resources"
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
			Description: "Username to login to the XiIoT",
			DefaultFunc: schema.EnvDefaultFunc("XI_LOGIN", nil),
		},
		"password": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Password to login to the XiIoT",
			DefaultFunc: schema.EnvDefaultFunc("XI_PASSWORD", nil),
		},
		"endpoint": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The URL to the XiIoT API",
			DefaultFunc: schema.EnvDefaultFunc("XI_URL", nil),
		},
	}
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		//		"xiiot_category": resources.XiIoTCategory(),
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	//  TODO : xi iot client
	// c := nitro.NewNitroClient(d.Get("endpoint").(string), d.Get("username").(string), d.Get("password").(string))

	return nil, nil
}
