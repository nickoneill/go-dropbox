# go-dropbox

A dropbox library written in Go. I whipped up a few quick calls that I needed for a separate project but by no means does it cover the entire [Dropbox API](https://www.dropbox.com/developers/reference/api). Enjoy, and please contribute if you require more API methods.

### dependencies

The library uses garyburd/go-oauth, usually you can just do `goinstall github.com/garyburd/go-oauth` to install.

### quick start

Use `goinstall github.com/nickoneill/go-dropbox` to install, then `import "github.com/nickoneill/go-dropbox"` to use in your project. Create a new client with your app key and secret, then use the exported oauth client to request credentials. After that you're free to make your dropbox calls. Use the example project if you need more instruction.
