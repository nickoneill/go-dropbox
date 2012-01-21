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
const api_content_url = "https://api-content.dropbox.com/1/files/sandbox/"

type DropboxClient struct {
	Token string
	Client *http.Client
	Oauth *oauth.Client
}

type AccountInfo struct {
	Referral_link string
	Display_name string
	Country string
	Email string
	Uid uint32
	Quota_info *QuotaInfo
}

type QuotaInfo struct {
	Shared uint64
	Quota uint64
	Normal uint64
}

type DropFile struct {
	Size string
	Rev string
	Thumb_exists bool
	Bytes uint64
	Modified string
	Path string
	Is_dir bool
	Icon string
	Root string
	Mime_type string
	Revision uint32
	Contents []*DropFile
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

//
func (drop *DropboxClient) GetFile(creds *oauth.Credentials, path string) *DropFile {
	file := new(DropFile)
	
	err := getUrl(creds, api_content_url + path, nil, file)
	if err != nil {
		fmt.Printf("error getting file: %v",err)
	}
	
	return file
}

//
func (drop *DropboxClient) GetFileMeta(creds *oauth.Credentials, path string) *DropFile {
	file := new(DropFile)
	
	err := getUrl(creds, api_url + "metadata/sandbox/" + path, nil, file)
	if err != nil {
		fmt.Printf("error getting file: %v",err)
	}
	
	return file
}

// getUrl signs our API GET requests with our oauth credentials
func getUrl(creds *oauth.Credentials, getUrl string, params url.Values, data interface{}) error {
	if params == nil {
		params = make(url.Values)
	}
	
	oauthClient.SignParam(creds, "GET", getUrl, params)
	res, err := http.Get(getUrl + "?" + params.Encode())
	if err != nil {
		return err
	}
	defer res.Body.Close()
	
	// b, _ := ioutil.ReadAll(res.Body)
	// fmt.Printf("file: %v",string(b))
	
	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("Get request for %s returned %d, %s",getUrl,res.StatusCode,string(b))
	}
	
	return json.NewDecoder(res.Body).Decode(&data)
}