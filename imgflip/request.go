package imgflip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
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

	//Convert struct to query parameters
	if i != nil {
		v, _ := query.Values(i)
		url += "?" + v.Encode()
	}

	//create http request and send
	req, err := http.NewRequest(method, url, nil)
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
