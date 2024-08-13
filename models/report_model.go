package models

type Report struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type ReportContent struct {
	Sections []ReportSection `json:"sections"`
}

type ReportSection struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
