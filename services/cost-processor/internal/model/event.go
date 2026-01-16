package model

type UploadEvent struct {
	Bucket string `json:"bucket"`
	Path string `json:"path"`
	Tenant string `json:"tenant"`
}