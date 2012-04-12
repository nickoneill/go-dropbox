# go-dropbox

A dropbox library written in Go. I whipped up a few quick calls that I needed for a separate project but by no means does it cover the entire [Dropbox API](https://www.dropbox.com/developers/reference/api). Enjoy, and please contribute if you require more API methods.

It's compatible with go1 or at least a recent weekly release.

### dependencies

The library uses the most recent release of [garyburd/go-oauth](http://github.com/garyburd/go-oauth) and will be automatically installed if you use `go get` to grab the package.

### quick start

Install the library: `go get github.com/nickoneill/go-dropbox`, then `import "github.com/nickoneill/go-dropbox"` to use in your project. Create a new client with your app key and secret, then use the exported oauth client to request credentials.

Check the example project for more detailed usage.