package models

type Controller struct {
	Model *Model
}

type Model struct {
	Games []struct {
		Title    string `json:"title"`
		Platform string `json:"platform"`
	}
}
