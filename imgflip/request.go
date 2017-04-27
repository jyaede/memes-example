package imgflip

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//RespData ...
type RespData struct {
	Success      bool            `json:"success"`
	ErrorMessage string          `json:"error_message"`
	Data         json.RawMessage `json:"data"`
}

func (c Client) request(path, method string, i, o interface{}) error {

	//merge the path with the base url
	url := c.baseURL + path

	//Convert struct to json bytes
	var buf *bytes.Buffer
	if method == "POST" && i != nil {
		jsonValue, err := json.Marshal(i)
		if err != nil {
			return nil
		}
		buf = bytes.NewBuffer(jsonValue)
	} else {
		buf = bytes.NewBuffer(nil)
	}

	//create http request and send
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//Convert response body to bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//Convert body to struct
	var r RespData
	if err := json.Unmarshal(body, &r); err != nil {
		return err
	}
	//Check if request failed
	if !r.Success {
		return fmt.Errorf("imgflip request failed - %s", r.ErrorMessage)
	}

	//Convert response data to needed struct
	if err := json.Unmarshal(r.Data, o); err != nil {
		return err
	}

	return nil
}
