CREATE EXTENSION IF NOT EXISTS pgcrypto;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS job_roles;

CREATE TABLE job_roles (
    role_id INT PRIMARY KEY,
    job_title VARCHAR(100) NOT NULL,
    department VARCHAR(50) NOT NULL
);

INSERT INTO job_roles (role_id, job_title, department) VALUES
(1, 'Assembly Line Worker', 'Production'),
(2, 'Robotics Technician', 'Engineering'),
(3, 'Quality Inspector', 'Quality Control'),
(4, 'Plant Manager', 'Management');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role_id INT NOT NULL REFERENCES job_roles(role_id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);