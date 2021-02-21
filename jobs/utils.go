package jobs

import (
	"net/url"
)

// GetProviderFromURI ...
func GetProviderFromURI(uri string) string {
	u, _ := url.Parse(uri)
	return u.Scheme
}
