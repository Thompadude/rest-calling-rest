# GoLang Rest Application Fetching JSON
For learning purposes. Application redirects a `GET` request to the given URL. The URL is determined by the path param `getUrl`.
## Prerequisites
GoLang installed and GOPATH set
## How To
* Navigate to project directory
* Run `go run`
* The server is now running on port `8080`
* Send request to URL (expecting `JSON` response) using the `getUrl` path parameter, example URL: `http://localhost:8080/get-json?getUrl=<REST endpoint URL>`  
If an empty string is provided the application will send a request to `https://jsonplaceholder.typicode.com/posts/1`
