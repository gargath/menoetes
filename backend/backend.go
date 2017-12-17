package backend

type Backend interface {
  GetModuleDetails(string, string, string, string) (string, error)
  GetLatestVersion(string, string, string) (string, error)
  GetDownloadURL(string, string, string, string) (string, error)
  GetModuleVersions(string, string, string) (string, error)
}
