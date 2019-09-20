Build and expose endpoints on port 8080

```
go run *.go $(cat oauth_credentials.txt)
```

Install ngrok to expose local server to the internet.

Then, run the following command:

```
ngrok http 8080
```
