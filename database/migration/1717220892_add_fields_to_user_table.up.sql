ALTER TABLE "user"
ADD COLUMN password VARCHAR(255),
ADD COLUMN created_at TIMESTAMPTZ,
ADD COLUMN updated_at TIMESTAMPTZ,
ADD COLUMN deleted_at TIMESTAMPTZ;

CREATE INDEX idx_deleted_at ON "user"(deleted_at);
