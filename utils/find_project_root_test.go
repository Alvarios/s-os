package sosutils

import (
	"strings"
	"testing"
)

func TestFindProjectRoot(t *testing.T) {
	_, err := FindProjectRoot("s-os")
	if err != nil {
		t.Errorf("Error getting project root : %s", err.Error())
	}

	_, err = FindProjectRoot("kushuh-go")
	if err == nil {
		t.Error("No error when providing non existing root name.")
	} else if !strings.HasPrefix(err.Error(), "root not found") {
		t.Errorf("Unexpected error : %s", err.Error())
	}
}
