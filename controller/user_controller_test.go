package controller_test

import (
	"testing"

	"github.com/derfinlay/basecrm/controller"
)

func TestGenerateUsername(t *testing.T) {
	username := controller.GenerateUsernameFromName("Finlay Sandy Gress")
	if username != "FiSaGr" {
		t.Errorf("Username was incorrect, got: %s, want: %s.", username, "FiSaGr")
	}
}
