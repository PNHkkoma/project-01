package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"xrplatform/arworld/backend/middleware/mysql"
	"xrplatform/arworld/backend/models"
	"xrplatform/arworld/backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestGetSessionData(t *testing.T) {
	// test cases
	tests := []models.SessionGetData{
		{
			SessionID: "123456",
		},
		{
			SessionID: "12345",
		},
	}

	// test cases result
	expected := []string{
		"{\"data\":\"test chiều 6/3\",\"status\":200}",
		"{\"error\":\"session ID already exists\",\"status\":500}",
	}

	// set up a test router
	router := gin.Default()

	// connect db
	db := mysql.ConnectDB(router)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	// test route
	router.POST("/ar-world/v1/session-data/get", routes.GetSessionData)

	for i, test := range tests {
		// parse test data to json
		data, _ := json.Marshal(test)

		// create a test request with session ID
		req, err := http.NewRequest("POST", "/ar-world/v1/session-data/get", bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		// create a test HTTP recorder
		recorder := httptest.NewRecorder()

		// Server the HTTP request to the test recorder
		router.ServeHTTP(recorder, req)

		// check the status code
		if status := recorder.Code; status != http.StatusOK {
			t.Errorf("handeler return wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// check the response body
		log.Println(recorder.Body.String())
		if recorder.Body.String() != expected[i] {
			t.Errorf("handler returned unexpected body: got %v want %v",
				recorder.Body.String(), expected[i])
		}
	}
}

func TestUploadSessionData(t *testing.T) {
	// test cases
	tests := []models.SessionUploadData{
		{
			SessionID:   "123456x",
			SessionData: "test",
		},
		{
			SessionID:   "123456x",
			SessionData: "test",
		},
	}

	// test cases result
	expected := []string{
		"{\"data\":\"success\",\"status\":200}",
		"{\"error\":\"session ID already exists\",\"status\":500}",
	}

	// create router
	router := gin.Default()

	// connect db
	db := mysql.ConnectDB(router)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	// test route
	router.POST("/ar-world/v1/session-data/upload", routes.UploadSessionData)

	for i, test := range tests {
		// parse test data to json
		data, _ := json.Marshal(test)

		// create a test request with session ID and session data
		req, err := http.NewRequest("POST", "/ar-world/v1/session-data/upload", bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		// create a test HTTP recorder
		recorder := httptest.NewRecorder()

		// Server the HTTP request to the test recorder
		router.ServeHTTP(recorder, req)
		assert.Equal(t, http.StatusOK, recorder.Code)

		if recorder.Body.String() != expected[i] {
			t.Errorf("handler returned unexpected body: got %v want %v",
				recorder.Body.String(), expected[i])
		}
	}
}
