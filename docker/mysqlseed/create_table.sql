use app;

CREATE TABLE IF NOT EXISTS comments (
	pk BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
	status  VARCHAR(50) DEFAULT '',
	created_date DATETIME,
	activated_date DATETIME,
	is_deleted TINYINT NULL DEFAULT 0,
	deleted_at  DATETIME NULL,
	deleted_by_user_id  BINARY(16) NULL,
	job_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)),
	user_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)),
	content TEXT,
	rating FLOAT(7,2) DEFAULT 0,
	is_approved TINYINT
);

CREATE TABLE IF NOT EXISTS companies (
	pk BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
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
	pk BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
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
	pk BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
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
	pk BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
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


DELIMITER //
CREATE PROCEDURE create_jobs(
    IN _pk VARCHAR(36), 
    IN _status VARCHAR(50), 
    IN _created_date DATETIME,
	IN _activated_date DATETIME,
	IN _is_deleted TINYINT ,
	IN _deleted_at  DATETIME ,
	IN _deleted_by_user_id  BINARY(16) ,
	IN _order_id BINARY(16),
	IN _content TEXT,
	IN _author_id  BINARY(16),
	IN _title  VARCHAR(255),
	IN _tags  VARCHAR(255),
	IN _company_id BINARY(16),
	IN _rating FLOAT(7,2),
	IN _short_description  VARCHAR(255),
	IN _full_description TEXT,
	IN _topic_id BINARY(16),
	IN _code  VARCHAR(6),
	IN _image_thumbnail_path VARCHAR(500)  ,
	IN _video_thumbnail_path VARCHAR(500),
    IN _is_approved TINYINT
    )
BEGIN
    INSERT INTO jobs(
        PK, 
        status, 
        created_date, 
        activated_date, 
        order_id, 
        content, 
        author_id, 
        title, 
        tags, 
        company_id, 
        rating, 
        short_description, 
        full_description,
        topic_id,
        code,
        image_thumbnail_path,
        video_thumbnail_path,
        is_approved
        )
    VALUES
    (
        UUID_TO_BIN(_PK, true), 
        _status, 
        _created_date, 
        _activated_date,
        UUID_TO_BIN(_order_id, true),
        _content,
        UUID_TO_BIN(_author_id, true),
        _title,
        _tags,
        UUID_TO_BIN(_company_id, true),
        _rating,
        _short_description,
        _full_description,
        UUID_TO_BIN(_topic_id, true),
        _code,
        _image_thumbnail_path,
        _video_thumbnail_path,
        _is_approved
    );
    SELECT LAST_INSERT_ID(); 
END//
DELIMITER ;

DELIMITER //
CREATE PROCEDURE update_jobs(
    IN _pk VARCHAR(36), 
    IN _status VARCHAR(50), 
	IN _activated_date DATETIME,
	IN _content TEXT,
	IN _title  VARCHAR(255),
	IN _tags  VARCHAR(255) ,
	IN _rating FLOAT(7,2) ,
	IN _short_description  VARCHAR(255),
	IN _full_description TEXT,
	IN _topic_id BINARY(16),
	IN _code  VARCHAR(6),
	IN _image_thumbnail_path VARCHAR(500) ,
	IN _video_thumbnail_path VARCHAR(500) ,
    IN _is_approved TINYINT
    )
BEGIN
    UPDATE jobs 
    SET status = _status, 
        activated_date = _activated_date, 
        content = _content, 
        title = _title, 
        tags = _tags, 
        rating = _rating, 
        short_description = _short_description, 
        full_description = _full_description,
        topic_id = _topic_id,
        code = _code,
        image_thumbnail_path = _image_thumbnail_path,
        video_thumbnail_path = _video_thumbnail_path,
        is_approved = _is_approved
    WHERE PK = UUID_TO_BIN(_pk, true);

END//
DELIMITER ;

DELIMITER //
CREATE PROCEDURE delete_jobs(IN _pk VARCHAR(36),  IN _deleted_by_user_id VARCHAR(36),  IN _deleted_at DATETIME)
BEGIN
    UPDATE jobs
    SET deleted_at = _deleted_at, 
        deleted_by_user_id = _deleted_by_user_id, 
        is_deleted = 1
    WHERE PK = UUID_TO_BIN(_pk, true);
END//
DELIMITER ;

DELIMITER //
CREATE PROCEDURE search_jobs(IN term VARCHAR(255), IN takes INT, IN skips INT)
BEGIN
    DECLARE search VARCHAR(257);
    SET search = (SELECT concat('%', term, '%'));
    SELECT BIN_TO_UUID(jobs.PK, true), 
        jobs.status, 
        jobs.created_date, 
        jobs.activated_date,
        BIN_TO_UUID(jobs.order_id, true),
        jobs.content,
        BIN_TO_UUID(jobs.author_id, true),
        jobs.title,
        jobs.tags,
        BIN_TO_UUID(jobs.company_id, true),
        jobs.rating,
        jobs.short_description,
        jobs.full_description,
        BIN_TO_UUID(jobs.topic_id, true),
        jobs.code,
        jobs.image_thumbnail_path,
        jobs.video_thumbnail_path,
        jobs.is_approved,
        app.comments.*,
        app.users.*,
        app.companies.*
        FROM app.jobs 
        JOIN app.comments on comments.job_id = jobs.PK
        JOIN app.users on users.PK = jobs.author_id
        JOIN app.companies on companies.PK = jobs.company_id
        WHERE BIN_TO_UUID(jobs.PK, true)  like search
        OR jobs.status like search
        OR BIN_TO_UUID(jobs.order_id, true) like search
        OR BIN_TO_UUID(jobs.author_id, true) like search
        OR BIN_TO_UUID(jobs.company_id, true) like search
        OR BIN_TO_UUID(jobs.topic_id, true) like search
        OR jobs.title like search
        OR jobs.content like search
        OR jobs.tags like search
        OR jobs.short_description like search
        OR jobs.full_description like search
        OR jobs.code like search
        LIMIT skips,takes;
END//
DELIMITER ;


DELIMITER //
CREATE PROCEDURE search_companies(IN term VARCHAR(255), IN takes INT, IN skips INT)
BEGIN
    DECLARE search VARCHAR(257);
    SET search = (SELECT concat('%', term, '%'));
    SELECT BIN_TO_UUID(companies.PK, true), 
        companies.status, 
        companies.created_date, 
        companies.activated_date,
        companies.company_name,
        companies.company_legal_name,
        companies.address,
        companies.phone_number,
        companies.tax_id,
        companies.registration_id,
        companies.parent_company_id,
        companies.code,
        companies.rating,
        companies.image_thumbnail_path,
        companies.video_thumbnail_path,
        companies.is_approved,
        app.company_reviews.*

        FROM app.companies 
        JOIN app.companies_reviews on companies.PK = companies_reviews.company_id
        WHERE BIN_TO_UUID(companies.PK, true)  like search
            OR companies.status like search
            OR BIN_TO_UUID(companies.order_id, true) like search
            OR BIN_TO_UUID(companies.author_id, true) like search
            OR companies.title like search
            OR companies.content like search
            OR companies.tags like search
            OR companies.short_description like search
            OR companies.full_description like search
            OR companies.code like search
        LIMIT skips,takes;
END//
DELIMITER ;

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

