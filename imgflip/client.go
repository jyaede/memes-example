package imgflip

import "net/http"

const (
	baseURL = "https://api.imgflip.com"
)

//Client imgflip client interface
// https://api.imgflip.com
type Client struct {
	username   string
	password   string
	httpClient *http.Client
	baseURL    string
}

//SetBaseURL Allow overrid of request url. useful for testing
func (c *Client) SetBaseURL(url string) {
	c.baseURL = url
}

//SetHTTPClient can be useful for testing
func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

//New ...
func New(username, password string) *Client {
	return &Client{username, password, http.DefaultClient, baseURL}
}

//GetMemes ...
func (c Client) GetMemes() ([]Meme, error) {
	var resp MemeResp
	err := c.request("/get_memes", "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Memes, nil
}

//CaptionImage add a caption to image to generate meme
func (c Client) CaptionImage(f CaptionForm) (*CaptionRespData, error) {
	creq := CaptionRequest{
		Username:   c.username,
		Password:   c.password,
		TemplateID: f.ID,
		Text0:      f.Text0,
		Text1:      f.Text1,
	}
	var cap *CaptionRespData
	err := c.request("/caption_image", "POST", creq, &cap)
	if err != nil {
		return nil, err
	}
	return cap, nil
}
