CREATE TABLE projects (
    id UUID PRIMARY KEY,
    project_id TEXT NOT NULL UNIQUE,
    provider TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE cost_daily (
    project_id TEXT NOT NULL,
    service TEXT NOT NULL,
    cost NUMERIC NOT NULL,
    usage_date DATE NOT NULL,
    PRIMARY KEY (project_id,service,usage_date)
);