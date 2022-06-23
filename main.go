package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type InfoApp struct {
	NameAPP string // `json:"name_app"`
	Url     string // `json:"url"`
}

type ResponseInfoApps struct {
	NameAPP   string `json:"name_app"`
	Url       string `json:"url"`
	Unhealthy bool   `json:"unhealthy"`
}

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		apis := []InfoApp{
			{
				NameAPP: "CIA",
				Url:     "https://mad.smarters.io/health-check",
			},
			{
				NameAPP: "FBI",
				Url:     "https://mad-sock.smarters.io/health-check",
			},
		}

		var responseInfos []ResponseInfoApps
		for _, api := range apis {
			resp, err := http.Get(api.Url)

			if err != nil {
				fmt.Println("Erro", err)
			}

			var unhealthy bool

			if resp.StatusCode == 200 {
				unhealthy = false
			} else {
				unhealthy = true
			}

			res := ResponseInfoApps{
				NameAPP:   api.NameAPP,
				Url:       api.Url,
				Unhealthy: unhealthy,
			}

			responseInfos = append(responseInfos, res)
		}

		json.NewEncoder(w).Encode(responseInfos)
	})
	http.ListenAndServe(":8080", r)
}
