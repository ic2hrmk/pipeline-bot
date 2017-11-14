package bitbucket

import "time"

const (
	PIPELINE_IN_PROGRESS = "INPROGRESS"
	PIPELINE_SUCCESSFUL = "SUCCESSFUL"
	PIPELINE_FAILURE = "FAILURE"
)

type WebHook struct {
	Actor struct {
		Username    string `json:"username"`
		Type        string `json:"type"`
		DisplayName string `json:"display_name"`
		UUID        string `json:"uuid"`
		Links       struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
	} `json:"actor"`
	Repository struct {
		Scm     string `json:"scm"`
		Website string `json:"website"`
		Name    string `json:"name"`
		Links   struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
		Project struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Type string `json:"type"`
			UUID string `json:"uuid"`
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"project"`
		FullName string `json:"full_name"`
		Owner    struct {
			Username    string `json:"username"`
			Type        string `json:"type"`
			DisplayName string `json:"display_name"`
			UUID        string `json:"uuid"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
		} `json:"owner"`
		Type      string `json:"type"`
		IsPrivate bool   `json:"is_private"`
		UUID      string `json:"uuid"`
	} `json:"repository"`
	CommitStatus struct {
		Links struct {
			Commit struct {
				Href string `json:"href"`
			} `json:"commit"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
		Repository struct {
			FullName string `json:"full_name"`
			Type     string `json:"type"`
			Name     string `json:"name"`
			Links    struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			UUID string `json:"uuid"`
		} `json:"repository"`
		URL       string    `json:"url"`
		Refname   string    `json:"refname"`
		Name      string    `json:"name"`
		State     string    `json:"state"`
		Key       string    `json:"key"`
		UpdatedOn time.Time `json:"updated_on"`
		Commit    struct {
			Hash  string `json:"hash"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Comments struct {
					Href string `json:"href"`
				} `json:"comments"`
				Patch struct {
					Href string `json:"href"`
				} `json:"patch"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Diff struct {
					Href string `json:"href"`
				} `json:"diff"`
				Approve struct {
					Href string `json:"href"`
				} `json:"approve"`
				Statuses struct {
					Href string `json:"href"`
				} `json:"statuses"`
			} `json:"links"`
			Author struct {
				Raw  string `json:"raw"`
				Type string `json:"type"`
				User struct {
					Username    string `json:"username"`
					Type        string `json:"type"`
					DisplayName string `json:"display_name"`
					UUID        string `json:"uuid"`
					Links       struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
						HTML struct {
							Href string `json:"href"`
						} `json:"html"`
						Avatar struct {
							Href string `json:"href"`
						} `json:"avatar"`
					} `json:"links"`
				} `json:"user"`
			} `json:"author"`
			Date    time.Time `json:"date"`
			Message string    `json:"message"`
			Type    string    `json:"type"`
		} `json:"commit"`
		Type        string    `json:"type"`
		CreatedOn   time.Time `json:"created_on"`
		Description string    `json:"description"`
	} `json:"commit_status"`
}