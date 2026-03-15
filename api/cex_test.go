package api_test

import (
	"go/crypto/api"
	"testing"
)

func TestAPIcall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Error not found")
	}
}
