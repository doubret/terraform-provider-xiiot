package resources

import (
	"log"

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
				Required: true,
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
			"users": &schema.Schema{
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
			"edge_selectors": &schema.Schema{
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

	// "users": &schema.Schema{
	// 	Type: schema.TypeSet,
	// 	Elem: &schema.Resource{
	// 		Schema: map[string]*schema.Schema{
	// 			"user_id": {
	// 				Type:     schema.TypeString,
	// 				Required: true,
	// 			},
	// 			"role": {
	// 				Type:     schema.TypeString,
	// 				Required: true,
	// 			},
	// 		},
	// 	},
	// 	Required: true,
	// },
	// "edge_selectors": &schema.Schema{
	// 	Type: schema.TypeSet,
	// 	Elem: &schema.Resource{
	// 		Schema: map[string]*schema.Schema{
	// 			"id": {
	// 				Type:     schema.TypeString,
	// 				Required: true,
	// 			},
	// 			"value": {
	// 				Type:     schema.TypeString,
	// 				Required: true,
	// 			},
	// 		},
	// 	},
	// 	Optional: true,
	// },

	resource := api_models.Project{
		ID:                 d.Id(),
		Name:               &name,
		Description:        &description,
		CloudCredentialIds: utils.Convert_set_to_string_array(d.Get("cloud_credential_ids").(*schema.Set)),
		DockerProfileIds:   utils.Convert_set_to_string_array(d.Get("docker_profile_ids").(*schema.Set)),
		EdgeIds:            utils.Convert_set_to_string_array(d.Get("edge_ids").(*schema.Set)),
		EdgeSelectorType:   &edgeSelectorType,
		// Users
		// EdgeSelectors
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
	// Users
	// EdgeSelectors
	d.SetId(resource.ID)
}

func createProject(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In createProject")

	config := meta.(configuration.Configuration)

	model := getProject(d)
	_, err := config.Client.Operations.ProjectCreate(api_operations.NewProjectCreateParams().WithDoc(model), config.Auth)

	if err != nil {
		log.Print("Failed to create resource : ", err)

		return err
	}

	setProject(d, model)

	return nil
}

func readProject(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In readProject")

	id := d.Id()

	config := meta.(configuration.Configuration)

	model, err := config.Client.Operations.ProjectGet(api_operations.NewProjectGetParams().WithProjectID(id), config.Auth)

	if err != nil {
		log.Print("Failed to read resource : ", err)

		return err
	}

	setProject(d, model.Payload)

	return nil
}

func updateProject(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In updateProject")

	config := meta.(configuration.Configuration)

	model := getProject(d)
	_, err := config.Client.Operations.ProjectUpdateV2(api_operations.NewProjectUpdateV2Params().WithDoc(model), config.Auth)

	if err != nil {
		log.Print("Failed to update resource : ", err)

		return err
	}

	setProject(d, model)

	return nil
}

func deleteProject(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] xiiot-provider: In deleteProject")

	id := d.Id()

	config := meta.(configuration.Configuration)

	_, err := config.Client.Operations.ProjectDelete(api_operations.NewProjectDeleteParams().WithID(id), config.Auth)

	if err != nil {
		log.Print("Failed to delete resource : ", err)

		return err
	}

	d.SetId("")

	return nil
}
