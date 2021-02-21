# clerk

An executable JSON validator with JSON schema to validate documents from local files or databases.

## Install

```bash
git clone git@github.com:lisez/clerk.git
make install
make build
```

## Usage

```bash
$ clerk -h
Usage:
  clerk [<schema> <file> [files...]] [flags]

Flags:
  -c, --config string   config file
  -h, --help            help for clerk
  -v, --version         version for clerk
```

### Example 1: Validate local files

```bash
$ clerk ./testlib/data/test_schema.json ./testlib/data/test_invalid_doc.json ./testlib/data/test_valid_doc.json
./testlib/data/test_invalid_doc.json: invalid JSON, reasons: [age: Invalid type. Expected: integer, given: string]
./testlib/data/test_valid_doc.json: pass
```

### Example 2: Validate documents from MongoDB

```bash
$ clerk -c ./testlib/data/test_config.yml
ObjectID("60328b85f18db7837a0d600e"): pass
ObjectID("60328b8ff18db7837a0d601c"): invalid JSON, reasons: [age: Invalid type. Expected: integer, given: string]
```

## Config

You can find the file at `testlib/data/test_config.yml`:

```yaml
clerk:
  schema:
    $id: "https://example.com/person.schema.json"
    $schema: "http://json-schema.org/draft-07/schema#"
    properties:
      age:
        minimum: 0
        type: integer
      firstName:
        type: string
      lastName:
        type: string
    title: Person
    type: object
  sourceRemotes:
    -
      args:
        database: test
        collection: people
      provider: mongodb
      uri: "mongodb://localhost:27017"
  targetFiles: # not implemented now
    - test_valid_doc.json
    - test_invalid_doc.json
```
