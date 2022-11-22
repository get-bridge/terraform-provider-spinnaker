package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func stageEnabledResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"expression": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
