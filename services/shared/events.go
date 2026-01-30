package events

type EventMeta struct {
	PlatformProjectID string `json:"platform_project_id"`
	HookID            string `json:"hook_id"`
	ExternalProjectID string `json:"external_project_id"`
	Cloud             string `json:"cloud"`
	SignalType        string `json:"signal_type"`
	Timestamp         string `json:"timestamp"`
}

type UploadEvent struct {
	Bucket string    `json:"bucket"`
	Path   string    `json:"path"`
	Meta   EventMeta `json:"meta"`
}
