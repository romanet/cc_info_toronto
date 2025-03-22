module cc-info-toronto.org

go 1.24.1

replace cc-info-toronto.org/centre => ./centre

replace cc-info-toronto.org/http_utils => ./http_utils

replace cc-info-toronto.org/conf => ./conf

replace cc-info-toronto.org/programs => ./programs

replace cc-info-toronto.org/db_utils => ./db_utils

require (
	cc-info-toronto.org/centre v0.0.0-00010101000000-000000000000
	cc-info-toronto.org/conf v0.0.0-00010101000000-000000000000
	cc-info-toronto.org/db_utils v0.0.0-00010101000000-000000000000
	cc-info-toronto.org/programs v0.0.0-00010101000000-000000000000
)

require (
	cc-info-toronto.org/http_utils v0.0.0-00010101000000-000000000000 // indirect
	github.com/jmoiron/sqlx v1.4.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
	golang.org/x/text v0.23.0 // indirect
)
