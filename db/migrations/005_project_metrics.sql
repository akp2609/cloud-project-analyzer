CREATE TABLE project_metrics (
    project_id TEXT NOT NULL,
    service_name TEXT NOT NULL,

    metric_type TEXT NOT NULL, 

    window_start TIMESTAMP NOT NULL,
    window_end TIMESTAMP NOT NULL,

    value DOUBLE PRECISION NOT NULL,

    created_at TIMESTAMP DEFAULT now(),

    PRIMARY KEY (
        project_id,
        service_name,
        metric_type,
        window_start
    )
);
