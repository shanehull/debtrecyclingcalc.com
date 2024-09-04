package handlers

import (
	"net/http"

	"debtrecyclingcalc.com/internal/buildinfo"
	"debtrecyclingcalc.com/internal/templates"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)

		c := templates.NotFound()
		err := templates.Layout(c, "Not Fountempl.WithStatus(http.StatusNotFound)d", buildinfo.GitTag, buildinfo.BuildDate).
			Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	index := templates.Index(
		templates.Hero(),
		templates.Form(),
		templates.Blank(),
	)

	err := templates.Layout(index, "Debt Recycling Calculator", buildinfo.GitTag, buildinfo.BuildDate).
		Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
