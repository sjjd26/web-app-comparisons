-- Create a table for storing user credentials and metadata
CREATE TABLE users (
    id SERIAL PRIMARY KEY,          -- Unique ID for each user
    email VARCHAR(255) NOT NULL UNIQUE,     -- User's unique email
    password_hash TEXT NOT NULL,    -- Hashed password for security
    salt TEXT NOT NULL,             -- Salt for hashing
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- When the account was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- When the account was last updated
);

-- Create a table for tracking user sessions
CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,          -- Unique ID for the session
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE, -- Associated user
    token TEXT NOT NULL UNIQUE,     -- Session token
    expires_at TIMESTAMP NOT NULL,  -- Expiration date and time
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Session creation time
);