# PII Logger

To test Personally Identifiable Information filtering in logs.

## Usage

`go run ./cmd/main.go`

`go test ./...`

## TODO?

- delay between messages
- just entities
- output transformer
- natural language
- entity configuration
- entity locale support
- load your own entities
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
  - 


### Format
- entity name
  - local
      - entries