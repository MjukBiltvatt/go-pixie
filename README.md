go-pixie is a simple Go wrapper for the [Pixie SMS API](https://www.pixie.se/Home/Index).

# Getting started

## Installation
```
go mod init github.com/my/repo
go get github.com/jomla97/go-pixie
```

## Importing
``` go
import "github.com/jomla97/go-pixie"
```

## Complete example
* Numbers in `to` should be separated with a comma (`,`) and should each be lead with the country-code (e.g `+46` for Sweden)
* For a newline in `message`, use `\n`
``` go
//Create a new Client
px := pixie.NewClient("username", "password")

//Set a standard country code (not required, but if the first character of the phone number is `0` an error is returned otherwise)
px.StandardCountryCode = "+46"

//Send the SMS
err := px.Send("sender", "to", "message")
if err != nil {
    //Handle error
}
```
