package metrics

type PubSubMessage struct {
	Message struct {
		Data string `json:"data"`
	} `json:"message"`
}

type MetricsEvent struct {
	ProjectID string `json:"project_id"`
}
