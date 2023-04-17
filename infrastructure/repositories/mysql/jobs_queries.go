package mysql

const (
	//JOB_GET get specific job base on input ID
	// 	params: input string
	// 	out: full job entity from database
	JOB_GET string = `SELECT * FROM app.jobs WHERE PK = UUID_TO_BIN(?, true);`

	//JOBS_GET get sub set of Jobs
	// 	params:
	//		term string,
	//		take int,
	//		skip int
	// 	out:
	//		jobs entities from database match fetch and offset provided
	JOBS_GET string = `call app.search_jobs(?, ?, ?)`

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
	JOBS_ADD string = `call app.create_jobs(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

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
	JOBS_UPDATE string = `call app.update_jobs(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	//JOBS_ADD add new Job
	// 	params:
	// 		PK guid,
	// 		deleted_by_user_id guid,
	// 		deleted_at date,
	// 	out:
	//		row affected
	JOBS_DELETE string = `call app.delete_jobs(?, ?, ?)`
)
