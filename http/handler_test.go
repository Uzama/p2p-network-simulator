package http

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

var h = newHandler()

type FakeReader int

func (FakeReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("error occurred while reading")
}

func TestJoin(t *testing.T) {
	tableTest := []struct {
		name               string
		reader             io.Reader
		expectedStatusCode int
		expectedOutput     string
	}{
		{
			name:               "test readall error",
			reader:             FakeReader(0),
			expectedStatusCode: http.StatusBadRequest,
			expectedOutput:     `{"message":"error occurred while reading","error":true,"data":null}`,
		},
		{
			name:               "no body passed",
			reader:             bytes.NewReader(nil),
			expectedStatusCode: http.StatusBadRequest,
			expectedOutput:     `{"message":"unexpected end of JSON input","error":true,"data":null}`,
		},
		{
			name:               "id < 1",
			reader:             bytes.NewReader([]byte(`{"id": 0}`)),
			expectedStatusCode: http.StatusBadRequest,
			expectedOutput:     `{"message":"id must be a positive integer","error":true,"data":null}`,
		},
		{
			name:               "negative value for capacity",
			reader:             bytes.NewReader([]byte(`{"id":1, "capacity":-1}`)),
			expectedStatusCode: http.StatusBadRequest,
			expectedOutput:     `{"message":"capacity must be none negative","error":true,"data":null}`,
		},
		{
			name:               "happy path",
			reader:             bytes.NewReader([]byte(`{"id":1, "capacity":1}`)),
			expectedStatusCode: http.StatusCreated,
			expectedOutput:     `{"message":"successfully joined","error":false,"data":1}`,
		},
		{
			name:               "provide node 1 twice",
			reader:             bytes.NewReader([]byte(`{"id":1, "capacity":1}`)),
			expectedStatusCode: http.StatusUnprocessableEntity,
			expectedOutput:     `{"message":"id 1 already reserved","error":true,"data":null}`,
		},
	}

	for _, testCase := range tableTest {
		t.Run(testCase.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/join", testCase.reader)
			if err != nil {
				t.Fatal(err)
			}

			// we create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()

			// our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			h.Join(rr, req)

			// check the status code is what we expect.
			if rr.Code != testCase.expectedStatusCode {
				t.Errorf("expected %v, but got %v", testCase.expectedStatusCode, rr.Code)
			}

			// check the response body is what we expect.
			if rr.Body.String() != testCase.expectedOutput {
				t.Errorf("expected %v, but got %v", testCase.expectedOutput, rr.Body.String())
			}
		})
	}
}

func TestTrace(t *testing.T) {
	tableTest := []struct {
		name               string
		expectedStatusCode int
		expectedOutput     string
	}{
		{
			name:               "happy case",
			expectedStatusCode: http.StatusOK,
			expectedOutput:     `{"message":"trace received","error":false,"data":["1(0/1)"]}`,
		},
	}

	for _, testCase := range tableTest {
		t.Run(testCase.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/trace", nil)
			if err != nil {
				t.Fatal(err)
			}

			// we create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()

			// our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			h.Trace(rr, req)

			// check the status code is what we expect.
			if rr.Code != testCase.expectedStatusCode {
				t.Errorf("expected %v, but got %v", testCase.expectedStatusCode, rr.Code)
			}

			// check the response body is what we expect.
			if rr.Body.String() != testCase.expectedOutput {
				t.Errorf("expected %v, but got %v", testCase.expectedOutput, rr.Body.String())
			}
		})
	}
}

func TestLeave(t *testing.T) {
	tableTest := []struct {
		name               string
		id                 int
		expectedStatusCode int
		expectedOutput     string
	}{
		{
			name:               "happy case",
			id:                 1,
			expectedStatusCode: http.StatusAccepted,
			expectedOutput:     `{"message":"successfully left","error":false,"data":1}`,
		},
		{
			name:               "negative value",
			id:                 -1,
			expectedStatusCode: http.StatusBadRequest,
			expectedOutput:     `{"message":"id must be a positive integer","error":true,"data":null}`,
		},
		{
			name:               "not number",
			id:                 0, // set to an alphabat
			expectedStatusCode: http.StatusBadRequest,
			expectedOutput:     `{"message":"strconv.Atoi: parsing \"a\": invalid syntax","error":true,"data":null}`,
		},
		{
			name:               "not exists node",
			id:                 2,
			expectedStatusCode: http.StatusUnprocessableEntity,
			expectedOutput:     `{"message":"cannot locate id 2 node","error":true,"data":null}`,
		},
	}

	for _, testCase := range tableTest {
		t.Run(testCase.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodDelete, "/leave", nil)
			if err != nil {
				t.Fatal(err)
			}

			vars := map[string]string{
				"id": strconv.Itoa(testCase.id),
			}

			// only for one case
			if testCase.id == 0 {
				vars["id"] = "a"
			}

			req = mux.SetURLVars(req, vars)

			// we create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()

			// our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			h.Leave(rr, req)

			// check the status code is what we expect.
			if rr.Code != testCase.expectedStatusCode {
				t.Errorf("expected %v, but got %v", testCase.expectedStatusCode, rr.Code)
			}

			// check the response body is what we expect.
			if rr.Body.String() != testCase.expectedOutput {
				t.Errorf("expected %v, but got %v", testCase.expectedOutput, rr.Body.String())
			}
		})
	}
}
