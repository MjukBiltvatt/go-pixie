go-pixie is a simple Go wrapper for the [Pixie SMS API](https://www.pixie.se/Home/Index).

## Installation
```
go get github.com/MjukBiltvatt/go-pixie
```

## Importing
``` go
import "github.com/MjukBiltvatt/go-pixie"
```

## Complete example
* Multiple phone numbers in `to` should be separated with a comma (`,`)
* For a newline in `message`, use `\n`
``` go
//Create a new Client
px := pixie.New("username", "password")

//Set a standard country code (not required, but if the first character of the phone number is `0` an error is returned otherwise)
px.StandardCountryCode = "+46"

//Send the SMS
err := px.Send("sender", "to", "message")
if err != nil {
    //Handle error
}
```
