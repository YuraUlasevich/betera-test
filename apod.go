package beteratest

type Apod struct {
	Date string `json:"date"`
	Url  string `json:"url"`
}

type ApodDB struct {
	Id    int    `json:"-"`
	Url   string `json:"url"`
	Image []byte `json:"image"`
}
