package auth

import "golang.org/x/oauth2"

var (
	OAuth oAuth
)

type oAuth struct {
	Google oauth2.Config
}
