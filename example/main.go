package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/go-oauth"
	"github.com/nickoneill/go-dropbox"
	"io/ioutil"
	"net/http"
	"time"
)

const app_key = "ENTER_YOUR_APP_KEY_HERE"
const app_secret = "ENTER_YOUR_APP_SECRET_HERE"
const callback_url = "http://www.someurl.com/callback"

func main() {
	fmt.Printf("howdy\n")

	drop := dropbox.NewClient(app_key, app_secret)

	creds, err := load("config.json")
	if err != nil {
		tempcred, err := drop.Oauth.RequestTemporaryCredentials(http.DefaultClient, callback_url)
		if err != nil {
			fmt.Printf("err! %v", err)
			return
		}

		fmt.Printf("token stuff: %v %v\n", tempcred.Token, tempcred.Secret)

		url := drop.Oauth.AuthorizationURL(tempcred)
		fmt.Printf("auth url: %v\n", url)

		time.Sleep(15e9)

		newcreds, _, err := drop.Oauth.RequestToken(http.DefaultClient, tempcred, "")
		fmt.Printf("newcreds: %v", newcreds)
		err = save("config.json", newcreds.Token, newcreds.Secret)

	} else {
		fmt.Printf("loaded creds: %v\n", creds)

		// test for account info
		fmt.Printf("results: %v\n", drop.AccountInfo(creds))

		// test for get file
		// drop.GetFile(creds,"LLTP5.jpg")

		// test for file meta
		newfilemeta := drop.GetFileMeta(creds, "folder")
		fmt.Printf("files: %#v\n", newfilemeta)
		for i, thing := range newfilemeta.Contents {
			fmt.Printf("file %v: %#v\n", i, thing)
		}
	}
}

func save(fileName string, accessToken string, accessSecret string) error {
	config := oauth.Credentials{
		Token:  accessToken,
		Secret: accessSecret,
	}

	b, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, b, 0600)
	if err != nil {
		return err
	}

	return nil
}

func load(fileName string) (*oauth.Credentials, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	config := new(oauth.Credentials)

	err = json.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
