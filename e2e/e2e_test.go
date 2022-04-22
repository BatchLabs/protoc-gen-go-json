package main

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/protoc-gen-go-json/e2e/batch.com/myproto"
)

func TestGen(t *testing.T) {
	t.Run("uuid", func(t *testing.T) {
		uuid := &myproto.UUID{Data: make([]byte, 16)}
		copy(uuid.Data[:], "\xde\xad\xbe\xef\xde\xad\xbe\xef\xde\xad\xbe\xef\xde\xad\xbe\xef")

		var foo struct {
			UUID *myproto.UUID
		}
		foo.UUID = uuid

		data, err := json.Marshal(foo)
		if err != nil {
			t.Error(err)
		}

		const exp = `{"UUID":"deadbeef-dead-beef-dead-beefdeadbeef"}`
		if got := string(data); got != exp {
			t.Errorf("expected %q, got %q", exp, got)
		}
	})
}
