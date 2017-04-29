package model

import (
	"testing"

	"github.com/itsubaki/apst/client"
)

func TestFeed(t *testing.T) {
	b := client.Ranking(10, "", "grossing", "jp")
	if b == nil {
		t.Error("http get failed.")
	}

	f := NewAppFeed(b)
	if f == nil {
		t.Error("feed unmarshal failed.")
	}

}
