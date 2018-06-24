package main

import (
	"flag"
	"net/http"
)

func main() {
	var (
		path string
		port string
	)

	flag.StringVar(&path, "d", ".", "ファイルサーバのルートディレクトリ")
	flag.StringVar(&port, "p", "8080", "割り当てるポート")
	flag.Parse()

	server := http.Server{
		Addr: ":" + port,
	}

	http.Handle("/", http.StripPrefix("/",
		http.FileServer(http.Dir(path))))

	server.ListenAndServe()
}
