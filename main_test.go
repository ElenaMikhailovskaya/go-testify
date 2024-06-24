package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NotEqual(t, nil, responseRecorder.Code)
	require.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestMainHandlerWhenCityNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=tula", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NotEqual(t, nil, responseRecorder.Code)
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	expected := `wrong city value`
	assert.Equal(t, expected, responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.NotEqual(t, nil, responseRecorder.Code)
	require.Equal(t, http.StatusOK, responseRecorder.Code)

	expected := totalCount
	answer := responseRecorder.Body.String()
	mapAnswer := strings.Split(answer, ",")

	assert.Len(t, mapAnswer, expected)
}
