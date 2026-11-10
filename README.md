# go-lib

Small Go helpers for filters, collections, logging, MongoDB and MySQL.
Requires Go 1.26 or newer. Validate changes with `go test ./...` and `go vet ./...`;
CI also runs the race detector and checks reachable vulnerabilities.

The MongoDB integration retains the v1 driver types exposed by the existing API.
A v2 driver migration is a separate breaking change for consumers. No service is
deployed by this repository.

Filter struct field order defines the Participle grammar. Do not reorder those
fields for memory alignment: `identifier operator value` must stay in that order.
