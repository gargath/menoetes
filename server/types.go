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
