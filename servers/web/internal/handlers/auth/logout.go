package auth

import (
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"zotes/servers/web/internal/app"
)

func Logout(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_CLIENT_UI_DOMAIN") + "/v2/logout/")
		if err != nil {
			app.Logger.Error("Logout", slog.Any("error", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// session := app.GetRedisSessionManager()
		app.Session.DeleteAuthUser(r.Context())
		// app.Session().Destroy(r.Context())

		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		returnTo, err := url.Parse(scheme + "://" + r.Host)
		if err != nil {
			app.Logger.Error("Logout", slog.Any("error", err.Error()))
			// 500
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		parameters := url.Values{}
		parameters.Add("returnTo", returnTo.String())
		parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_UI_CLIENT_ID"))
		logoutUrl.RawQuery = parameters.Encode()

		http.Redirect(w, r, logoutUrl.String(), http.StatusSeeOther)
	}
}
