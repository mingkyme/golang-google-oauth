package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"example.com/session"
	"example.com/util"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type googleResponeBody struct {
	Email string "json:email"
}

func Setup_GoogleOauth(app *fiber.App) {
	googleOauthConfig := &oauth2.Config{
		RedirectURL:  "http://localhost/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	OAuth.Google = *googleOauthConfig
}
func GoogleLogin(c *fiber.Ctx) error {
	state := util.RandomString(16)
	cookie := new(fiber.Cookie)
	cookie.Name = "googleState"
	cookie.Value = state
	cookie.Expires = time.Now().Add(10 * time.Minute)
	c.Cookie(cookie)
	url := OAuth.Google.AuthCodeURL(state)
	return c.Redirect(url)
}
func GoogleCallback(c *fiber.Ctx) error {
	if c.FormValue("state") != c.Cookies("googleState") {
		fmt.Println("state is not vaild")
		return c.Redirect("/")
	}
	token, err := OAuth.Google.Exchange(c.Context(), c.FormValue("code"))
	if err != nil {
		fmt.Println("could not get token")
		return c.Redirect("/")
	}

	respone, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Println("could not create get method")
		return c.Redirect("/")
	}
	defer respone.Body.Close()

	body, err := ioutil.ReadAll(respone.Body)
	if err != nil {
		return c.Redirect("/")
	}

	googleResponeJson := googleResponeBody{}
	json.Unmarshal(body, &googleResponeJson)
	session, err := session.Store.Get(c)
	if err != nil {
		return c.Redirect("/")
	}
	session.Set("email", googleResponeJson.Email)
	session.Save()

	return c.Redirect("/")
}
