# ḎOME Network - Library

## Requirements

* `docker`
* `docker-compose`
* `go v1.18+`
* `godoc`
* `swag`
* `swagger`

## Build

The build system takes advantages of build tags to allow the inclusion of specific files for targeted architectures.

There is no direct building of this library.

## Test

To test the Go code run `bash scripts/test.sh` in a console.

## Structure

A generalized overview of the folder structure.

* `/config` - environment specific configurations
* `/pkg` - various packages
  * `/api` - common service API functions and types
  * `/cfg` - configuration helpers
  * `/codec` - standardized encodings (gob, hex, json)
  * `/coin` - supported wallet coins
  * `/common` - global types
  * `/crypto` - convenience functions
  * `/db` - key-value stores and SQL databases
  * `/eth` - a Go+JS compatible Ethereum client
  * `/fetch` - HTTP client that automatically handles authorization using the identity key
  * `/io` - IO interface for local persistence (file system, JS localStorage)
  * `/log` - internal log functions
  * `/metric` - Prometheus compatible VictoriaMetrics client adapter with Gin handler
  * `/test` - simplify service API handler testing
  * `/valid` - validation functions for common types
  * `/wallet` - HD wallet interface definition and basic implementation
* `/scripts` - helpful bash scripts
  * `/build.sh` - simplified building options
  * `/test.sh` - run tests, runs `go test ./...` but used to do more
