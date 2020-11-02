# App Metadata

Runs on **:8080**

Start the server:

```sh
go run main.go
```

Create a file `payload.yml` in this directory with the contents:

```yml
title: Valid App 1
version: 0.0.1
maintainers:
  - name: firstmaintainer app1
    email: firstmaintainer@hotmail.com
  - name: secondmaintainer app1
    email: secondmaintainer@gmail.com
company: Random Inc.
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: |-
  ### Interesting Title
  Some application content, and description
```

Run this request

```sh
curl -i -X POST http://localhost:8080 --data-binary '@payload.yml'
```

## Get the payload

```sh
curl -i http://localhost:8080/Valid%20App%201
```

## Get all payloads

```sh
curl -i http://localhost:8080/
```

## Filter by company

```sh
curl -i http://localhost:8080/?company=Random%20Inc.
```

## Update a company

```sh
curl -i -X PUT http://localhost:8080/Valid%20App%201 --data-binary '@payload.yml'
```
