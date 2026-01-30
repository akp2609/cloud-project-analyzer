CREATE TABLE project_features (
    project_id TEXT PRIMARY KEY REFERENCES linked_projects(project_id),

    cost_enabled BOOLEAN DEFAULT TRUE,
    metrics_enabled BOOLEAN DEFAULT TRUE,
    logs_enabled BOOLEAN DEFAULT FALSE,

    slo_enabled BOOLEAN DEFAULT FALSE,

    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
)