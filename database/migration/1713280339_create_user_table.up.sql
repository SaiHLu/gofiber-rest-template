CREATE TABLE IF NOT EXISTS "user" (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR (50) NOT NULL,
    email VARCHAR (300) UNIQUE NOT NULL
)
