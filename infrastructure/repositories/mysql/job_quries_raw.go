package mysql

const (
	//JOB_GET get specific job base on input ID
	// 	params: input string
	// 	out: full job entity from database
	JOB_GET_RAW string = `SELECT 
		app.jobs.*, 
		app.comments.*,
		app.users.*,
		app.companies.*
	FROM app.jobs 
	JOIN app.comments on comments.job_id = jobs.PK
	JOIN app.users on users.PK = jobs.author_id
	JOIN app.companies on companies.PK = jobs.company_id
	WHERE BIN_TO_UUID(jobs.PK, true)  = ?;`

	//JOBS_GET get sub set of Jobs
	// 	params:
	//		term string,
	//		take int,
	//		skip int
	// 	out:
	//		jobs entities from database match fetch and offset provided
	JOBS_GET_RAW string = `SELECT BIN_TO_UUID(jobs.PK, true), 
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
		WHERE jobs.title like ?
		OR jobs.content like ?
		OR jobs.short_description like ?
		LIMIT ?, ?;`

	//JOBS_ADD add new Job
	// 	params:
	// 		PK guid,
	// 		status string,
	// 		created_date UTC Date,
	// 		activated_date UTC Date,
	// 		order_id guid,
	// 		content string,
	// 		author_id guid,
	// 		title string,
	// 		tags string,
	// 		company_id guid,
	// 		rating float64,
	// 		short_description string,
	// 		full_description string,
	// 		topic_id guid,
	// 		code string,
	// 		image_thumbnail_path string,
	// 		video_thumbnail_path string,
	// 		is_approved bool
	// 	out:
	//		uuid of new jobs
	JOBS_ADD_RAW string = `INSERT INTO jobs(
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
        UUID_TO_BIN(?, true), 
        ?, 
        ?, 
        ?,
        UUID_TO_BIN(?, true),
        ?,
        UUID_TO_BIN(?, true),
        ?,
        ?,
        UUID_TO_BIN(?, true),
        ?,
        ?,
        ?,
        UUID_TO_BIN(?, true),
        ?,
        ?,
        ?,
        ?
    );`

	//JOBS_UPDATE update existing Job by pk
	// 	params:
	// 		PK guid,
	// 		status string,
	// 		activated_date UTC Date,
	// 		content string,
	// 		title string,
	// 		tags string,
	// 		rating float64,
	// 		short_description string,
	// 		full_description string,
	// 		topic_id guid,
	// 		image_thumbnail_path string,
	// 		video_thumbnail_path string,
	// 		is_approved bool
	// 	out:
	//		row affected
	JOBS_UPDATE_RAW string = `UPDATE jobs 
    SET status = ?, 
        activated_date = ?, 
        content = ?, 
        title = ?, 
        tags = ?, 
        rating = _rating, 
        short_description = ?, 
        full_description = ?,
        topic_id = UUID_TO_BIN(?, true),
        code = ?,
        image_thumbnail_path = ?,
        video_thumbnail_path = ?,
        is_approved = ?
    WHERE PK = UUID_TO_BIN(?, true)`

	//JOBS_ADD add new Job
	// 	params:
	// 		PK guid,
	// 		deleted_by_user_id guid,
	// 		deleted_at date,
	// 	out:
	//		row affected
	JOBS_DELETE_RAW string = `call app.delete_jobs(?, ?, ?)`
)
