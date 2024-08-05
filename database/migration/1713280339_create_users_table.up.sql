CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR (50),
    email VARCHAR (300) UNIQUE NOT NULL
)