package store

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ModuleStore interface {
	GetModuleDetails(string, string, string, string) (Module, error)
	GetLatestVersion(string, string, string) (string, error)
	GetDownloadURL(string, string, string, string) (string, error)
	GetModuleVersions(string, string, string) ([]Module, error)
	GetModulesList(string, string, string) ([]Module, error)
}

type Module struct {
	Id           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Owner        string        `json:"owner"`
	Namespace    string        `json:"namespace"`
	Name         string        `json:"name"`
	Version      string        `json:"version"`
	Provider     string        `json:"provider"`
	Description  string        `json:"description"`
	Source       string        `json:"source"`
	Published_at time.Time     `json:"published_at"`
	Downloads    int           `json:"downloads"`
	Verified     bool          `json:"verified"`
}
