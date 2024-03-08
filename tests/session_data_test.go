package tests

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"xrplatform/arworld/backend/middleware/mysql"
	"xrplatform/arworld/backend/routes"
)

func TestGetSessionData(t *testing.T) {
	// set up a test router
	router := gin.Default()

	// connect db
	db := mysql.ConnectDB(router)
	defer db.Close()

	// test route
	router.POST("/ar-world/v1/session-data/get", routes.GetSessionData)

	// create a test request with sesion ID
	//jsonStr := []byte(`id=123456`)
	jsonStr := []byte(`{"SessionID": "123456"}`)
	req, err := http.NewRequest("POST", "/ar-world/v1/session-data/get", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Type", "application/json")

	// create a test HTTP recorder
	rr := httptest.NewRecorder()

	// Server the HTTP request to the test recorder
	router.ServeHTTP(rr, req)

	//check connect db

	//check send body json (raw)
	//

	// check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handeler return wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// check the response body
	expected := `{"data":"test chiều 6/3","status":200}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUploadSessionData(t *testing.T) {
	// set up a test router
	router := gin.Default()

	// connect db
	db := mysql.ConnectDB(router)
	defer db.Close()
	//check connect db

	// test route
	router.POST("/ar-world/v1/session-data/upload", routes.UploadSessionData)

	//tạo 2 input khác nhau
	// create a test request with sesion ID,
	SessionID1 := `"123456"`
	SessionData1 := `"test"`

	jsonStr1 := []byte(`{"SessionID": ` + SessionID1 + `, "SessionData": ` + SessionData1 + `}`)
	//SessionID2 := `"123456"`
	//SessionData2 := `"test"`
	//jsonStr2 := []byte(`{"SessionID": ` + SessionID2 + `, "SessionData": ` + SessionData2 + `}`)

	//creat request
	req, err := http.NewRequest("POST", "/ar-world/v1/session-data/upload", bytes.NewBuffer(jsonStr1))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// create a test HTTP recorder
	rr := httptest.NewRecorder()
	status := rr.Code
	// Server the HTTP request to the test recorder
	router.ServeHTTP(rr, req)

	//sửa ở đây, ko cần db, input khác nhau
	//ở đây check status, nếu nó trả về 500, sẽ phân thêm thành các trươnng hợp
	if status == http.StatusInternalServerError {
		t.Errorf("handeler return wrong status code: got %v want %v",
			status, http.StatusInternalServerError)

		// check the response body
		expectedExists := `{"error":"session ID already exists"}`
		if rr.Body.String() != expectedExists {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expectedExists)
		}
	} else if status == http.StatusOK {
		t.Errorf("handeler return wrong status code: got %v want %v",
			status, http.StatusOK)

		// check the response body
		expected := `{"data":"test chiều 6/3","status":200}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	} else {
		//expected := `{"data":"test chiều 6/3","status":200}`
		//if rr.Body.String() != expected {
		//	t.Errorf("handler returned unexpected body: got %v want %v",
		//		rr.Body.String(), expected)
	}
}
