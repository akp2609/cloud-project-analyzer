package logs

type LogEntry struct {
	Severity  string `json:"severity"`
	Timestamp string `json:"timestamp"`

	Resource struct {
		Labels struct {
			ProjectID string `json:"project_id"`
		} `json:"labels"`
	} `json:"resource"`
}
