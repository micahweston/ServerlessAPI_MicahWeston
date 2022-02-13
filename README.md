# Serverless IP Info API

API with GET only function to send in IPaddress and receive data about the IPAddress supplied.

## Description

This project is a code challenge to create a serverless API with AWS Lambda, AWS APIGateway, and Terraform. The API is created in Golang, with a simple function to call IPinfo API
to gather information about the IP address that was sent in with the GET request.

## Getting Started
### To send a API call you can submit from the browser, or use POSTman to make a GET request to the URL below. 
```
https://1joz2mq5pi.execute-api.us-west-2.amazonaws.com/api/?ipAddress=[YOUR_IP_GOES_HERE]
```
### Sample Response
```json
{
    "ipAddress": "8.8.8.8",
    "hostName": "dns.google",
    "city": "Mountain View",
    "region": "California",
    "country": "US",
    "location": "37.4056,-122.0775",
    "org": "AS15169 Google LLC",
    "postal": "94043",
    "timezone": "America/Los_Angeles"
}
```

## Help

Only thing to remember is to add the ip address you are wanting to retrieve information from in the query params as shown above.


## Authors

Contributors names and contact info

Micah Weston 
[@mmwest55](https://twitter.com/mmwest55)

### Inspiration, code snippets, etc.
* [IPInfo.io API](https://github.com/ipinfo/go)
