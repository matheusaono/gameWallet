package main

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	Model *Model
}

type Model struct {
	Games []struct {
		Title    string `json:"title"`
		Platform string `json:"platform"`
	}
}

func main() {
	// Cria o modelo
	model := &Model{
		Games: []struct {
			Title    string `json:"title"`
			Platform string `json:"platform"`
		}{
			{Title: "Elden Ring", Platform: "PC"},
			{Title: "God of War Ragnarok", Platform: "PS5"},
			{Title: "Starfield", Platform: "PC, Xbox Series X/S"},
		},
	}

	// Cria o controlador
	controller := &Controller{Model: model}

	// Registra o roteador
	http.HandleFunc("/games", controller.GetGames)

	// Inicia o servidor
	http.ListenAndServe(":8080", nil)
}

func (c *Controller) GetGames(w http.ResponseWriter, r *http.Request) {
	// Obtém os jogos
	games := c.Model.Games

	// Converte os jogos para JSON
	json, err := json.Marshal(games)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Envia a resposta
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
