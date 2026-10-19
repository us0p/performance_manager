CREATE TABLE IF NOT EXISTS friendship (
    user_id UUID NOT NULL REFERENCES app_user (id),
    follows_id UUID NOT NULL REFERENCES app_user (id),
    PRIMARY KEY (user_id, follows_id)
);
