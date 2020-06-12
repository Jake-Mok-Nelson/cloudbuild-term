package builds

import (
	"context"

	"google.golang.org/api/cloudbuild/v1"
)

// FetchBuilds - Fetches all builds from a given project
func FetchBuilds(projectID string, limit int64) (builds []byte, err error) {
	ctx := context.Background()
	cloudbuildService, err := cloudbuild.NewService(ctx)
	if err != nil {
		return nil, err
	}

	// Get builds for a given project
	buildsList, err := cloudbuildService.Projects.Builds.List(projectID).
		Fields(
			"builds/buildTriggerId",
			"builds/logsBucket",
			"builds/startTime",
			"builds/finishTime",
			"builds/id",
			"builds/projectId",
			"builds/results",
			"builds/source",
			"builds/status",
			"builds/queueTtl",
			"builds/statusDetail",
			"builds/tags").
		PageSize(limit).
		Do()

	if err != nil {
		return nil, err
	}

	// Convert slice of cloudbuild.build type to JSON
	d, err := buildsList.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return d, nil
}
