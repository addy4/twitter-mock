package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var url_hydra_admin_introspect_token = "https://trusting-tereshkova-12o8uqnuqz.projects.oryapis.com/admin/oauth2/introspect"
var oryToken = "ory_pat_92Eavnc5HtAlxnoCLqMljVAGjQuMuIgy"

type TokenJSON struct {
	Active   bool   `json:"active,omitempty"`
	ClientId string `json:"client_id,omitempty"`
}

func ValidateToken(access_token string) *TokenJSON {

	var bearer = "Bearer " + oryToken

	params := url.Values{}
	params.Add("token", access_token)

	req, _ := http.NewRequest("POST", url_hydra_admin_introspect_token, strings.NewReader(params.Encode()))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	tokenData := &TokenJSON{}

	json.Unmarshal(body, &tokenData)

	return tokenData
}
