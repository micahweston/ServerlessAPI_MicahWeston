package main

import (
	"os"
	"context"
	"encoding/json"
	"log"
	"net"
	"github.com/joho/godotenv"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ipinfo/go/v2/ipinfo"
)

type Response struct {
	IP       string `json:"ipAddress"`
	HostName string `json:"hostName"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"location"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	token := "$MY_TOKEN"

	// Check if ipAddress is in queryStringParameters. If not return bad request response.
	if request.QueryStringParameters["ipAddress"] == "" {
		return events.APIGatewayProxyResponse{Body: "Invalid request paramaters", StatusCode: 400}, nil
	}
	ipAddress := request.QueryStringParameters["ipAddress"]

	if request.HTTPMethod == "GET" {
		ipAddress = request.QueryStringParameters["ipAddress"]
		client := ipinfo.NewClient(nil, nil, token)
		info, err := client.GetIPInfo(net.ParseIP(ipAddress))

		if err != nil {
			log.Fatal(err)
			return events.APIGatewayProxyResponse{Body: "Request failed please contact support", StatusCode: 400}, nil
		}

		// saving info from ipinfo to struct
		var ipResponse = Response{
			IP:       info.IP.String(),
			HostName: info.Hostname,
			City:     info.City,
			Region:   info.Region,
			Country:  info.Country,
			Loc:      info.Location,
			Org:      info.Org,
			Postal:   info.Postal,
			Timezone: info.Timezone,
		}

		// returning json object for get request
		resp, err := json.Marshal(ipResponse)
		if err != nil {
			log.Fatal(err)
			return events.APIGatewayProxyResponse{Body: "Request failed please contact support.", StatusCode: 400}, nil
		}
		return events.APIGatewayProxyResponse{Body: string(resp), StatusCode: 200}, nil
	} else {
		return events.APIGatewayProxyResponse{Body: "Invalid HTTP request", StatusCode: 428}, nil
	}
}
