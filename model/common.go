package model

type VersionedData struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

type Image struct {
	Full   string `json:"full"`
	Group  string `json:"group"`
	Sprite string `json:"sprite"`
	H      int    `json:"h"`
	W      int    `json:"w"`
	Y      int    `json:"y"`
	X      int    `json:"x"`
}

type MapData struct {
	Data map[string]MapDetails `json:"data"`
	*VersionedData
}

type MapDetails struct {
	MapName               string  `json:"mapName"`
	Image                 Image   `json:"image"`
	MapID                 int64   `json:"mapId"`
	UnpurchasableItemList []int64 `json:"unpurchasableItemList"`
}

type MasteryList struct {
	Data map[string]Mastery `json:"data"`
	Tree MasteryTree        `json:"tree"`
	*VersionedData
}

type Mastery struct {
	Prereq               string   `json:"prereq"`
	MasteryTree          string   `json:"masteryTree"`
	Name                 string   `json:"name"`
	Ranks                int      `json:"ranks"`
	Image                Image    `json:"image"`
	SanitizedDescription []string `json:"sanitizedDescription"`
	ID                   int      `json:"id"`
	Description          []string `json:"description"`
}

type MasteryTree struct {
	Resolve  []MasteryTreeList `json:"Resolve"`
	Ferocity []MasteryTreeList `json:"Ferocity"`
	Cunning  []MasteryTreeList `json:"Cunning"`
}

type MasteryTreeList struct {
	MasteryTreeItems []MasteryTreeItem `json:"masteryTreeItem"`
}

type MasteryTreeItem struct {
	MasteryID int    `json:"masteryId"`
	Prereq    string `json:"prereq"`
}

type ProfileIcons struct {
	Data map[string]ProfileDetails `json:"data"`
	*VersionedData
}

type ProfileDetails struct {
	Image Image `json:"image"`
	ID    int64 `json:"id"`
}

type Realm struct {
	LG             string            `json:"lg"`
	DD             string            `json:"dd"`
	L              string            `json:"l"`
	N              map[string]string `json:"n"`
	ProfileIconMax int               `json:"profileiconmax"`
	Store          string            `json:"store"`
	V              string            `json:"v"`
	CDN            string            `json:"cdn"`
	CSS            string            `json:"css"`
}
