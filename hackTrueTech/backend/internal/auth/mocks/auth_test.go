package mocks

import (
	"bytes"
	"encoding/json"
	"github.com/wlcmtunknwndth/hackBPA/internal/auth"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth_LogIn(t *testing.T) {

	testCases := []struct {
		testName   string
		usr        auth.User
		isAdmin    bool
		statusCode int
	}{
		{
			testName: "Valid creds",
			usr: auth.User{
				Username: "idkidkidk",
				Password: "idkidkidk",
			},
			isAdmin:    false,
			statusCode: 200,
		},
		{
			testName: "Invalid login",
			usr: auth.User{
				Username: "",
				Password: "1234432",
			},
			isAdmin:    false,
			statusCode: 401,
		},
		{
			testName: "Invalid password",
			usr: auth.User{
				Username: "aasd",
				Password: "",
			},
			isAdmin:    false,
			statusCode: 401,
		},
	}

	db := NewStorage(t)
	//db.GetPassword()
	authSrv := auth.Auth{Db: db}

	//srv := httptest.NewServer(authSrv)

	for _, val := range testCases {

		data, err := json.Marshal(val.usr)
		if err != nil {
			t.Errorf("couldn't marshall test case: %s", err.Error())
			return
		}

		req := httptest.NewRequest(http.MethodPost, "", bytes.NewReader(data))

		w := httptest.NewRecorder()

		db.Mock.On("GetPassword", val.usr.Username).Return(val.usr.Password, nil)
		db.Mock.On("IsAdmin", val.usr.Username).Return(val.isAdmin)

		authSrv.LogIn(w, req)

		res := w.Result()

		defer res.Body.Close()

		if res.StatusCode != val.statusCode {
			t.Errorf("wrong status code: expected %d, but got %d", val.statusCode, res.StatusCode)
			//return
		}

	}
}
