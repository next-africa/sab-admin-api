package persistence

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/aetest"
	"testing"
)

func GetNewContext(t *testing.T) (ctx context.Context, done func()) {
	ctx, done, err := aetest.NewContext()

	if err != nil {
		t.Fatal(err)
	}
	return
}
