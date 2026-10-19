CREATE TABLE IF NOT EXISTS task_category (
    task_id UUID NOT NULL REFERENCES task (id),
    category_id UUID NOT NULL REFERENCES category (id),
    PRIMARY KEY (task_id, category_id)
);
