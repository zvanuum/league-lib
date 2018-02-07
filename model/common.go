package model

type Image struct {
	Full   string `json:"full,omitempty"`
	Group  string `json:"group,omitempty"`
	Sprite string `json:"sprite,omitempty"`
	H      int    `json:"h,omitempty"`
	W      int    `json:"w,omitempty"`
	Y      int    `json:"y,omitempty"`
	X      int    `json:"x,omitempty"`
}
