package utils

import (
	"github.com/hashicorp/terraform/helper/schema"
	"strings"
)

func Id_to_state(val interface{}) string {
	return strings.ToLower(val.(string))
}

func Compare_ids(k, old, new string, d *schema.ResourceData) bool {
	if strings.ToLower(old) == strings.ToLower(new) {
		return true
	}
	return false
}
