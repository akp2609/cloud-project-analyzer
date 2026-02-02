CREATE TABLE cost_anomalis (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    project_id TEXT NOT NULL,
    service TEXT NOT NULL,

    date DATE NOT NULL,
    cost NUMERIC NOT NULL,
    baseline_cost NUMERIC NOT NULL,
    spike_ratio NUMERIC NOT NULL,

    severity TEXT NOT NULL CHECK (
        severity IN ('LOW','MEDIUM','HIGH')
    ),

    created_at TIMESTAMP DEFAULT now(),

    uniq_cost_anomaly UNIQUE (project_id, service, date)
)