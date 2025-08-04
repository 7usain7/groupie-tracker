package funcs

import "net/http"

func AboutHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		RenderTemplate(w, "error.html", MethodNotAllowed)
		return
	}

	RenderTemplate(w, "about.html", nil)

}
