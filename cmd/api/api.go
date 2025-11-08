package api

import (
	"database/sql"
	"log"
	"myapi/service/product"
	"myapi/service/user"
	"myapi/utils"
	"net/http"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db}
}
func (s *ApiServer) Run() error{
    router := mux.NewRouter()
	subrouter :=router.PathPrefix("/api/v1").Subrouter()

	// User
	store := user.NewStore(s.db)
	handler :=user.NewHandler(store)
	handler.RegisterRoutes(subrouter)

	// Product
	productStore:= product.NewStore(s.db)
	productHandler :=product.NewHandler(productStore,store)
	productHandler.RegisterRoutes(subrouter)


	log.Printf("Listening on http://localhost%s",s.addr)
	utils.PrintRoutes(router)
	return  http.ListenAndServe(s.addr,router);
}
