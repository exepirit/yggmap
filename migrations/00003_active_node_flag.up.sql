ALTER TABLE nodes ADD is_active INTEGER;
UPDATE nodes SET is_active = 0;