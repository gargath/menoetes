package server

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	mux "github.com/gorilla/mux"
	"net/http"
	"encoding/json"

  backend "github.com/gargath/menoetes/backend"
)

type moduleAPI struct {
  backend backend.Backend
}

func (api *moduleAPI) listHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
  	log.Printf("Request for %s\n", r.URL.Path)
    vars := mux.Vars(r)
    namespace, ok := vars["namespace"]
    if (ok) {
      name, ok := vars["name"]
      if (ok) {
        fmt.Fprintf(w, "hello, you've hit %s and will receive providers for module %s in namespace %s", r.URL.Path, name, namespace)
      } else {
        fmt.Fprintf(w, "hello, you've hit %s and will receive modules in namespace %s\n", r.URL.Path, namespace)
      }
    } else {
      fmt.Fprintf(w, "hello, you've hit %s and will receive ALL modules\n", r.URL.Path)
    }
		return
}

func (api *moduleAPI) versionDetailsHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  version, ok := vars["version"]
  if (!ok) {
    v, err := api.backend.GetLatestVersion(vars["namespace"], vars["name"], vars["provider"])
    if err != nil {
      http.Error(w, "Failed to retrieve module details. See logs for more...", http.StatusInternalServerError)
      log.Error(err)
    }
    version = v
  }
  details, err := api.backend.GetModuleDetails(vars["namespace"], vars["name"], vars["provider"], version)
  if err != nil {
    http.Error(w, "Failed to retrieve module details. See logs for more...", http.StatusInternalServerError)
    log.Error(err)
  } else {
    w.Header().Set("Content-Type", "application/json")
		out, _ := json.Marshal(details)
    fmt.Fprintf(w, string(out))
  }
}

func (api *moduleAPI) latestDownloadHandler(w http.ResponseWriter, r *http.Request) {
  log.Info("Call to download redirector")
  vars := mux.Vars(r)
  latest, err := api.backend.GetLatestVersion(vars["namespace"], vars["name"], vars["provider"])
  if err != nil {
    http.Error(w, "Failed to retrieve module details. See logs for more...", http.StatusInternalServerError)
    log.Error(err)
  } else {
    u := fmt.Sprintf("https://%s/v1/modules/%s/%s/%s/%s/download", r.Host, vars["namespace"], vars["name"], vars["provider"], latest)
    log.WithFields(log.Fields{"url": u}).Info("Redirecting to:")
    http.Redirect(w, r, u, 302)
  }
}

func (api *moduleAPI) downloadHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  log.Info("Call to download handler")
  u, err := api.backend.GetDownloadURL(vars["namespace"], vars["name"], vars["provider"], vars["version"])
  if err != nil {
    http.Error(w, "Failed to retrieve module download URL. See logs for more...", http.StatusInternalServerError)
    log.Error(err)
  }
  w.Header().Add("X-Terraform-Get", u)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusNoContent)
}

func (api *moduleAPI) versionsHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  log.Info("Call to versions handler")
  versions, err := api.backend.GetModuleVersions(vars["namespace"], vars["name"], vars["provider"])
  if err != nil {
    http.Error(w, "Failed to retrieve module versions. See logs for more...", http.StatusInternalServerError)
    log.Error(err)
  } else {
    w.Header().Set("Content-Type", "application/json")
		out, _ := json.Marshal(versions)
    fmt.Fprintf(w, string(out))
  }
}

func registerModulesAPI(r *mux.Router) {
  api := &moduleAPI{backend: backend.NewMockBackend()}
  s := r.PathPrefix("/v1/modules").Subrouter()
  s.HandleFunc("", use(api.listHandler, tokenAuth)).Methods("GET")
  s.HandleFunc("/", use(api.listHandler, tokenAuth)).Methods("GET")
  s.HandleFunc("/{namespace}", use(api.listHandler, tokenAuth)).Methods("GET")
  s.HandleFunc("/{namespace}/{name}", use(api.listHandler, tokenAuth)).Methods("GET")

  s.HandleFunc("/{namespace}/{name}/{provider}/versions", use(api.versionsHandler, tokenAuth)).Methods("GET")
  s.HandleFunc("/{namespace}/{name}/{provider}/{version}", use(api.versionDetailsHandler, tokenAuth)).Methods("GET")
  s.HandleFunc("/{namespace}/{name}/{provider}", use(api.versionDetailsHandler, tokenAuth)).Methods("GET")

  s.HandleFunc("/{namespace}/{name}/{provider}/download", use(api.latestDownloadHandler, tokenAuth)).Methods("GET")

  s.HandleFunc("/{namespace}/{name}/{provider}/{version}/download", use(api.downloadHandler, tokenAuth)).Methods("GET")

}
