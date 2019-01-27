package utils

import (
	"github.com/hashicorp/terraform/helper/schema"
	"strings"
)

func Compare_ids(k, old, new string, d *schema.ResourceData) bool {
	if strings.ToLower(old) == strings.ToLower(new) {
		return true
	}
	return false
}
