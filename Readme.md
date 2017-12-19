# sk8

### Instructions

##### Development

* Run the client by changing to the client directory and executing "npm run dev"
* Run the server by running "go run \*.go" from the project root

##### Production

Check out the deploy script! I have an ssh shortcut for gcloud and I push my app up there automatically.

Otherwise, make sure to do things in this order. The fully built client will be incorported into the binary file.

* Build the client by changing to the client directory and executing "npm run build"
* In the project root directory, execute:

```
rice embed-go
go build
```

PS - make sure to set your GOOS and GOARCH environment variables accordingly for the target server
