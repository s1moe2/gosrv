package server

import (
	"fmt"
	"net/http"

	"github.com/s1moe2/gosrv/config"
	"github.com/s1moe2/gosrv/db"
	"github.com/s1moe2/gosrv/handlers"
	"github.com/s1moe2/gosrv/repositories"
)

func Start() {
	conf := config.New()

	dbConn := db.ConnectDB(conf.Database)
	userRepo := repositories.NewUserRepo(dbConn)

	h := handlers.NewBaseHandler(userRepo)

	http.HandleFunc("/", h.HelloWorld)

	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "5000"),
	}

	s.ListenAndServe()

}
