# appstore-api

## Feature

 - AppStore Review Search API Server with GAE

## Required

 - go 1.8

# How to Build

## Install

```console
$ go get github.com/itsubaki/appstore-api
```

## Make GAE Application

```go
package main

import "github.com/itsubaki/appstore-api"

func init() {
    api.Init()
}
```

```yaml
# app.yaml
runtime: go
api_version: go1

handlers:
- url: /.*
  script: _go_app
```

```yaml
# cron.yaml
- description: "Capture Review"
  url: /review/capture?id=${IOS_APP_ID}&name=${APP_NAME}
  schedule: every 1 hours
- description: "Capture Ranking"
  url: /ranking/capture
  schedule: every 24 hours
```

## Deploy

```console
$ ls
app.yaml cron.yaml main.go
$ gcloud app deploy app.yaml  --project ${PROJECT_ID}
$ gcloud app deploy cron.yaml --project ${PROJECT_ID}
```

## API Example

### AppInfo

```console
$ curl "https://${PROJECT_ID}.appspot.com/app"
```


### Ranking

```console
$ curl "https://${PROJECT_ID}.appspot.com/ranking"
$ curl "https://${PROJECT_ID}.appspot.com/ranking/search?id=${IOS_APP_ID}"
$ curl "https://${PROJECT_ID}.appspot.com/ranking/search?query=${IOS_APP_ARTIST}"
```

### Review

```console
$ curl "https://${PROJECT_ID}.appspot.com/review"
$ curl "https://${PROJECT_ID}.appspot.com/review/search?id=${IOS_APP_ID}&query=Rating:5"
```
