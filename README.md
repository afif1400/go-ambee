# **Ambee API client**

A wrapper for the [Ambee API](https://www.getambee.com)

Ambee provides accurate environmental intelligence for the world.

Its API are for - Weather - Air Quality - Pollen - Water Vapour - Soil - Fire

## **Docs**

## Installation

```go
go get github.com/afif1400/go-ambee
```

## Registration

For using the client you need to signup on [getambee.com](https://api-dashboard.getambee.com/#/signup)
and get the API Key to be used for initialization of client.

## Usage

After installation of the library using `go get`, and importing it into your project, ex:-

```go
import "github.com/afif1400/go-ambee"
```

### **Initializing client**

Create an instance of the AmbeeClient using the NewClient() function and pass the API Key accessed from the Ambee API dashboard.

```go
c := ambee.NewClient("your-api-key")
```

### **Functions**

### Get air quality by city

```go
res, err := ambee.GetAirQualityBycity("any-city-name")

// error handling
if err != nil{
    log.Fatalf("error: %v", err)
}
```

### **Response**

The function return a response and an error, the response struct [apiResponse](https://pkg.go.dev)

###
