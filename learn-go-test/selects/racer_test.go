package selects

import (
	"testing"
)

func TestRacer(t *testing.T) {

	//slowServer := httptest.NewServer(http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
	//	time.Sleep(20 * time.Millisecond)
	//	w.WriteHeader(http.StatusOK)
	//}))
	//
	//fastServer := httptest.NewServer(http.HandleFunc("", func(writer http.ResponseWriter, r *http.Request) {
	//	writer.WriteHeader(http.StatusOK)
	//}))

	slowUrl := "http://www.facebook.com"
	fastUrl := "http://www.quii.co.uk"

	want := fastUrl
	got, err := RacerSelect(slowUrl, fastUrl)
	if err != nil {
		t.Errorf("err msg%v", err)
	}
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
