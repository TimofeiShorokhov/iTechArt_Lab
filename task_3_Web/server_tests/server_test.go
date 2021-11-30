package server_tests

import (
	"github.com/stretchr/testify/assert"
	"iTechArt_Lab/task_3_Web/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRequest(t *testing.T) {

	r, _ := http.NewRequest("GET", "/students", nil)
	w := httptest.NewRecorder()

	handlers.GetStudents(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
