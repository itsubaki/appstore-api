package appinfo

import (
	"sort"

	"github.com/itsubaki/apstapi/model"
	"golang.org/x/net/context"

	"google.golang.org/appengine/capability"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func Select(ctx context.Context, kind string, limit int) model.AppInfoList {
	if !capability.Enabled(ctx, "datastore_v3", "*") {
		log.Warningf(ctx, "datastore is currently unavailable.")
		return model.AppInfoList{}
	}

	list := model.AppInfoList{}
	q := datastore.NewQuery(kind).Order("ID").Limit(limit)
	for t := q.Run(ctx); ; {
		var r model.AppInfo
		_, err := t.Next(&r)
		if err == datastore.Done {
			break
		}
		if err != nil {
			log.Warningf(ctx, err.Error())
			return list
		}
		list = append(list, r)
	}
	sort.Sort(sort.Reverse(list))
	return list
}
