package projects

// // List all Google projects we can see with the current credentials
// func FetchFromGCP() (projects []string, err error) {

// 	// Get all projects
// 	ctx := context.Background()

// 	cloudresourcemanagerService, err := cloudresourcemanager.NewService(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var projectsFromApiCall []string
// 	req := cloudresourcemanagerService.Projects.List()
// 	if err := req.Pages(ctx, func(page *cloudresourcemanager.ListProjectsResponse) error {
// 		for _, project := range page.Projects {
// 			projectsFromApiCall = append(projectsFromApiCall, project.ProjectId)
// 		}
// 		return nil
// 	}); err != nil {
// 		return nil, err
// 	}

// 	// Convert the google projects into a simple slice of strings
// 	return projectsFromApiCall, nil
// }

type Project struct {
	Name string
	ID   string
}
