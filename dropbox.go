package dropbox

import (
	"fmt"
	"net/url"
	"net/http"
	// "os"
	"io/ioutil"
	"encoding/json"
	// "bytes"
	// "time"
	// "alloy-d/goauth"
	"github.com/garyburd/go-oauth"
)

const api_url = "https://api.dropbox.com/1/"

type DropboxClient struct {
	Token string
	Client *http.Client
	Oauth *oauth.Client
}

type QuotaInfo struct {
	shared int32
	quota int32
	normal int32
}

type AccountInfo struct {
	referral_link string
	display_name string
	uid int16
	country string
	quota_info *QuotaInfo
}

var ( 
	oauthClient = oauth.Client{
		TemporaryCredentialRequestURI: "https://api.dropbox.com/1/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://www.dropbox.com/1/oauth/authorize",
		TokenRequestURI:               "https://api.dropbox.com/1/oauth/access_token",
	}
)

func NewClient(app_key string, app_secret string) *DropboxClient {
	oauthClient.Credentials = oauth.Credentials{
		Token:  app_key,
		Secret: app_secret,
	}
	
	return &DropboxClient{"", new(http.Client), &oauthClient}
}

func (drop *DropboxClient) AccountInfo(creds *oauth.Credentials) *AccountInfo {
	
	info := new(AccountInfo) 
	
	getUrl(creds, "https://api.dropbox.com/1/account/info", make(url.Values), info)
	
	fmt.Printf("info: %v",info)
	return info

	// if drop.oauth.Authorized() {
	// 	res, err := drop.oauth.Get("https://api.dropbox.com/1/account/info", map[string]string{})
	// 	if err != nil { return }
	// 
	// 	b, err := ioutil.ReadAll(res.Body)
	// 	fmt.Printf("res: %v",string(b))
	// } else {
	// 	fmt.Printf("Er, not authorized")
	// }
	
}

func getUrl(creds *oauth.Credentials, url string, params url.Values, info *AccountInfo) error {
	// if params == nil {
	// 	params = make(url.Values)
	// }
	
	oauthClient.SignParam(creds, "GET", url, params)
	res, err := http.Get(url + "?" + params.Encode())
	if err != nil {
		return err
	}
	
	defer res.Body.Close()
	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("Get request for %s returned %d, %s",url,res.StatusCode,string(b))
	}
	
	b, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("ok: %v\n",string(b))
	
	// var acctinfo []map[string]interface{}
	// json.NewDecoder(res.Body).Decode(acctinfo)
	// fmt.Printf("eh: %v\n",acctinfo)
	
	return json.NewDecoder(res.Body).Decode(info)
}

// func (drop *DropboxClient) RootFiles() {
// 	
// 	if drop.oauth.Authorized() {
// 		res, err := drop.oauth.Get("https://api.dropbox.com/1/metadata/sandbox/", map[string]string{})
// 		if err != nil { return }
// 
// 		b, err := ioutil.ReadAll(res.Body)
// 		fmt.Printf("res: %v",string(b))
// 	} else {
// 		fmt.Printf("Er, not authorized")
// 	}
// 	
// }
// 
// func (drop *DropboxClient) CreateFolder() {
// 	
// 	if drop.oauth.Authorized() {
// 		res, err := drop.oauth.Post("https://api.dropbox.com/1/fileops/create_folder", map[string]string{"root":"sandbox","path":"bloooog"})
// 		if err != nil { return }
// 
// 		b, err := ioutil.ReadAll(res.Body)
// 		fmt.Printf("res: %v",string(b))
// 	} else {
// 		fmt.Printf("Er, not authorized")
// 	}
// 	
// }


// func (drop *DropboxClient) getToken() (*http.Response, os.Error) {
// 	body := new(bytes.Buffer)
// 	write := multipart.NewWriter(body)
// 	oauth_consumer_key, _ := w.CreateFormField("oauth_consumer_key")
// 	oauth_consumer_key.Write([]byte(""))
// 	
// 	// oauth_signature_method:
// 	// oauth_signature:
// 	// oauth_timestamp:
// 	// oauth_nonce:
// 	
// 	req, _ := http.NewRequest("POST", fmt.Sprintf("%voauth/request_token",api_url), body)
// 	
// 	defer res.Body.Close()
// }