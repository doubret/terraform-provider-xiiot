package resources

import (
	api_operations "github.com/doubret/terraform-provider-xiiot/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/doubret/terraform-provider-xiiot/xiiot/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"strings"
)

func XiIoTApplication() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createApplication,
		Read:          readApplication,
		Update:        updateApplication,
		Delete:        deleteApplication,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"yaml_data": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				DiffSuppressFunc: utils.Compare_ids,
			},
			"edge_ids": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					DiffSuppressFunc: utils.Compare_ids,
				},
				Optional: true,
			},
			"edge_selector": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
							DiffSuppressFunc: utils.Compare_ids,
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

func getApplication(d *schema.ResourceData) *api_models.Application {
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	yamlData := d.Get("yaml_data").(string)
	projectID := d.Get("project_id").(string)

	resource := api_models.Application{
		ID:            d.Id(),
		Name:          &name,
		Description:   description,
		YamlData:      &yamlData,
		ProjectID:     projectID,
		EdgeIds:       utils.Convert_set_to_string_array(d.Get("edge_ids").(*schema.Set)),
		EdgeSelectors: utils.Convert_set_to_categoryinfo_array(d.Get("edge_selector").(*schema.Set)),
	}

	return &resource
}

func setApplication(d *schema.ResourceData, resource *api_models.Application) {
	d.Set("name", resource.Name)
	d.Set("description", resource.Description)
	d.Set("yaml_data", resource.YamlData)
	d.Set("project_id", resource.ProjectID)
	d.Set("edge_ids", resource.EdgeIds)
	d.Set("edge_selector", utils.Convert_categoryinfo_array_to_set(resource.EdgeSelectors))
}

func createApplication(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	result, err := config.Client.Operations.ApplicationCreate(api_operations.NewApplicationCreateParams().WithBody(getApplication(d)), config.Auth)

	if err != nil {
		return err
	}

	d.SetId(*result.Payload.ID)

	return readApplication(d, meta)
}

func readApplication(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	resource, err := config.Client.Operations.ApplicationGet(api_operations.NewApplicationGetParams().WithID(d.Id()), config.Auth)

	if err != nil {
		if httperr, ok := err.(*api_operations.ApplicationGetDefault); ok {
			if httperr.Code() == 404 {
				d.SetId("")

				return nil
			}
		}

		return err
	}

	setApplication(d, resource.Payload)

	return nil
}

func updateApplication(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.ApplicationUpdateV2(api_operations.NewApplicationUpdateV2Params().WithID(d.Id()).WithBody(getApplication(d)), config.Auth)

	if err != nil {
		return err
	}

	return readApplication(d, meta)
}

func deleteApplication(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.ApplicationDelete(api_operations.NewApplicationDeleteParams().WithID(d.Id()), config.Auth)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
