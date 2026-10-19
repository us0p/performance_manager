CREATE TABLE IF NOT EXISTS task (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES app_user (id),
    name VARCHAR NOT NULL,
    priority VARCHAR NOT NULL DEFAULT 'normal' CONSTRAINT task_priority CHECK (priority IN ('normal', 'medium', 'hight')),
    due_date TIMESTAMPTZ,
    detail TEXT,
    board VARCHAR NOT NULL DEFAULT 'backlog' CONSTRAINT task_board CHECK (board IN ('backlog', 'todo', 'in_progress', 'completed'))
);
