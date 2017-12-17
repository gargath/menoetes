package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	mux "github.com/gorilla/mux"
	"net/http"
)

func discoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Debug(vars["disco_path"])
	if (vars["disco_path"] == "terraform.json") {
		log.Printf("Service Discovery Request: %s", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\n\"modules.v1\": \"https://%s/v1/modules\"\n}", r.Host)
	} else {
		http.Error(w, "not here", http.StatusNotFound)
		return
	}
}
