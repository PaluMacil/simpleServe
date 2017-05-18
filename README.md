# simpleServe
Super simple Go static web server. No timeouts or SSL certs are configured, so please only use this for testing locally.

You can hit endpoint localhost:3033/echo to see an echo of your http request.
```
GET /echo HTTP/1.1
Host: localhost:3033
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
Accept-Encoding: gzip, deflate, sdch, br
Accept-Language: en-US,en;q=0.8
Connection: keep-alive
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36
```

Otherwise URI will be used to match against files in the `static` child folder of wherever you run this command. This folder comes with Bootstrap by default. This functionality might change.