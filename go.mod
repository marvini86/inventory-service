module github.com/marvini86/inventory-service

go 1.23.0

toolchain go1.23.11

require (
	github.com/DATA-DOG/go-sqlmock v1.5.2
	github.com/jinzhu/copier v0.4.0
	github.com/joho/godotenv v1.5.1
	github.com/marvini86/car-service-protos v0.0.0
	github.com/stretchr/testify v1.10.0
	google.golang.org/grpc v1.73.0
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.30.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/marvini86/car-service-protos => ../car-service-protos
