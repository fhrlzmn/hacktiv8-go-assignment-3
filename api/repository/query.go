package repository

var createTableQuery = `
	CREATE TABLE IF NOT EXISTS data (
		id SERIAL PRIMARY KEY,
		wind INT,
		water INT,
		wind_status VARCHAR(255),
		water_status VARCHAR(255),
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
`

var insertDataQuery = `
	INSERT INTO data (wind, water, wind_status, water_status) VALUES ($1, $2, $3, $4);
`

var getDataQuery = `
	SELECT * FROM data ORDER BY id DESC LIMIT 1;
`
