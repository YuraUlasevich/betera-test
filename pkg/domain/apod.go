package domain

type ApodRespArr struct {
	ApodRespArr []ApodResp
}

type ApodResp struct {
	Date string `json:"date"`

	Url string `json:"url"`
}
