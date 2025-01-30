package domain

type SettingID int64

type Content struct {
	Content string `json:"content"`
}

type Setting struct {
	SettingID SettingID
	Content   Content
	Title     string
}
