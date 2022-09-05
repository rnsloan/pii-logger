# PII Logger

To test Personally Identifiable Information redacting in logs. Go ports of the [faker library](https://github.com/faker-js/faker) have limited locale support.

## Usage

`go run ./cmd/main.go`

### Commands

`--delay` the time in seconds between outputs. Default: `5`
`--locale` the time between outputs. Default: `en-AU`
`--entitiesFilePath` the time between outputs. Default: `./pkg/pii/entities.toml`

## Development

`go test ./...`

## TODO?

- configure what entities to use
- format entities e.g '04## ### ###'
- natural language
## Entities

[https://cloud.google.com/dialogflow/es/docs/reference/system-entities](https://cloud.google.com/dialogflow/es/docs/reference/system-entities)

- Date & time
  - date
  - time
- Currency
  - amounts
- Geography
  - addresses
- Contacts
  - email address
  - phone number
  - IP address
- Personal
  - names
  - religion
  - race
  - height
  - weight
- Numbers
  - credit card numbers
  - social security number
  - passport
  - drivers licence
  - tax number
  - vehicle registration number
  - medical numbers
