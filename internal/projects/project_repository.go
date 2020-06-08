package projects

import (
	"context"

	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
)

// FetchProjects - Fetch all Google projects we can see with the current credentials
func FetchProjects() (allProjects []string, err error) {

	ctx := context.Background()
	cloudresourcemanagerService, err := cloudresourcemanager.NewService(ctx, option.WithScopes(cloudresourcemanager.CloudPlatformReadOnlyScope))
	if err != nil {
		return nil, err
	}
	projectsList, err := cloudresourcemanagerService.Projects.List().Do()
	if err != nil {
		return nil, err
	}

	projSlice := make([]string, len(projectsList.Projects))

	for i, project := range projectsList.Projects {
		projSlice[i] = project.ProjectId
	}

	return projSlice, err
}
