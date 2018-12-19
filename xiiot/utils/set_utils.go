package utils

import (
	api_models "github.com/doubret/terraform-provider-xiiot/xiiot/client/models"
	"github.com/hashicorp/terraform/helper/schema"
)

func Convert_set_to_string_array(set *schema.Set) []string {
	var results []string

	for _, item := range set.List() {
		results = append(results, item.(string))
	}

	return results
}

func Convert_set_to_int_array(set *schema.Set) []int {
	var results []int

	for _, item := range set.List() {
		results = append(results, item.(int))
	}

	return results
}

func Convert_set_to_categoryinfo_array(set *schema.Set) []*api_models.CategoryInfo {
	var results []*api_models.CategoryInfo

	for _, raw := range set.List() {
		item := raw.(map[string]interface{})

		id := item["id"].(string)
		value := item["value"].(string)

		results = append(results, &api_models.CategoryInfo{
			ID:    &id,
			Value: &value,
		})
	}

	return results
}

func Convert_categoryinfo_array_to_set(array []*api_models.CategoryInfo) []interface{} {
	var results []interface{}

	for _, item := range array {
		raw := make(map[string]string)

		raw["id"] = *item.ID
		raw["value"] = *item.Value

		results = append(results, raw)
	}

	return results
}

func Convert_set_to_projectuserinfo_array(set *schema.Set) []*api_models.ProjectUserInfo {
	var results []*api_models.ProjectUserInfo

	for _, raw := range set.List() {
		item := raw.(map[string]interface{})

		userID := item["user_id"].(string)
		role := item["role"].(string)

		results = append(results, &api_models.ProjectUserInfo{
			UserID: &userID,
			Role:   &role,
		})
	}

	return results
}
