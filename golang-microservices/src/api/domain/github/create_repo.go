package github

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

type CreateRepoResponse struct {
	Id         int64          `json:"id"`
	Name       string         `json:"name"`
	FullName   string         `json:"full_name"`
	Owner      RepoOwner      `json:"owner"`
	Permission RepoPermission `json:"permission"`
}

type RepoOwner struct {
	Id      uint64 `json:"id"`
	Url     string `json:"url"`
	Login   string `json:"login"`
	HtmlUrl string `json:"html_url"`
}

type RepoPermission struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}
