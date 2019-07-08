package translation

import "testing"

func TestErrTranslation_Error(t *testing.T) {
	if ErrConfFallbackEmpty.Error() == "" {
		t.Error("failed to get an error description")
	}
}
