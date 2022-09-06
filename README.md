# PII Logger

To test Personally Identifiable Information redacting in logs. Go ports of the [faker library](https://github.com/faker-js/faker) have limited locale support.

The Entities toml file supports reverse regular expression generation using the [regenerator package](https://pkg.go.dev/github.com/zach-klippenstein/goregen). Wrap the regular expression in `/` characters e.g. `"/04[0-9]{8}/"`.

## Usage

`go run ./cmd/main.go`

### Commands

`--delay` the time in seconds between outputs. Default: `5`

`--locale` the language locale to use. Supported locales are listed here: [https://stackoverflow.com/a/3191729](https://stackoverflow.com/a/3191729). Default: `en-AU`

`--entitiesFilePath` path to the entities file. Any table header, e.g. `[phone]`, that does not exist in the default file will be ignored. Default: `./pkg/pii/entities.toml`

`--specificEntities` specific entities to use e.g. `name,IPAddress`. Default: `all`

## Development

`go test ./...`

## TODO?

- natural language
- add more entities to default file

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
