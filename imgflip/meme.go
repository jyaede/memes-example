package imgflip

//MemeResp ...
type MemeResp struct {
	Memes []Meme `json:"memes"`
}

//Meme ...
type Meme struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
