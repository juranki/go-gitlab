package gitlab

import (
	"fmt"
	"net/url"
	"time"
)

// ReleasesService handles communication with the releases methods
// of the GitLab API.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/releases/index.html
type ReleasesService struct {
	client *Client
}

// Release represents a project release.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/releases/index.html#list-releases
type Release struct {
	TagName         string     `json:"tag_name"`
	Name            string     `json:"name"`
	Description     string     `json:"description,omitempty"`
	DescriptionHTML string     `json:"description_html,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	Author          *struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"author,omitempty"`
	Commit *Commit `json:"commit,omitempty"`
	Assets *struct {
		Count   int `json:"count"`
		Sources []struct {
			Format string `json:"format"`
			URL    string `json:"url"`
		} `json:"sources"`
		Links []ReleaseLink `json:"links"`
	} `json:"assets,omitempty"`
}

// ListReleasesOptions represents ListReleases() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/releases/index.html#list-releases
type ListReleasesOptions ListOptions

// ListReleases gets a pagenated of releases accessible by the authenticated user.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/releases/index.html#list-releases
func (s *ReleasesService) ListReleases(pid interface{}, opt *ListReleasesOptions, options ...OptionFunc) ([]*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var r []*Release
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, err
}

// GetRelease returns a single release, identified by a tag name.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/releases/index.html#get-a-release-by-a-tag-name
func (s *ReleasesService) GetRelease(pid interface{}, tag string, options ...OptionFunc) (*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases/%s", url.QueryEscape(project), tag)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var r *Release
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, err
}

// ReleaseAssets represents release assets in CreateRelease() options
//
// GitLab API docs: https://docs.gitlab.com/ce/api/releases/index.html#create-a-release
type ReleaseAssets struct {
	Links []ReleaseAssetLink `url:"links" json:"links"`
}

// ReleaseAssetLink represents release asset link in CreateRelease() options
//
// GitLab API docs: https://docs.gitlab.com/ce/api/releases/index.html#create-a-release
type ReleaseAssetLink struct {
	Name string `url:"name" json:"name"`
	URL  string `url:"ref" json:"ref"`
}

// CreateReleaseOptions represents CreateRelease() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/releases/index.html#create-a-release
type CreateReleaseOptions struct {
	Name        string         `url:"name" json:"name"`
	TagName     string         `url:"tag_name" json:"tag_name"`
	Description string         `url:"description" json:"description"`
	Ref         string         `url:"ref,omitempty" json:"ref,omitempty"`
	Assets      *ReleaseAssets `url:"assets,omitempty" json:"assets,omitempty"`
}

// CreateRelease creates a release.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/releases/index.html#create-a-release
func (s *ReleasesService) CreateRelease(pid interface{}, opts *CreateReleaseOptions, options ...OptionFunc) (*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases", url.QueryEscape(project))

	req, err := s.client.NewRequest("POST", u, opts, options)
	if err != nil {
		return nil, nil, err
	}

	var r *Release
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, err
}

// UpdateReleaseOptions represents UpdateRelease() options.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/releases/index.html#update-a-release
type UpdateReleaseOptions struct {
	Name        string `url:"name" json:"name"`
	Description string `url:"description" json:"description"`
}

// UpdateRelease updates a release.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/releases/index.html#update-a-release
func (s *ReleasesService) UpdateRelease(pid interface{}, tag string, opts *UpdateReleaseOptions, options ...OptionFunc) (*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases/%s", url.QueryEscape(project), tag)

	req, err := s.client.NewRequest("PUT", u, opts, options)
	if err != nil {
		return nil, nil, err
	}

	var r *Release
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, err
}

// DeleteRelease deletes a release.
//
// GitLab API docs: https://docs.gitlab.com/ce/api/releases/index.html#delete-a-release
func (s *ReleasesService) DeleteRelease(pid interface{}, tag string, options ...OptionFunc) (*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases/%s", url.QueryEscape(project), tag)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var r *Release
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, err
}
