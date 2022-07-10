# catalog-deployer

## Introduction

Tool observes files and subdirectories in given directory and shares its contents via API.

If searched path is a directory then the list of files and subdirectories (not recursive) within will be shown, else the
file content.
So if we have for example a `ROOT_DIR=/storage`

```text
/
└── storage
    └── substorage1
        └── file.txt
```

and we want to check the content of file.txt, we need to send request GET
to `$(DEPLOY_HOST):$(DEPLOY_PORT)/substorage1<slash>file.txt`.

It is important to use `<slash>` and to not add `/storage` in request link. `$ROOT_DIR` is hidden behind first `/` not
written as `<slash>`.
`<slash>` is a trick which allow api to notice new directories on the lowest height in a tree.

## Launch

It is required to set `ROOT_DIR` variable with path to directory that should be observed and shared.

API is deployed on `0.0.0.0:8080` by default, it can be overridden by setting environment variables `DEPLOY_HOST`
and `DEPLOY_PORT`.

Example launch
```shell
ROOT_DIR=./storage go run .
```

Example `docker-compose.yml` setup.

```yaml
version: '3.8'

services:
  catalog_deployer:
    build:
      context: .
    ports:
      - "8080:8080"
    environment:
      - ROOT_DIR=/storage
    volumes:
      - ./tests/test_dir:/storage

volumes:
  storage:
```

## Code Example

```go

func GetEntityInfo(path string) Entity {
var (
eType   string
content string
)
name := filepath.Base(path)
if isDir(path) {
eType = "directory"
content = strings.Join(listEntities(path), ",")
} else {
fileContent, err := ioutil.ReadFile(path)
if err != nil {
log.Fatal(err)
}
eType = getFileType(name)
content = string(fileContent)
}

return Entity{
Name:    name,
Type:    eType,
Content: content,
}
}
```

