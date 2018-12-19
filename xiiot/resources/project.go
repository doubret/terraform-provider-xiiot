package resources

import (
	api_operations "github.com/doubret/terraform-provider-xiiot/xiiot/client/client/operations"
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/doubret/terraform-provider-xiiot/xiiot/configuration"
	"github.com/doubret/terraform-provider-xiiot/xiiot/utils"
	"github.com/hashicorp/terraform/helper/schema"
)

func XiIoTProject() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createProject,
		Read:          readProject,
		Update:        updateProject,
		Delete:        deleteProject,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_credential_ids": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"docker_profile_ids": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"user": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"role": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
				Required: true,
			},
			"edge_selector_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"edge_ids": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
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

func getProject(d *schema.ResourceData) *api_models.Project {
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	edgeSelectorType := d.Get("edge_selector_type").(string)

	resource := api_models.Project{
		ID:                 d.Id(),
		Name:               &name,
		Description:        &description,
		CloudCredentialIds: utils.Convert_set_to_string_array(d.Get("cloud_credential_ids").(*schema.Set)),
		DockerProfileIds:   utils.Convert_set_to_string_array(d.Get("docker_profile_ids").(*schema.Set)),
		EdgeIds:            utils.Convert_set_to_string_array(d.Get("edge_ids").(*schema.Set)),
		EdgeSelectorType:   &edgeSelectorType,
		EdgeSelectors:      utils.Convert_set_to_categoryinfo_array(d.Get("edge_selector").(*schema.Set)),
		Users:              utils.Convert_set_to_projectuserinfo_array(d.Get("user").(*schema.Set)),
	}

	return &resource
}

func setProject(d *schema.ResourceData, resource *api_models.Project) {
	d.Set("name", resource.Name)
	d.Set("description", resource.Description)
	d.Set("cloud_credential_ids", resource.CloudCredentialIds)
	d.Set("docker_profile_ids", resource.DockerProfileIds)
	d.Set("edge_ids", resource.EdgeIds)
	d.Set("edge_selector_type", resource.EdgeSelectorType)
	d.Set("edge_selector", utils.Convert_categoryinfo_array_to_set(resource.EdgeSelectors))
	d.Set("user", utils.Convert_projectuserinfo_array_to_set(resource.Users))
}

func createProject(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	result, err := config.Client.Operations.ProjectCreate(api_operations.NewProjectCreateParams().WithDoc(getProject(d)), config.Auth)

	if err != nil {
		return err
	}

	d.SetId(*result.Payload.ID)

	return readProject(d, meta)
}

func readProject(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	resource, err := config.Client.Operations.ProjectGet(api_operations.NewProjectGetParams().WithProjectID(d.Id()), config.Auth)

	if err != nil {
		return err
	}

	setProject(d, resource.Payload)

	return nil
}

func updateProject(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.ProjectUpdateV2(api_operations.NewProjectUpdateV2Params().WithDoc(getProject(d)), config.Auth)

	if err != nil {
		return err
	}

	return readProject(d, meta)
}

func deleteProject(d *schema.ResourceData, meta interface{}) error {
	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.ProjectDelete(api_operations.NewProjectDeleteParams().WithID(d.Id()), config.Auth)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
