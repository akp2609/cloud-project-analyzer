package main

import (
	"context"
	"fmt"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
)

func verifyProject(ctx context.Context, projectID string) error {
	client, err := resourcemanager.NewProjectsClient(ctx)
	if err != nil {
		return fmt.Errorf("rm client: %w", err)
	}
	defer client.Close()

	req := &resourcemanagerpb.GetProjectRequest{
		Name: "projects/" + projectID,
	}

	_, err = client.GetProject(ctx,req)
	if err != nil {
		return err
	}
	return nil
}