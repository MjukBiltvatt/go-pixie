package pixie

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

//Client contains the API credentials
type Client struct {
	Username            string
	Password            string
	StandardCountryCode string
}

/*
Send an SMS via the Pixie API

- `sender` is the name that shows up on the receivers' phone.

- `to` is the phonenumber to which the SMS should be sent to,
multiple phonenumbers should be separated with commas (,).

- `message` is the message to send, use `\n` for newline.
*/
func (c Client) Send(sender string, to string, message string) error {
	if sender == "" {
		return errors.New("the specified `sender` is not valid")
	} else if to == "" {
		return errors.New("the specified `to` is not valid")
	} else if message == "" {
		return errors.New("the specified `message` is not valid")
	}

	reg, err := regexp.Compile("[^,+0-9]+")
	if err != nil {
		return fmt.Errorf("failed to compile regexp: %v", err.Error())
	}
	to = reg.ReplaceAllString(to, "")

	if c.StandardCountryCode != "" {
		var standardCountryCode = c.StandardCountryCode
		if standardCountryCode[0:1] != "+" {
			standardCountryCode = "+" + standardCountryCode
		}

		if to[0:1] == "0" {
			//First character in phone number is a zero, replace it with the standard country code
			var n []rune
			p := []rune(to)
			for i := 1; i < len(p); i++ {
				n = append(n, p[i])
			}

			to = standardCountryCode + string(n)
		}

		to = strings.ReplaceAll(to, ",0", ","+standardCountryCode)
	}

	resp, err := http.DefaultClient.Get(fmt.Sprintf("http://smsserver.pixie.se/sendsms.asp?account=%v&pwd=%v&receivers=%v&sender=%v&message=%v", url.QueryEscape(c.Username), url.QueryEscape(c.Password), url.QueryEscape(to), url.QueryEscape(sender), url.QueryEscape(message)))
	if err != nil {
		return fmt.Errorf("failed to GET: %v", err.Error())
	}

	//Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err.Error())
	}

	//Unmarshal xml
	var response Response
	err = xml.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal xml: %v", err.Error())
	} else if response.Code != "0" {
		return errors.New(response.Description)
	}

	return nil
}
