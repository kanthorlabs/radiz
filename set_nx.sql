INSERT INTO radiz_string (key, value) VALUES (?, ?) ON CONFLICT(key) DO NOTHING;