package brightcoveapi_test

import (
	"testing"

	// Frameworks
	"github.com/djthorpe/ytapi/brightcoveapi"
)

func Test_000(t *testing.T) {
	if client, err := brightcoveapi.NewClientWithCredentials(&brightcoveapi.Credentials{}); err != nil {
		t.Error(err)
	} else if client == nil {
		t.Error("Invalid client == nil")
	} else {
		t.Log(client)
	}
}
