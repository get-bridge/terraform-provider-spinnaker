package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func canaryConfigClassifierResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"group_weights": {
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}
}
