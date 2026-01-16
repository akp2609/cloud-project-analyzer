package writer

import (
	"os"
	"encoding/json"
	"path/filepath"

	"github.com/amanpandey1910/cloud-project-analyzer/cost-processor/internal/model"
)

func WriteAggredatedJSON(
	baseDir string,
	bucket string,
	objectPath string,
	data []model.AggregatedCost,
) (string, error) {

	fullPath := filepath.Join(baseDir, bucket, objectPath)

	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return "", err
	}

	f, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")

	if err := enc.Encode(data); err != nil{
		return "", err
	}

	return fullPath, nil
}