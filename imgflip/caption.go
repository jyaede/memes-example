package imgflip

//CaptionResp ...
type CaptionResp struct {
	Success bool            `json:"succes"`
	Data    CaptionRespData `json:"data"`
}

//CaptionRespData ...
type CaptionRespData struct {
	URL     string `json:"url"`
	PageURL string `json:"page_url"`
}

//CaptionForm ...
type CaptionForm struct {
	ID    string    `json:"id"`
	Text0 string `json:"text0"`
	Text1 string `json:"text1"`
}

//CaptionRequest ...
type CaptionRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	TemplateID string    `json:"template_id"`
	Text0      string `json:"text0"`
	Text1      string `json:"text1"`
}
