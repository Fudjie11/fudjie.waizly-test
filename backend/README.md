
Prerequisite
- golang 1.21.6
- buf 1.xx.xx

How to run
1. `buf dep update`
2. `go mod tidy`
3. `go run serve-{app_type}`


Every update protobuf, should generate new protobuf definitions by running
`buf generate`