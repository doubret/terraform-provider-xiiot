package resources

import (
	api_operations "github.com/doubret/terraform-provider-xiiot/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/doubret/terraform-provider-xiiot/xiiot/utils"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTEdge() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createEdge,
		Read:          readEdge,
		Update:        updateEdge,
		Delete:        deleteEdge,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"label": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Optional: true,
			},
		},
	}
}

func getEdge(d *schema.ResourceData) *api_models.Edge {
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	gateway := d.Get("gateway").(string)
	ipAddress := d.Get("ip_address").(string)
	serialNumber := d.Get("serial_number").(string)
	subnet := d.Get("subnet").(string)

	resource := api_models.Edge{
		ID:           d.Id(),
		Name:         &name,
		Description:  description,
		Gateway:      &gateway,
		IPAddress:    &ipAddress,
		SerialNumber: &serialNumber,
		Subnet:       &subnet,
		Labels:       utils.Convert_set_to_categoryinfo_array(d.Get("label").(*schema.Set)),
	}

	return &resource
}

func setEdge(d *schema.ResourceData, resource *api_models.Edge) {
	d.Set("name", resource.Name)
	d.Set("description", resource.Description)
	d.Set("gateway", resource.Gateway)
	d.Set("ip_address", resource.IPAddress)
	d.Set("serial_number", resource.SerialNumber)
	d.Set("subnet", resource.Subnet)
	d.Set("label", utils.Convert_categoryinfo_array_to_set(resource.Labels))
}

func createEdge(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	result, err := config.Client.Operations.EdgeCreate(api_operations.NewEdgeCreateParams().WithBody(getEdge(d)), config.Auth)

	if err != nil {
		return err
	}

	d.SetId(*result.Payload.ID)

	return readEdge(d, meta)
}

func readEdge(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	resource, err := config.Client.Operations.EdgeGet(api_operations.NewEdgeGetParams().WithEdgeID(d.Id()), config.Auth)

	if err != nil {
		if httperr, ok := err.(*api_operations.EdgeGetDefault); ok {
			if httperr.Code() == 404 {
				d.SetId("")

				return nil
			}
		}
		
		return err
	}

	setEdge(d, resource.Payload)

	return nil
}

func updateEdge(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.EdgeUpdateV2(api_operations.NewEdgeUpdateV2Params().WithID(d.Id()).WithBody(getEdge(d)), config.Auth)

	if err != nil {
		return err
	}

	return readEdge(d, meta)
}

func deleteEdge(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.EdgeDelete(api_operations.NewEdgeDeleteParams().WithEdgeID(d.Id()), config.Auth)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
