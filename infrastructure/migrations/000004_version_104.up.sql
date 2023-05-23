use app;
ALTER TABLE jobs
DROP COLUMN test;

ALTER TABLE jobs
Add  code_new varchar(10);

Update jobs set code = '';