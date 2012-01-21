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
	"github.com/garyburd/go-oauth"
)

const api_url = "https://api.dropbox.com/1/"

type DropboxClient struct {
	Token string
	Client *http.Client
	Oauth *oauth.Client
}

type QuotaInfo struct {
	Shared uint64
	Quota uint64
	Normal uint64
}

type AccountInfo struct {
	Referral_link string
	Display_name string
	Country string
	Email string
	Uid uint32
	Quota_info *QuotaInfo
}

var ( 
	oauthClient = oauth.Client{
		TemporaryCredentialRequestURI: "https://api.dropbox.com/1/oauth/request_token",
		ResourceOwnerAuthorizationURI: "https://www.dropbox.com/1/oauth/authorize",
		TokenRequestURI:               "https://api.dropbox.com/1/oauth/access_token",
	}
)

// returns a new dropbox object you can use to authenticate with and subsequently make API requests against
func NewClient(app_key string, app_secret string) *DropboxClient {
	oauthClient.Credentials = oauth.Credentials{
		Token:  app_key,
		Secret: app_secret,
	}
	
	return &DropboxClient{"", new(http.Client), &oauthClient}
}

// returns the account info for the credentialed user
func (drop *DropboxClient) AccountInfo(creds *oauth.Credentials) *AccountInfo {
	info := new(AccountInfo) 
	
	err := getUrl(creds, api_url + "account/info", nil, info)
	if err != nil {
		fmt.Printf("error getting account info: %v",err)
	}
	
	return info
}

// getUrl can sign any API GET requests with our oauth credentials
func getUrl(creds *oauth.Credentials, getUrl string, params url.Values, info *AccountInfo) error {
	if params == nil {
		params = make(url.Values)
	}
	
	oauthClient.SignParam(creds, "GET", getUrl, params)
	res, err := http.Get(getUrl + "?" + params.Encode())
	if err != nil {
		return err
	}
	
	defer res.Body.Close()
	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("Get request for %s returned %d, %s",getUrl,res.StatusCode,string(b))
	}
	
	return json.NewDecoder(res.Body).Decode(&info)
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