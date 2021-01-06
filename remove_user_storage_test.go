package wxgameod

import (
	"errors"
	"github.com/birjemin/wxgameod/utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestRemoveUserStorage test removeUserStorage
func TestRemoveUserStorage(t *testing.T) {

	var sessionKey = "tiihtNczf5v6AKRyjwEUhQ=="
	var httpClient = &utils.HTTPClient{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	var ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.EscapedPath()
		if path != removeUserStorageURI {
			t.Fatalf("path is invalid: %s, %s'", removeUserStorageURI, path)
		}

		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}

		for _, v := range []string{"access_token", "signature"} {
			content := r.Form.Get(v)
			if content == "" {
				t.Fatalf("param %v can not be empty", v)
			}
		}

		body, _ := ioutil.ReadAll(r.Body)
		kvList := string(body)

		if kvList == "" {
			t.Fatal("body is empty")
		}

		if Encrypt(sessionKey, kvList) != r.Form.Get("signature") {
			t.Fatal("signature is invalid")
		}

		w.WriteHeader(http.StatusOK)

		raw := `{"errcode":0,"errmsg":""}`
		if _, err := w.Write([]byte(raw)); err != nil {
			t.Fatal(err)
		}
	}))

	defer ts.Close()

	var m = RemoveUserStorage{
		AccessToken: "ACCESSTOKEN",
		SessionKey:  sessionKey,
		OpenID:      "odkx20ENSNa2w5y3g_qOkOvBNM1g",
		Key:         "{\"key\":[\"gold\",\"score\"]}",
		HTTPRequest: httpClient,
	}

	if ret, err := m.doRemoveUserStorage(ts.URL); err != nil {
		t.Error(err)
	} else {
		if ret.ErrCode != 0 {
			t.Error(errors.New("msg: " + ret.ErrMsg))
		}
	}
}
