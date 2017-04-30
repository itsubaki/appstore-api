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
  url: /capture/review?id=${YOUR_APP_ID}
  schedule: every 1 hours
```

## Deploy

```console
$ ls
app.yaml cron.yaml	main.go
$ gcloud app deploy app.yaml --project ${YOUR_GAE_PROJECT}
```
