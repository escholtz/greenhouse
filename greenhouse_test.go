package greenhouse

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestBoard(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.EscapedPath() {
		case "/v1/boards/github":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"name":"GitHub","content":""}`))
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"status":404,"error":"Job board not found"}`))
		}
	}))
	defer ts.Close()

	client := NewClient().WithHTTPClient(ts.Client())
	client.baseURL = ts.URL
	board, err := client.Board("github")
	if err != nil {
		t.Fatal(err)
	}
	if board.Name != "GitHub" {
		t.Errorf("got %s, want %s", board.Name, "GitHub")
	}

	board, err = client.Board("404")
	if err == nil {
		t.Errorf("expected 404 error for unknown board token")
	} else if err.Error() != "Job board not found" {
		t.Errorf("got %s, want %s", err.Error(), "Job board not found")
	}
}

func TestJobs(t *testing.T) {
	path := filepath.Join("testdata", "jobs-github.json")
	jobsStub, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.EscapedPath() {
		case "/v1/boards/github/jobs":
			w.WriteHeader(http.StatusOK)
			w.Write(jobsStub)
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"status":404,"error":"Jobs not found"}`))
		}
	}))
	defer ts.Close()

	client := NewClient().WithHTTPClient(ts.Client())
	client.baseURL = ts.URL
	jobs, err := client.Jobs("github")
	if err != nil {
		t.Fatal(err)
	}
	if jobs.Meta.Total != 71 {
		t.Errorf("got %v, want %v", jobs.Meta.Total, 71)
	}

	jobs, err = client.Jobs("404")
	if err == nil {
		t.Errorf("expected 404 error for unknown board token")
	} else if err.Error() != "Jobs not found" {
		t.Errorf("got %s, want %s", err.Error(), "Jobs not found")
	}
}
