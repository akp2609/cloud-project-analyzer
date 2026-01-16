package writer

import (
	"context"
	"encoding/json"

	"github.com/amanpandey1910/cloud-project-analyzer/cost-processor/internal/model"
	"cloud.google.com/go/storage"
)

func WriteAggredatedJSONToGCS(
	ctx context.Context,
	client *storage.Client,
	bucket string,
	objectPath string,
	data []model.AggregatedCost,
) error {

	w := client.Bucket(bucket).Object(objectPath).NewWriter(ctx)
	w.ContentType = "application/json"

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	if err := enc.Encode(data); err != nil  {
		_ = w.Close()
		return err
	}

	if err := w.Close(); err != nil {
		return err
	}

	return nil

}