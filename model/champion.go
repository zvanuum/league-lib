package model

type Champions struct {
	Data map[string]Champion `json:"data"`
}

type Champion struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Name  string `json:"name"`
	Title string `json:"title"`
}
