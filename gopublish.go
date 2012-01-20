package main

import (
	"fmt"
	// "mustache"
	// "path"
	// "os"
	// "bytes"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"time"
	"github.com/garyburd/go-oauth"
	// "./rsync"
	"./dropbox"
)

type Post struct {
	Title string
	Body string
}

const app_key = "ylg2zoaj78ol2dz"
const app_secret = "i2863bf9odkbdl7"
const callback_url = "http://nickoneill.name/callback"

func main() {
	fmt.Printf("howdy\n")
	
	drop := dropbox.NewClient(app_key,app_secret)
	
	creds, err := load("config.json")
	if err != nil {
		tempcred, err := drop.Oauth.RequestTemporaryCredentials(http.DefaultClient, callback_url)
	    if err != nil { fmt.Printf("err! %v", err); return }

		fmt.Printf("token stuff: %v %v\n",tempcred.Token,tempcred.Secret)

		url := drop.Oauth.AuthorizationURL(tempcred)
		fmt.Printf("auth url: %v\n",url)
		
		time.Sleep(15e9)
		
		newcreds, _, err := drop.Oauth.RequestToken(http.DefaultClient, tempcred, "")
		fmt.Printf("newcreds: %v",newcreds)
		err = save("config.json",newcreds.Token,newcreds.Secret)
		
	} else {
		fmt.Printf("loaded creds: %v\n",creds)
		
		fmt.Printf("results: %v\n",drop.AccountInfo(creds))
	}
	
	// drop.CreateFolder()
	
	// filename := path.Join(os.Getenv("PWD"),"source","index.mustache")
	// 
	// b := []byte(`{"Title":"hithere"}`)
	// //var p Post
	// var f interface{}
	// 
	// _ = json.Unmarshal(b, &f)
	// 
	// output := mustache.RenderFile(filename, f)
	// 
	// file, _ := os.OpenFile("static/index.html", os.O_RDWR | os.O_CREATE, 0666)
	// defer file.Close()
	// 
	// file.Write([]byte(output))
	// 
	// rsync.Rsync("static/index.html", "nickoneill", "nickoneill.name", "/var/www/nickoneill.name/public_html/test/")
}

func save(fileName string, accessToken string, accessSecret string) (error) {
	// file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	// defer file.Close()
	// if err != nil {
	// 	return err
	// }
	
	config := oauth.Credentials{
		Token: accessToken,
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