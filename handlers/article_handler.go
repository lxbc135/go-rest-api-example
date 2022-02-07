package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lxbc135/go_rest_api_example/models"
)

func AllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: AllArticles")
	json.NewEncoder(w).Encode(models.TheModel.AllArticles())
}
