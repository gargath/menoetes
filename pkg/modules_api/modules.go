package modules_api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	mux "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	store "github.com/gargath/menoetes/pkg/store"
)

type modulesAPI struct {
	store store.ModuleStore
}

func (api *modulesAPI) listHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Printf("Request for %s\n", r.URL.Path)
	vars := mux.Vars(r)
	queries := r.URL.Query()
	limit, _ := strconv.Atoi(queries.Get("limit"))
	if limit < 1 || limit > 10 {
		limit = 10
	}
	namespace, ok := vars["namespace"]
	if ok {
		name, ok := vars["name"]
		if ok {
			fmt.Fprintf(w, "hello, you've hit %s and will receive providers for module %s in namespace %s", r.URL.Path, name, namespace)
		} else {
			fmt.Fprintf(w, "hello, you've hit %s and will receive modules in namespace %s\n", r.URL.Path, namespace)
		}
	} else {
		fmt.Fprintf(w, "hello, you've hit %s and will receive ALL modules\n", r.URL.Path)
	}
	meta := &Metatype{
		Limit:          limit,
		Current_offset: 0,
		Next_url:       "bla",
	}
	list, _ := api.store.GetModulesList(vars["namespace"], vars["name"], vars["provider"])
	resp := &ListResponse{
		Meta:    *meta,
		Modules: list,
	}
	json.NewEncoder(w).Encode(resp)
}

func (api *modulesAPI) versionDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	version, ok := vars["version"]
	if !ok {
		v, err := api.store.GetLatestVersion(vars["namespace"], vars["name"], vars["provider"])
		if err != nil {
			http.Error(w, "Failed to retrieve module details. See logs for more...", http.StatusInternalServerError)
			log.Error(err)
		}
		version = v
	}
	details, err := api.store.GetModuleDetails(vars["namespace"], vars["name"], vars["provider"], version)
	if err != nil {
		http.Error(w, "Failed to retrieve module details. See logs for more...", http.StatusInternalServerError)
		log.Error(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		out, _ := json.Marshal(details)
		fmt.Fprintf(w, string(out))
	}
}

func (api *modulesAPI) latestDownloadHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Call to download redirector")
	vars := mux.Vars(r)
	latest, err := api.store.GetLatestVersion(vars["namespace"], vars["name"], vars["provider"])
	if err != nil {
		http.Error(w, "Failed to retrieve module details. See logs for more...", http.StatusInternalServerError)
		log.Error(err)
	} else {
		u := fmt.Sprintf("https://%s/v1/modules/%s/%s/%s/%s/download", r.Host, vars["namespace"], vars["name"], vars["provider"], latest)
		log.WithFields(log.Fields{"url": u}).Info("Redirecting to:")
		http.Redirect(w, r, u, 302)
	}
}

func (api *modulesAPI) downloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Info("Call to download handler")
	u, err := api.store.GetDownloadURL(vars["namespace"], vars["name"], vars["provider"], vars["version"])
	if err != nil {
		http.Error(w, "Failed to retrieve module download URL. See logs for more...", http.StatusInternalServerError)
		log.Error(err)
	}
	w.Header().Add("X-Terraform-Get", u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (api *modulesAPI) versionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Info("Call to versions handler")
	versions, err := api.store.GetModuleVersions(vars["namespace"], vars["name"], vars["provider"])
	if err != nil {
		http.Error(w, "Failed to retrieve module versions. See logs for more...", http.StatusInternalServerError)
		log.Error(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		mvs := make(map[string][]string)
		for _, v := range versions {
			fmt.Println(v)
			mvs[v.Version] = append(mvs[v.Version], v.Provider)
		}
		fmt.Println(mvs)
		resp := &VersionsResponse{}
		resp.Modules = append(resp.Modules, VersionsModulestype{
			Source: versions[0].Source,
		})
		for v := range mvs {
			n := &VersionsVersiontype{
				Version:    v,
				Submodules: make([]string, 0),
				Root: VersionsRoottype{
					Dependencies: make([]string, 0),
				},
			}
			for _, p := range mvs[v] {
				r := &VersionsProviderstype{
					Name:    p,
					Version: "",
				}
				n.Root.Providers = append(n.Root.Providers, *r)
			}
			resp.Modules[0].Versions = append(resp.Modules[0].Versions, *n)
		}
		json.NewEncoder(w).Encode(resp)
	}
}
