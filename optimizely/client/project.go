package client

import (
	"encoding/json"
	"fmt"

	"github.com/pffreitas/optimizely-terraform-provider/optimizely/project"
)

func (c OptimizelyClient) GetProject(projectId string) (project.Project, error) {
	respBody, err := c.sendHttpRequest("GET", fmt.Sprintf("v2/projects/%s", projectId), nil)
	if err != nil {
		return project.Project{}, err
	}

	var projectResp project.Project
	json.Unmarshal(respBody, &projectResp)

	return projectResp, nil
}
