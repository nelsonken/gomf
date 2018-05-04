package gomf

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func Test_Demo(t *testing.T) {
	fb := New()
	fb.WriteField("name", "accountName")
	fb.WriteField("password", "pwd")
	fb.WriteFile("picture", "icon.png", "image/jpeg", []byte(strings.Repeat("0", 100)))

	t.Log(fb.GetBuffer().String())

	req, err := fb.GetHTTPRequest(context.Background(), "http://127.0.0.1:8080/up.php")
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)

	t.Log(res.StatusCode)
	t.Log(res.Status)

	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(b))
}
