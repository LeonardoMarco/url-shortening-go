CREATE TABLE urls (
   id BIGSERIAL primary key,
   hash TEXT not null,
   url TEXT,
   created_at TIMESTAMP default now()
);