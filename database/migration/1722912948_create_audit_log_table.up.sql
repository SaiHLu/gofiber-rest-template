CREATE TABLE IF NOT EXISTS audit_log (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid,
    entity_type VARCHAR (255),
    entity_id uuid,
    action VARCHAR (255),
    data jsonb,
    created_at TIMESTAMPTZ,

    FOREIGN KEY (user_id) REFERENCES "user"(id)
);
