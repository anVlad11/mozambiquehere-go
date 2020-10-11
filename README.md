# Mozambiquehe.re API Client

This library is a Golang API wrapper for https://apexlegendsapi.com/ .

## Install

```golang
go get github.com/anVlad11/mozambiquehere-go/v4
```

## Usage example

```golang
apexClient := client.DefaultClient(token)
data, err := apexClient.GetUserByName(models.PlatformPC, "anvlad11")
```


