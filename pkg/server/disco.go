package server

import (
	"fmt"
	"net/http"

	mux "github.com/gorilla/mux"
)

func discoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["disco_path"] == "terraform.json" {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\n\"modules.v1\": \"https://%s/v1/modules\"\n}", r.Host)
	} else {
		http.Error(w, "not here", http.StatusNotFound)
		return
	}
}
