package product

import (
	"myapi/service/auth"
	"myapi/types"
	"myapi/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	types.IProductStore
	types.IUserStore
}

func NewHandler(store types.IProductStore, userStore types.IUserStore) *Handler {
	return &Handler{
		IProductStore: store,
		IUserStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products",auth.WithJWTAuth(h.getProducts,h.IUserStore)).Methods("GET")
}

func (h *Handler) getProducts(response http.ResponseWriter, request *http.Request) {

	products, err := h.IProductStore.GetProducts()
	if err != nil {
		utils.WriteError(response, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(response, http.StatusOK, products)
}
