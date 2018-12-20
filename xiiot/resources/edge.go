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
		Labels:          utils.Convert_set_to_categoryinfo_array(d.Get("label").(*schema.Set)),
	}

	return &resource
}

func setEdge(d *schema.ResourceData, resource *api_models.Edge) {
	d.Set("name", resource.Name)
	d.Set("description", resource.Description)
	d.Set("edge_devices", resource.EdgeDevices)
	d.Set("gateway", resource.Gateway)
	d.Set("ip_address", resource.IPAddress)
	d.Set("serial_number", resource.SerialNumber)
	d.Set("storage_capacity", resource.StorageCapacity)
	d.Set("storage_usage", resource.StorageUsage)
	d.Set("subnet", resource.Subnet)
	d.Set("connected", resource.Connected)
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
		return err
	}

	setEdge(d, resource.Payload)

	return nil
}

func updateEdge(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.EdgeUpdateV2(api_operations.NewEdgeUpdateV2Params().WithBody(getEdge(d)), config.Auth)

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
