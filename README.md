# apstlib
app store data capture tool for google app engine


## Install & Deploy

```console
$ go get github.com/itsubaki/apstlib
$ ls
app.yaml cron.yaml main.go
$ gcloud app deploy app.yaml --project ${YOUR_GAE_PROJECT_ID}
```

## GAE Application

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
- description: "Capture Ranking"
  url: /ranking/capture
  schedule: every 24 hours
```

## Example

### Ranking

```console
$ curl "https://${YOUR_GAE_PROJECT_ID}.appspot.com/ranking?limit=100"
$ curl "https://${YOUR_GAE_PROJECT_ID}.appspot.com/ranking/search?query=ドラゴンズ&output=json"
```

### Review

```console
$ curl "https://${YOUR_GAE_PROJECT_ID}.appspot.com/review"
$ curl "https://${YOUR_GAE_PROJECT_ID}.appspot.com/review/search?id=${YOUR_IOS_APP_ID}&query=楽しい"
```
