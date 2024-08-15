package models

import (
	"encoding/json"
	"errors"
)

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

func (r *Report) Validate() error {
	if r.Title == "" {
		return errors.New("title is required")
	}

	if r.Author == "" {
		return errors.New("author is required")
	}

	if r.Content == "" {
		return errors.New("content is required")
	}

	return nil
}

func (r *Report) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Report) FromJSON(data []byte) error {
	return json.Unmarshal(data, r)
}
