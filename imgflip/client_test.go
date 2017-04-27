package imgflip

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMemes(t *testing.T) {
	//Set up mock server and response
	body := `{"success":true,"data":{"memes":[{"id":"61579","name":"One Does Not Simply","url":"http://i.imgflip.com/1bij.jpg","width":568,"height":335},{"id":"101470","name":"Ancient Aliens","url":"http://i.imgflip.com/26am.jpg","width":500,"height":437}]}}`
	server, c := testClient(body)
	defer server.Close()

	//do request with mock server
	memes, err := c.GetMemes()

	//validate response
	assert.Nil(t, err)
	assert.Len(t, memes, 2)
	assert.NotEmpty(t, memes[0].ID)
	assert.NotEmpty(t, memes[0].Name)
	assert.NotEmpty(t, memes[0].URL)
	assert.NotEmpty(t, memes[0].Height)
	assert.NotEmpty(t, memes[0].Width)
}

func TestGetMemesError(t *testing.T) {
	//Set up mock server and response
	body := `{"success":false,"error_message":"Some hopefully-useful statement about why it failed"}`
	server, c := testClient(body)
	defer server.Close()

	//do request with mock server
	memes, err := c.GetMemes()

	//validate response
	assert.Nil(t, memes)
	assert.NotNil(t, err)
}

func TestCaptionImage(t *testing.T) {
	//Set up mock server and response
	body := `{"success":true,"data":{"url":"http://i.imgflip.com/123abc.jpg","page_url":"https://imgflip.com/i/123abc"}}`
	server, c := testClient(body)
	defer server.Close()

	//do request with mock server
	f := CaptionForm{
		ID:    "61579",
		Text0: "Top Text",
		Text1: "Bottom Text",
	}
	image, err := c.CaptionImage(f)

	//validate response
	assert.Nil(t, err)
	assert.NotEmpty(t, image.PageURL)
	assert.NotEmpty(t, image.URL)
}

func TestCaptionImageError(t *testing.T) {
	//Set up mock server and response
	body := `{"success":false,"error_message":"Some hopefully-useful statement about why it failed"}`
	server, c := testClient(body)
	defer server.Close()

	//do request with mock server
	f := CaptionForm{
		ID:    "61579",
		Text0: "Top Text",
		Text1: "Bottom Text",
	}
	image, err := c.CaptionImage(f)

	//validate response
	assert.Nil(t, image)
	assert.NotNil(t, err)
}

func testClient(body string) (*httptest.Server, *Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))
	c := New("username", "password")
	c.SetBaseURL(server.URL)
	return server, c
}
