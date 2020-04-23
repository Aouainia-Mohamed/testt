package Controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatedataclip(t *testing.T) {
	var jsonStr = []byte(`{"Dataclipkey": "dc101","Name": "dcdriver","Description": "this dataclip will list all available drivers","Sqltext": "SELECT * FROM public.driver","Nargs": 3 ,"Argsdesc":"qsdqsfqsfqsfqsf"}`)
	request, err := http.NewRequest("POST", "/Dataclips", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateDataclip)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
