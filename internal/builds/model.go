package builds

import "time"

// BuildOverview - A high level view of a list of builds, for listing builds with minimal details
type BuildOverview struct {
	Builds []struct {
		BuildTriggerID string    `json:"buildTriggerId"`
		FinishTime     time.Time `json:"finishTime"`
		ID             string    `json:"id"`
		LogsBucket     string    `json:"logsBucket"`
		ProjectID      string    `json:"projectId"`
		QueueTTL       string    `json:"queueTtl"`
		Source         struct {
			RepoSource struct {
				CommitSha string `json:"commitSha"`
				ProjectID string `json:"projectId"`
				RepoName  string `json:"repoName"`
			} `json:"repoSource"`
		} `json:"source"`
		StartTime time.Time `json:"startTime"`
		Status    string    `json:"status"`
		Tags      []string  `json:"tags"`
	} `json:"builds"`
}
