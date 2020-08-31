package pixie

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//Client contains the API credentials
type Client struct {
	Username string
	Password string
}

/*
Send an SMS via the Pixie API

- `sender` is the name that shows up on the receivers' phone.

- `to` is the phonenumber to which the SMS should be sent to,
multiple phonenumbers should be separated with commas (,).

- `message` is the message to send, use `\n` for newline.
*/
func (c Client) Send(sender string, to string, message string) error {
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
