# apstweb

app store web api

 - Collect Review from AppStore and Store to GCP Datastore/Search
 - Provide WebAPI for Search in GCP Datastore/Search

## Required

 - go 1.8
 - GCP Account

# How to Build

## Install

```console
$ go get github.com/itsubaki/apstweb
```

## Make GAE Application

```go
// main.go
package example

import (
    "github.com/itsubaki/apstweb"
)

func init() {
    apstweb.AppEngine()
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
$ gcloud app deploy app.yaml --project ${PROJECT_ID}
$ gcloud app deploy cron.yaml
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
