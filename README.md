# DomeneShop-DynDNS
This is a DynDNS application for Domeneshop

its intended to run as a cronjob,  domeneshop-ddns-amd64 is compiled for Linux amd64
one or many domains are supported, it the record does not exist, it will be created to your public IP.

Get api key:
https://domene.shop/admin?view=api


Domeneshop api is currently in v0 and beta, so this may break without notice.

## How to run
all parameters can be passed into the app:
domeneshop-ddns-amd64 -domains=test.example.com,subdomain2.example.no -token=<token> -secret=<secret>

optionally, a app.yaml file can be placed in the working directory.

````
token: "your-token-here"
secret: "your-secret-here"
domains:
  - "test.example.no"
  - "sub.example.com"
````




## Build
GOOS=linux GOARCH=amd64 CGOENABLED=0 go build  -o domeneshop-ddns-amd64  main