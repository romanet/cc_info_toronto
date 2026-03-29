module cc-info-toronto.org/programs

go 1.25.0

replace cc-info-toronto.org/conf => ../conf

replace cc-info-toronto.org/http_utils => ../http_utils

replace cc-info-toronto.org/centre => ../centre

replace cc-info-toronto.org/db_utils => ../db_utils

require (
	cc-info-toronto.org/centre v0.0.0-00010101000000-000000000000
	cc-info-toronto.org/conf v0.0.0-00010101000000-000000000000
	cc-info-toronto.org/http_utils v0.0.0-00010101000000-000000000000
	golang.org/x/text v0.35.0
)

require (
	cc-info-toronto.org/db_utils v0.0.0-00010101000000-000000000000 // indirect
	github.com/apache/arrow-go/v18 v18.4.1 // indirect
	github.com/duckdb/duckdb-go-bindings v0.1.24 // indirect
	github.com/duckdb/duckdb-go-bindings/darwin-amd64 v0.1.24 // indirect
	github.com/duckdb/duckdb-go-bindings/darwin-arm64 v0.1.24 // indirect
	github.com/duckdb/duckdb-go-bindings/linux-amd64 v0.1.24 // indirect
	github.com/duckdb/duckdb-go-bindings/linux-arm64 v0.1.24 // indirect
	github.com/duckdb/duckdb-go-bindings/windows-amd64 v0.1.24 // indirect
	github.com/duckdb/duckdb-go/arrowmapping v0.0.27 // indirect
	github.com/duckdb/duckdb-go/mapping v0.0.27 // indirect
	github.com/duckdb/duckdb-go/v2 v2.5.4 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/google/flatbuffers v25.9.23+incompatible // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jmoiron/sqlx v1.4.0 // indirect
	github.com/klauspost/compress v1.18.2 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	golang.org/x/exp v0.0.0-20251209150349-8475f28825e9 // indirect
	golang.org/x/mod v0.33.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/sys v0.41.0 // indirect
	golang.org/x/telemetry v0.0.0-20260209163413-e7419c687ee4 // indirect
	golang.org/x/tools v0.42.0 // indirect
	golang.org/x/xerrors v0.0.0-20240903120638-7835f813f4da // indirect
)
