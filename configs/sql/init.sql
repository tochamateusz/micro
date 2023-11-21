CREATE DATABASE test;
\connect test;

CREATE TABLE test_table(
  id bigserial PRIMARY KEY,
  name varchar
);

INSERT INTO test_table (
  name
) VALUES ( 'test_entry' )

