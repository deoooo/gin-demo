package e2e

import (
	"github.com/deoooo/gin_demo/routers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	gin.SetMode("test")
	req, err := http.NewRequest("GET", "/health", nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()

	router := routers.InitRouter()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	assert.Equal(t, rr.Body.String(), "{\"code\":200,\"msg\":\"ok\"}")
}
