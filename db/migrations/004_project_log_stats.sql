CREATE TABLE project_log_stats (
    project_id TEXT PRIMARY KEY
        REFERENCES linked_projects(project_id)
        ON DELETE CASCADE,

    error_count BIGINT NOT NULL DEFAULT 0,
    last_error_at TIMESTAMP,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);
