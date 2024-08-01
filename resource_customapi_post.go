package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCustomAPIPost() *schema.Resource {
	return &schema.Resource{
		Create: resourceCustomAPIPostCreate,
		Read:   resourceCustomAPIPostRead,
		Update: resourceCustomAPIPostUpdate,
		Delete: resourceCustomAPIPostDelete,

		Schema: map[string]*schema.Schema{
			"organization_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"var_identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"auth_token": {
				Type:     schema.TypeString,
				Required: true,
			},
			"input_data": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"response": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCustomAPIPostCreate(d *schema.ResourceData, m interface{}) error {
	endpoint := "https://api.external.envtrack.com/postVariablesToEnvironment"
	authToken := d.Get("auth_token").(string)
	orgID := d.Get("organization_id").(string)
	projectID := d.Get("project_id").(string)
	varIdentifier := d.Get("var_identifier").(string)
	environmentID := d.Get("environment_id").(string)

	inputData := d.Get("input_data").(map[string]interface{})

	// transfrom inputData from map[string]interface{} to array of {key: string, value: string}
	inputDataArray := []map[string]string{}
	for k, v := range inputData {
		inputDataArray = append(inputDataArray, map[string]string{
			"name":  k,
			"value": v.(string),
		})
	}

	jsonData, err := json.Marshal(map[string]interface{}{
		"newVars": map[string]interface{}{
			"id":   varIdentifier,
			"vars": inputDataArray,
		},
	})
	if err != nil {
		return fmt.Errorf("error marshaling input data: %s", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf(
		endpoint+"?orgId=%s&prjId=%s&envId=%s",
		orgID, projectID, environmentID,
	), bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-carrier-auth", authToken)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	d.SetId(endpoint)
	d.Set("response", fmt.Sprintf("Successfully posted to %s", endpoint))

	return nil
}

func resourceCustomAPIPostRead(d *schema.ResourceData, m interface{}) error {
	// In this case, we don't need to do anything for the Read operation
	return nil
}

func resourceCustomAPIPostUpdate(d *schema.ResourceData, m interface{}) error {
	// For this example, update is the same as create
	return resourceCustomAPIPostCreate(d, m)
}

func resourceCustomAPIPostDelete(d *schema.ResourceData, m interface{}) error {
	// For this example, we don't need to do anything for delete
	d.SetId("")
	return nil
}
