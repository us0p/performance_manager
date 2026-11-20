ALTER TABLE app_user ALTER COLUMN id SET DEFAULT gen_random_uuid();
ALTER TABLE category ALTER COLUMN id SET DEFAULT gen_random_uuid();
ALTER TABLE task ALTER COLUMN id SET DEFAULT gen_random_uuid();
