use app;

CREATE TABLE IF NOT EXISTS comments (
	pk BINARY(16) PRIMARY KEY,
	status  VARCHAR(50) DEFAULT '',
	created_date DATETIME,
	activated_date DATETIME,
	is_deleted TINYINT NULL DEFAULT 0,
	deleted_at  DATETIME NULL,
	deleted_by_user_id  BINARY(16) NULL,
	job_id BINARY(16),
	user_id BINARY(16),
	content TEXT,
	rating FLOAT(7,2) DEFAULT 0,
	is_approved TINYINT
);

CREATE TABLE IF NOT EXISTS companies (
	pk BINARY(16) PRIMARY KEY,
	status  VARCHAR(50) DEFAULT '',
	created_date DATETIME,
	activated_date DATETIME,
	is_deleted TINYINT NULL DEFAULT 0,
	deleted_at  DATETIME NULL,
	deleted_by_user_id BINARY(16) NULL,
	company_name VARCHAR(255),
	company_legal_name VARCHAR(255),
	address VARCHAR(255) NULL DEFAULT '',
	phone_number VARCHAR(20) NULL DEFAULT '',
	tax_id VARCHAR(50) NULL DEFAULT '',
	registration_id VARCHAR(255) NULL DEFAULT '',
	parent_company_id BINARY(16) NULL,
	code VARCHAR(6) NULL DEFAULT '',
    rating FLOAT(7,2) DEFAULT 0,
    type VARCHAR(50) NULL DEFAULT '',
    logo_path VARCHAR(500) NULL DEFAULT '',
    background_path VARCHAR(500) NULL DEFAULT '',
    description TEXT NULL
);

CREATE TABLE IF NOT EXISTS jobs (
	pk BINARY(16) PRIMARY KEY,
	status  VARCHAR(50) DEFAULT '',
	created_date DATETIME,
	activated_date DATETIME,
	is_deleted TINYINT  DEFAULT 0,
	deleted_at  DATETIME NULL,
	deleted_by_user_id  BINARY(16) NULL,
	order_id BINARY(16),
	content TEXT,
	author_id  BINARY(16),
	title  VARCHAR(255) DEFAULT '',
	tags  VARCHAR(255) DEFAULT '',
	company_id BINARY(16),
	rating FLOAT(7,2) DEFAULT 0,
	short_description  VARCHAR(255),
	full_description TEXT,
	topic_id BINARY(16),
	code  VARCHAR(6),
	image_thumbnail_path VARCHAR(500) NULL DEFAULT '',
	video_thumbnail_path VARCHAR(500) NULL DEFAULT '',
    is_approved TINYINT
);


CREATE TABLE IF NOT EXISTS users (
	pk BINARY(16) PRIMARY KEY,
	status  VARCHAR(50) DEFAULT '',
	created_date DATETIME,
	activated_date DATETIME,
	is_deleted TINYINT DEFAULT 0,
	deleted_at  DATETIME NULL,
	deleted_by_user_id  BINARY(16) NULL,
	company_id BINARY(16) NULL,
	first_name VARCHAR(50),
	mid_name VARCHAR(50) NULL,
	last_name VARCHAR(50),
	address VARCHAR(250) NULL,
	phone_number VARCHAR(20) NULL,
	email VARCHAR(50),
	user_name VARCHAR(50),
    avatar_path VARCHAR(500) NULL DEFAULT '',
    rating FLOAT(7,2) DEFAULT 0
);

CREATE TABLE IF NOT EXISTS company_reviews (
	pk BINARY(16) PRIMARY KEY,
	status  VARCHAR(50) DEFAULT '',
	created_date DATETIME,
	activated_date DATETIME,
	is_deleted TINYINT DEFAULT 0,
	deleted_at  DATETIME NULL,
	deleted_by_user_id  BINARY(16) NULL,
	parent_review_id BINARY(16) NULL,
	company_id BINARY(16),
	user_id BINARY(16),
	content TEXT,
	rating FLOAT(7,2),
	is_approved TINYINT,
	is_annonymous TINYINT,
	reactions VARCHAR(500),
    title  VARCHAR(50) DEFAULT '',
    stworking_history_monthatus  FLOAT(7,2) DEFAULT 0
);


