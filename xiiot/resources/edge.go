package resources

import (
	"log"

	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
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
			"edge_devices": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
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
			"storage_capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"storage_usage": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"connected": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"labels": &schema.Schema{
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
	edgeDevices := float64(d.Get("edge_devices").(int))
	gateway := d.Get("gateway").(string)
	ipAddress := d.Get("ip_address").(string)
	serialNumber := d.Get("serial_number").(string)
	storageCapacity := float64(d.Get("storage_capacity").(int))
	storageUsage := float64(d.Get("storage_usage").(int))
	subnet := d.Get("subnet").(string)
	connected := d.Get("connected").(bool)

	resource := api_models.Edge{
		ID:              d.Id(),
		Name:            &name,
		Description:     description,
		EdgeDevices:     &edgeDevices,
		Gateway:         &gateway,
		IPAddress:       &ipAddress,
		SerialNumber:    &serialNumber,
		StorageCapacity: &storageCapacity,
		StorageUsage:    &storageUsage,
		Subnet:          &subnet,
		Connected:       connected,
		Labels:          utils.Convert_set_to_categoryinfo_array(d.Get("labels").(*schema.Set)),
	}

	return &resource
}

// func setEdge(d *schema.ResourceData, resource *api_models.Edge) {
// 	d.Set("name", resource.Name)
// 	d.Set("description", resource.Description)
// 	d.Set("cloud_credential_ids", resource.CloudCredentialIds)
// 	d.Set("docker_profile_ids", resource.DockerProfileIds)
// 	d.Set("edge_ids", resource.EdgeIds)
// 	d.Set("edge_selector_type", resource.EdgeSelectorType)
// 	// Users
// 	// EdgeSelectors
// 	d.SetId(resource.ID)
// }

func createEdge(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createEdge")

	config := meta.(configuration.Configuration)

	model := getEdge(d)
	result, err := config.Client.Operations.EdgeCreate(api_operations.NewEdgeCreateParams().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	d.SetId(*result.Payload.ID)

	return nil
}

func readEdge(d *schema.ResourceData, meta interface{}) error {
	// 	log.Printf("[DEBUG] xiiot-provider: In readEdge")

	// 	id := d.Id()

	// 	config := meta.(configuration.Configuration)

	// 	model, err := config.Client.Operations.EdgeGet(api_operations.NewEdgeGetParams().WithEdgeID(id), config.Auth)

	// 	if err != nil {
	// 		log.Print("Failed to read resource : ", err)

	// 		return err
	// 	}

	// 	setEdge(d, model.Payload)

	return nil
}

func updateEdge(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateEdge")

	config := meta.(configuration.Configuration)

	model := getEdge(d)
	result, err := config.Client.Operations.EdgeUpdateV2(api_operations.NewEdgeUpdateV2Params().WithBody(model), config.Auth)

	if err != nil {
		log.Print("Failed to update resource : ", err)

		return err
	}

	d.SetId(*result.Payload.ID)

	return nil
}

func deleteEdge(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In deleteEdge")

	id := d.Id()

	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.EdgeDelete(api_operations.NewEdgeDeleteParams().WithEdgeID(id), config.Auth)

	if err != nil {
		log.Print("Failed to delete resource : ", err)

		return err
	}

	d.SetId("")

	return nil
}
