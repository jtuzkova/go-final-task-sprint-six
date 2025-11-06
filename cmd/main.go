package main

import (
	"log"
	"os"
	s "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	flog, err := os.OpenFile(`server.log`, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }

    defer flog.Close()

    mylog := log.New(flog, `serv `, log.LstdFlags|log.Lshortfile)

	srv := s.CreateRouter(mylog)
	srv.Logger.Println(`Start server`)

	if err := srv.Server.ListenAndServe(); err != nil {
		srv.Logger.Fatalf("Server failed to start: %v", err)
	}
}
