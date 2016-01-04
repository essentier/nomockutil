package nomockutil

import "testing"

func TestValidateServiceName(t *testing.T) {
	err := ValidateServiceName("invalidName")
	if err == nil {
		t.Errorf("invalidName should not be a valid service name.")
	}

	err = ValidateServiceName("invalid/name")
	if err == nil {
		t.Errorf("invalid/name should not be a valid service name.")
	}

	err = ValidateServiceName(".valid-service_name.012")
	if err != nil {
		t.Errorf(".valid-service_name.012 should be a valid service name.")
	}
}
