package builds

import "time"

type BuildOverview struct {
	Builds []struct {
		BuildTriggerID string    `json:"buildTriggerId"`
		FinishTime     time.Time `json:"finishTime"`
		ID             string    `json:"id"`
		LogsBucket     string    `json:"logsBucket"`
		ProjectID      string    `json:"projectId"`
		QueueTTL       string    `json:"queueTtl"`
		Results        struct {
			BuildStepImages  []string `json:"buildStepImages"`
			BuildStepOutputs []string `json:"buildStepOutputs"`
		} `json:"results"`
		Source struct {
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
