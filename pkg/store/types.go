package store

import (
	"time"
)

type Store interface {
	GetModuleDetails(string, string, string, string) (Module, error)
	GetLatestVersion(string, string, string) (string, error)
	GetDownloadURL(string, string, string, string) (string, error)
	GetModuleVersions(string, string, string) ([]Module, error)
	GetModulesList(string, string, string) ([]Module, error)
	ValidateAccessToken(string) (string, error)
}

type Module struct {
	Id           string    `json:"id"`
	Owner        string    `json:"owner"`
	Namespace    string    `json:"namespace"`
	Name         string    `json:"name"`
	Version      string    `json:"version"`
	Provider     string    `json:"provider"`
	Description  string    `json:"description"`
	Source       string    `json:"source"`
	Published_at time.Time `json:"published_at"`
	Downloads    int       `json:"downloads"`
	Verified     bool      `json:"verified"`
}
