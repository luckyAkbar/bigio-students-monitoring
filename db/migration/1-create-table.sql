-- +migrate Up notransaction

CREATE TYPE role AS ENUM ('ADMIN', 'STUDENT', 'TEACHER');
CREATE TYPE mark AS ENUM ('A', 'B', 'C', 'D', 'E');

CREATE TABLE IF NOT EXISTS admins (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS grades (
    id BIGINT PRIMARY KEY,
    student_id BIGINT,
    teacher_id BIGINT,
    subject_id BIGINT,
    mark mark NOT NULL,
    value INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS "sessions" (
    id BIGINT PRIMARY KEY,
    access_token TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    user_id BIGINT,
    role role NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expired_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS students (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS subjects (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    teacher_id BIGINT
);

CREATE TABLE IF NOT EXISTS teachers (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL
);

ALTER TABLE grades ADD FOREIGN KEY ("student_id") REFERENCES students("id");
ALTER TABLE grades ADD FOREIGN KEY ("teacher_id") REFERENCES teachers("id");
ALTER TABLE grades ADD FOREIGN KEY ("subject_id") REFERENCES subjects("id");

ALTER TABLE subjects ADD FOREIGN KEY ("teacher_id") REFERENCES teachers("id");

-- +migrate Down
DROP TABLE IF EXISTS admins;
DROP TABLE IF EXISTS grades;
DROP TABLE IF EXISTS "sessions";
DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS subjects;
DROP TABLE IF EXISTS teachers;