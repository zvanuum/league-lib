package model

type LanguageStrings struct {
	Data map[string]string `json:"data"`
	*VersionedData
}
