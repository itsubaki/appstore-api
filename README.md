# apstlib
app store data capture tool for google app engine


## Install

```console
$ go get github.com/itsubaki/apstlib
```

## Make Google App Engine Application

### main.go

```go
package example

import (
	"github.com/itsubaki/apstlib"
)

func init() {
	apstlib.AppEngine()
}
```

### app.yaml

```yaml
runtime: go
api_version: go1

handlers:
- url: /.*
  script: _go_app
```

### cron.yaml

```yaml
- description: "Capture Review"
  url: /review/capture?id=${YOUR_IOS_APP_ID}
  schedule: every 1 hours
- description: "ranking"
  url: /ranking/capture
  schedule: every 24 hours
```

## Deploy

```console
$ ls
app.yaml cron.yaml main.go
$ gcloud app deploy app.yaml --project ${YOUR_GAE_PROJECT_ID}
```

## Example

### Ranking

```console
$ curl "https://${YOUR_GAE_PROJECT_ID}.appspot.com/ranking?output=json&limit=100"
$ curl "https://${YOUR_GAE_PROJECT_ID}.appspot.com/ranking/search?output=json&query=ドラゴンズ"
```

### Review

```console
$ curl "https://${YOUR_GAE_PROJECT_ID}.appspot.com/review?output=json"
$ curl "https://${YOUR_GAE_PROJECT_ID}.appspot.com/review/search?id=493470467&output=json&query=運営"
```
