package project

import (
	"context"

	"google.golang.org/api/cloudresourcemanager/v2"
)

// List all Google projects we can see with the current credentials
func ListAll() (projects []Project, err error) {

	// Get all projects

	ctx := context.Background()
	cloudresourcemanagerService, err := cloudresourcemanager.NewService(ctx)

	// Validate the list
	if err != nil {
		return err
	}

	// Convert google API call results into a slice of type Project
	allProjects, err := generateProjectsSlice()
	if err != nil {
		return nil, err
	}
}

// Take the results from the google call and turn them into a slice of Project
func generateProjectsSlice() {

}

// Project -  A Google project
type Project struct {
	Name string
	ID   string
}
