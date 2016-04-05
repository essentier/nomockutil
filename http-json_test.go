package nomockutil_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/essentier/nomockutil"
)

func TestWriteObjectOrErr(t *testing.T) {
	w := httptest.NewRecorder()
	var err error
	err = nil
	nomockutil.WriteObjectOrErr(w, "data", http.StatusUnauthorized, err)
	t.Logf("writer %v", w)
}
