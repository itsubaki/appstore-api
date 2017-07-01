package ranking

import (
	"github.com/itsubaki/apstapi/model"
	"golang.org/x/net/context"

	"google.golang.org/appengine/capability"
	"google.golang.org/appengine/delay"
	"google.golang.org/appengine/log"
)

var indexPutDelay = delay.Func("indexput", IndexPut)

func Taskq(ctx context.Context, name string, feed *model.AppFeed) {
	if !capability.Enabled(ctx, "taskqueue", "*") {
		log.Warningf(ctx, "taskqueue is currently unavailable.")
		return
	}

	indexPutDelay.Call(ctx, name, feed)
}
