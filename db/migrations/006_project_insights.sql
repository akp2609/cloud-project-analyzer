CREATE TABLE project_insights (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id TEXT NOT NULL,

    insight_type TEXT NOT NULL,
    severity TEXT NOT NULL,

    title TEXT NOT NULL,
    description TEXT NOT NULL,

    detected_at TIMESTAMP NOT NULL DEFAULT now(),
    metadeta JSONB,

    UNIQUE (project_id,insight_type,title)
)