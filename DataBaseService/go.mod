module github.com/ZeroZeroZerooZeroo/ChecklistApp/databaseservice

go 1.25.1

require (
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	golang.org/x/sys v0.34.0 // indirect
)

require (
	github.com/golang-migrate/migrate/v4 v4.19.0
	github.com/lib/pq v1.10.9
)

replace github.com/ZeroZeroZerooZeroo/ChecklistApp/proto => ../proto
