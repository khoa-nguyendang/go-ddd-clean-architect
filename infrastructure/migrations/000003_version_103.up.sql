use app;

ALTER TABLE jobs
ADD test varchar(100);

ALTER TABLE jobs
DROP COLUMN code;