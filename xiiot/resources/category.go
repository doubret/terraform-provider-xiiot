package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func XiIoTCategory() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_category,
		Read:          read_category,
		Update:        nil,
		Delete:        delete_category,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"purpose": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"values": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_category(d *schema.ResourceData) nitro.Category {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Category{
		Id:      d.Get("id").(string),
		Name:    d.Get("name").(string),
		Purpose: d.Get("purpose").(string),
		Values:  utils.Convert_set_to_string_array(d.Get("values").(*schema.Set)),
	}

	return resource
}

func set_category(d *schema.ResourceData, resource *nitro.Category) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("id", resource.Id)
	d.Set("name", resource.Name)
	d.Set("purpose", resource.Purpose)
	d.Set("values", resource.Values)

	var key []string

	key = append(key, resource.Id)
	d.SetId(strings.Join(key, "-"))
}

func get_category_key(d *schema.ResourceData) nitro.CategoryKey {

	key := nitro.CategoryKey{
		d.Get("id").(string),
	}
	return key
}

func create_category(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  xiiot-provider: In create_category")

	client := meta.(*nitro.NitroClient)

	resource := get_category(d)
	key := resource.ToKey()

	exists, err := client.ExistsCategory(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCategory(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_category(d, resource)
	} else {
		err := client.AddCategory(get_category(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCategory(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_category(d, resource)
	}

	return nil
}

func read_category(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] xiiot-provider:  In read_category")

	client := meta.(*nitro.NitroClient)

	resource := get_category(d)
	key := resource.ToKey()

	exists, err := client.ExistsCategory(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCategory(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_category(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_category(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  xiiot-provider: In delete_category")

	client := meta.(*nitro.NitroClient)

	resource := get_category(d)
	key := resource.ToKey()

	exists, err := client.ExistsCategory(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCategory(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
