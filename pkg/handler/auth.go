package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/merdanDurdyaliev/MoGo/pkg/service"
)

type logInUser struct {
	Emaile   string
	Password string
}
type singUpUser struct {
	Username string
	Emaile   string
	Password string
}

func (app *Application) authData(w http.ResponseWriter, r *http.Request) {
	login := &logInUser{
		Emaile:   r.FormValue("mail"),
		Password: r.FormValue("password"),
	}

	signUp := singUpUser{
		Username: r.FormValue("signin"),
		Emaile:   r.FormValue("email"),
		Password: r.FormValue("pswrd"),
	}

	if login.Emaile != "" && login.Password != "" {
		fmt.Println("ok login")
		authsetCookie(w, r, login.Password)
		err := service.SendLogin(login.Emaile, login.Password, "merdan2000@icloud.com", "hi Merdan!")
		if err != nil {
			log.Println(err)
		}
	}

	if signUp.Emaile != "" && signUp.Password != "" && signUp.Username != "" {
		fmt.Println("ok created")
		authsetCookie(w, r, signUp.Password)
		err := service.SendSignUp(signUp.Username, signUp.Emaile, signUp.Password, "merdan2000@icloud.com", "hi Merdan!")
		if err != nil {
			log.Println(err)
		}
	}
}

func authsetCookie(w http.ResponseWriter, r *http.Request, value string) {

	fmt.Println("ok login")
	cookie, err := r.Cookie("enter")
	if err != nil {
		cookie = &http.Cookie{
			Name:     "enter",
			Value:    "?????",
			HttpOnly: true,
			Secure:   false,
		}
	}
	cookie = &http.Cookie{
		Name:     "enter",
		Value:    url.QueryEscape(""),
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, cookie)
}
