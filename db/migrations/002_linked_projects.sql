CREATE TABLE linked_projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    project_id TEXT NOT NULL UNIQUE,
    project_number TEXT NOT NULL,
    billing_account TEXT,

    status TEXT NOT NULL CHECK (
        status IN ('PENDING','ACTIVE','ERROR')
    ),

    iam_verified BOOLEAN NOT NULL DEFAULT FALSE,
    billing_verified BOOLEAN NOT NULL DEFAULT FALSE,
    monitoring_verified BOOLEAN NOT NULL DEFAULT FALSE,

    last_verfied_at TIMESTAMP,
    error_reason TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
)