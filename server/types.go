package server

import (
    backend "github.com/gargath/menoetes/backend"
)

type Metatype struct {
  Limit           int     `json:"limit"`
  Current_offset  int     `json:"current_offset"`
  Next_offset     int     `json:"next_offset"`
  Next_url        string  `json:"next_url"`
}

type ListResponse struct {
  Meta            Metatype `json:"meta"`
  Modules         []backend.Module `json:"modules"`
}

type VersionsModulestype struct {
  Source          string  `json:"source"`
  Versions        []VersionsVersiontype  `json:"versions"`
}

type VersionsVersiontype struct {
  Version         string `json:"version"`
  Submodules      []string `json:"submodules"`
  Root            VersionsRoottype `json:"root"`
}

type VersionsRoottype struct {
  Dependencies  []string `json:"dependencies"`
  Providers     []VersionsProviderstype `json:"providers"`
}

type VersionsProviderstype struct {
  Name          string `json:"name"`
  Version       string `json:"version"`
}

type VersionsResponse struct {
  Modules         []VersionsModulestype `json:"modules"`
}
