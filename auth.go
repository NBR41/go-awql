package awql

import (
	"net/http"
)

// Auth struct for authenticated calls infos
type Auth struct {
	AdwordsID      string
	DeveloperToken string
	Client         *http.Client
}

// NewAuth returns new instance of Auth
func NewAuth(adwordsID string, devToken string, cli *http.Client) *Auth {
	return &Auth{
		AdwordsID:      adwordsID,
		DeveloperToken: devToken,
		Client:         cli,
	}
}
