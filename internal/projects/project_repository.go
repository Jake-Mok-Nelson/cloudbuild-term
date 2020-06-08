package projects

import (
	"context"

	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
)

/// Fetch all Google projects we can see with the current credentials
func FetchProjects() (reply interface{}, err error) {

	ctx := context.Background()
	cloudresourcemanagerService, err := cloudresourcemanager.NewService(ctx, option.WithScopes(cloudresourcemanager.CloudPlatformReadOnlyScope))
	if err != nil {
		return nil, err
	}
	projectsList, err := cloudresourcemanagerService.Projects.List().Do()
	if err != nil {
		return nil, err
	}

	print(projectsList.Projects)
	projectSlice := Projects{}

	for _, project := range projectsList.Projects {
		p := ProjectItem{
			ID:   project.ProjectId,
			Name: project.Name,
		}
		projectSlice.AddItem(p)
	}

	return projectSlice, err
}
