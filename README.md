# clerk

A command-line JSON validator to validate documents with JSON schema from local files or databases.

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
2021/02/22 02:25:51 ./testlib/data/test_invalid_doc.json: invalid JSON, reasons: [age: Invalid type. Expected: integer, given: string]
2021/02/22 02:25:51 ./testlib/data/test_valid_doc.json: pass
```

### Example 2: Validate documents from MongoDB

```bash
$ go run main.go -c ./testlib/data/test_config.yml
ObjectID("60328b85f18db7837a0d600e"): pass
ObjectID("60328b8ff18db7837a0d601c"): invalid JSON, reasons: [age: Invalid type. Expected: integer, given: string]
```
