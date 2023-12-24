package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/goccy/go-json"
)

//go:embed data/quotes.json web
var embeddedData embed.FS

type QuotesData struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
}

func main() {
	var serverPort uint
	flag.UintVar(&serverPort, "port", 8080, "Server port")
	flag.Parse()

	var quotes []QuotesData

	if data, err := embeddedData.ReadFile("data/quotes.json"); err != nil {
		log.Fatalln(err)
	} else {
		if err := json.Unmarshal(data, &quotes); err != nil {
			log.Fatalln(err)
		}
	}

	webFS, err := fs.Sub(embeddedData, "web")
	if err != nil {
		log.Fatalln(err)
	}

	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(webFS))))
	// mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("web"))))

	mux.Handle("/quotes", allowedHTTPMethod(quotesHandler(quotes), http.MethodGet))

	log.Println("Server listening to port", serverPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), mux); err != nil {
		log.Fatalln(err)
	}
}

func quotesHandler(quotes []QuotesData) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		countParam := r.URL.Query().Get("count")

		if countParam == "" || countParam == "0" {
			randNum := rand.Intn(len(quotes))
			data, err := json.Marshal(quotes[randNum])
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal server error: " + err.Error()))
				return
			}

			w.Header().Add("Content-Type", "application/json")
			w.Write(data)

			return
		}

		quoteCount, err := strconv.ParseUint(countParam, 10, 64)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request: invalid value '" + countParam + "' for count parameter"))
			return
		}

		data, err := json.Marshal(quotes[:quoteCount])
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error: " + err.Error()))
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	})
}

func allowedHTTPMethod(next http.Handler, methods ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, m := range methods {
			if r.Method == m {
				next.ServeHTTP(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})
}
