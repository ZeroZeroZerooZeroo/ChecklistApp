module github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice

go 1.25.1

require (
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	google.golang.org/grpc v1.75.1 // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)

require (
	github.com/golang-migrate/migrate/v4 v4.19.0
	github.com/lib/pq v1.10.9
)

require github.com/ZeroZeroZerooZeroo/ChecklistApp/proto v0.0.0

replace github.com/ZeroZeroZerooZeroo/ChecklistApp/proto => ../proto