INSERT INTO companies(
    pk ,
	status,
	created_date ,
	activated_date ,
	company_name ,
	company_legal_name ,
	address ,
	phone_number ,
	tax_id ,
	registration_id ,
	parent_company_id ,
	code ,
    rating,
    type
) 
VALUES ( 
    UUID_TO_BIN('b79b86ba-48eb-4062-b2d8-b70c1d0fcc4a', 1),
    'visible',
    UTC_TIMESTAMP(),
    UTC_TIMESTAMP(),
    'GalaxyAI',
    'GalaxyAI',
    '7 District, Ho Chi Minh city, Vietnam',
    '+(84)-88-888-6789',
    '123456789',
    '123456789',
    NULL,
    'G7',
    5.0,
    'product'
), ( 
    UUID_TO_BIN('c2458585-7762-4ed8-a09a-f8b374a71eb4', 1),
    'visible',
    UTC_TIMESTAMP(),
    UTC_TIMESTAMP(),
    'Company A',
    'Company A',
    '1 District, Ho Chi Minh city, Vietnam',
    '+(84)-88-888-8888',
    '123456780',
    '123456780',
    NULL,
    'G8',
    5.0,
    'product'
), ( 
    UUID_TO_BIN('e1a836e8-67d5-4b34-a3f8-2e4d512c5dd8', 1),
    'visible',
    UTC_TIMESTAMP(),
    UTC_TIMESTAMP(),
    'Company B',
    'Company B',
    '3 District, Ho Chi Minh city, Vietnam',
    '+(84)-88-888-9999',
    '123456781',
    '123456781',
    NULL,
    'G9',
    5.0,
    'product'
);

INSERT INTO users(
	pk,
	status,
	created_date,
	activated_date,
	company_id ,
	first_name ,
	mid_name ,
	last_name ,
	address ,
	phone_number ,
	email ,
	user_name ,
    avatar_path ,
    rating 
) 
VALUES (
    UUID_TO_BIN('0fcb476d-5dbb-490e-a6fb-e14a883f291b', 1),
    'online',
    UTC_TIMESTAMP(),
    UTC_TIMESTAMP(),
    UUID_TO_BIN('c2458585-7762-4ed8-a09a-f8b374a71eb4', 1),
    'Recruiter A',
    '',
    'Nguyen',
    '4 District, Ho Chi Minh city, Vietnam',
    '+(84)-88-888-8881',
    'recruitera@gmail.com',
    NULL,
    NULL,
    5.0
),(
    UUID_TO_BIN('f719fecb-3be9-4a5a-a86b-2e7100961fe9', 1),
    'online',
    UTC_TIMESTAMP(),
    UTC_TIMESTAMP(),
    UUID_TO_BIN('c2458585-7762-4ed8-a09a-f8b374a71eb4', 1),
    'Recruiter V',
    '',
    'Tran',
    '3 District, Ho Chi Minh city, Vietnam',
    '+(84)-88-888-8882',
    'recruiterv@gmail.com',
    NULL,
    NULL,
    5.0
),(
    UUID_TO_BIN('c3829ab1-01c5-4fdf-b7d7-cc88379cbb94', 1),
    'online',
    UTC_TIMESTAMP(),
    UTC_TIMESTAMP(),
    UUID_TO_BIN('c2458585-7762-4ed8-a09a-f8b374a71eb4', 1),
    'Recruiter E',
    '',
    'Phung',
    '2 District, Ho Chi Minh city, Vietnam',
    '+(84)-88-888-8883',
    'recruitere@gmail.com',
    NULL,
    NULL,
    5.0
);

