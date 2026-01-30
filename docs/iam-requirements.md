# Required IAM Roles for Linked Project

The following roles must be granted to the Analyzer service account:

- roles/logging.viewer
- roles/monitoring.viewer
- roles/cloudasset.viewer
- roles/resourcemanager.projectViewer

Optional (cost visibility):
- roles/billing.Viewer

No editor or owner permisiions are required