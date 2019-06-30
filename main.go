package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"strings"
)

func BasicAuthMiddleWare(user string, password string, handler http.Handler) http.Handler {
	if user == "" && password == "" {
		return handler
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqUser, reqPassword, ok := r.BasicAuth()
		if ok && reqUser == user && reqPassword == password {
			handler.ServeHTTP(w, r)
			return
		}
		w.Header().Set("WWW-Authenticate","Basic")
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, http.StatusText(http.StatusUnauthorized))
	})
}

func main() {
	var (
		path         string
		port         string
		basicAuthKey string
	)

	flag.StringVar(&path, "d", ".", "ファイルサーバのルートディレクトリ")
	flag.StringVar(&port, "p", "8080", "割り当てるポート")
	flag.StringVar(&basicAuthKey, "b", ":", "Basic認証を使用する[user:password]")
	flag.Parse()

	key := strings.Split(basicAuthKey, ":")
	if len(key) != 2 {
		flag.Usage()
		return
	}

	server := http.Server{
		Addr: ":" + port,
	}

	http.Handle("/",
		BasicAuthMiddleWare(key[0], key[1], http.StripPrefix("/", http.FileServer(http.Dir(path)))),
	)

	log.Fatal(server.ListenAndServe())
}
