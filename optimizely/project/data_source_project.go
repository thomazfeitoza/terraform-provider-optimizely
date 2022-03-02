package project

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Project struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func DataSourceProject() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIngredientsRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceIngredientsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(ProjectClient)
	project, err := client.GetProject(d.Get("id").(string))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Failed to get Optimizely project: %+v", err),
		})

		return diags
	}

	d.SetId(strconv.FormatInt(project.ID, 10))
	d.Set("name", project.Name)

	return diags
}
