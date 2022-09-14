# PII Logger

![preview](https://user-images.githubusercontent.com/2513462/189565365-ffeed618-3058-406e-8970-be4374cb296d.gif)

To test Personally Identifiable Information redacting in logs.

## Rationale:

1. Go ports of the [faker library](https://github.com/faker-js/faker) have limited locale support
2. To support Natural Language phrase output

## Usage

Download the appropriate executable from the releases page [https://github.com/rnsloan/pii-logger/releases](https://github.com/rnsloan/pii-logger/releases).

To run: `./pii-logger`.

PII 'Entities' data is stored in a toml file. It comes with a default file: `pkg/pii/entities.toml`. An alternative file can be specified via a [flag option](#options).

The entities file has support for reverse regular expression generation using the [regenerator package](https://pkg.go.dev/github.com/zach-klippenstein/goregen). To use, wrap the regular expression in `/` characters e.g. `"/04[0-9]{8}/"`. Support is basic and it is advised to test expressions with the regenerator package before usage as there are limitations with what the package can support versus the standard capabilities of a re2 regular expression.

### Options

`--delay` time in seconds between outputs. Default: `5`

`--locale` language locale to use. Supported locales are listed here: [https://stackoverflow.com/a/3191729](https://stackoverflow.com/a/3191729). Default: `en-AU`

`--entitiesFilePath` fully qualified path to the entities file. For a list of supported entities see the Entities section below. Default: `./pkg/pii/entities.toml`

`--specificEntities` specific entities to use e.g. `name,IPAddress`. Default: `all`

`--naturalLanguage` use the entities in natural language sentences if available. Options: `yes`, `no`,`always`. Default: `yes`

## Entities

Supported Entities:

- phone
- name
- IPAddress
- email
- religion
- creditCard
- medicalCard
- driversLicence
- vehicleRegistration
- pinNumber
- loyaltyCard
- gender
- password
- time
- date
- address
- currency

## Development

- `go run ./cmd/main.go`
- `go test ./...`

To build a new release:

1. add a new git tag
2. push the tag to GitHub
3. `make release`
4. create the new release in GitHub and add the files in `/build` generated from the previous step as assets
