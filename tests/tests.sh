# This script run all unit tests and store coverage output into coverage.out
# Require go >= 1.13

go test ./... -v -race -coverprofile=coverage.out -covermode=atomic