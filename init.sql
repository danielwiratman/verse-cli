CREATE TABLE verses (
  id SERIAL PRIMARY KEY,
  address TEXT,
  content TEXT,
  created TIMESTAMP DEFAULT NOW()
);

INSERT INTO verses (address, content) VALUES
  ('Genesis 1:1', 'In the beginning, God created the heavens and the earth.');
