package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var url_hydra_admin = "https://trusting-tereshkova-12o8uqnuqz.projects.oryapis.com/admin/oauth2/introspect"
var oryToken = "ory_pat_92Eavnc5HtAlxnoCLqMljVAGjQuMuIgy"
var introspectEndpoint = "/oauth2/instrospect"

type TokenJSON struct {
	Active   bool   `json:"active,omitempty"`
	ClientId string `json:"client_id,omitempty"`
}

func ValidateToken(access_token string) *TokenJSON {

	var bearer = "Bearer " + oryToken

	params := url.Values{}
	params.Add("token", access_token)

	req, _ := http.NewRequest("POST", url_hydra_admin, strings.NewReader(params.Encode()))
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(params.Encode())))

	fmt.Println(req)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	tokenData := &TokenJSON{}

	json.Unmarshal(body, &tokenData)

	fmt.Println("testtttttt")

	fmt.Println(string(body))

	fmt.Println(tokenData)
	fmt.Println(tokenData.Active)

	return tokenData
}
