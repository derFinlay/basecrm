package service_test

import (
	"testing"

	"github.com/derfinlay/basecrm/service"
)

func TestGenerateRandomString(t *testing.T) {
	randomString := service.GenerateRandomString(12)
	if len(randomString) != 12 {
		t.Errorf("Random string was incorrect, got: %s, want: %s.", randomString, "12")
	}

}
