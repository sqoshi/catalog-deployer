package api

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

type mockSvc struct {
}

// have 'mockSvc' implement the interface

func TestGetMaterialByFilter(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// now you can set mockSvc into the test context
	c.Set("svc_context_key", &mockSvc{})

}
