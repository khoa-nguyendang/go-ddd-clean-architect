package mysql

const (
	//COMPANY_GET get specific company base on input ID
	// 	params: input string
	// 	out: full company entity from database
	COMPANY_GET string = `SELECT * FROM app.companies WHERE PK = UUID_TO_BIN(?, true);`

	//COMPANIES_GET get sub set of companies
	// 	params:
	//		term string,
	//		take int,
	//		skip int
	// 	out:
	//		companies entities from database match fetch and offset provided
	COMPANIES_GET string = `call app.search_companies(?, ?, ?)`
)
