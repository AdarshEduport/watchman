package user

import (
	"fmt"
	"myapi/service/auth"
	"myapi/types"
	"myapi/utils"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sikozonpc/ecom/configs"
)

type Handler struct {
	types.IUserStore
}

func NewHandler(store types.IUserStore) *Handler {
	return &Handler{
		IUserStore: store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.login).Methods("POST")
	router.HandleFunc("/register", h.register).Methods("POST")
}

func (h *Handler) login(response http.ResponseWriter, request *http.Request) {
	// Get the json payload
	var payload types.AuthUserPayload
	if err := utils.ParseJSON(request, &payload); err != nil {
		utils.WriteError(response, http.StatusBadRequest, err)
		return
	}
	//Validate the payload
	if err:=utils.Validate.Struct(payload); err!=nil{
		utils.WriteError(response,http.StatusBadRequest,err)
		return
	}

	// Get user
	user,err:= h.GetUserByEmail(payload.Email)
	if err!=nil{
		utils.WriteError(response,http.StatusNotFound,err)
		return
	}
	// Compare password 
	if !auth.ComparePassword(user.Password,payload.Password){
		utils.WriteError(response,http.StatusBadRequest,fmt.Errorf("invalid credentials"))
		return
	}
	// Generate token
	token,err:=auth.CreateJWT([]byte(configs.Envs.JWTSecret),user.Id)
	if err!=nil{
		utils.WriteError(response,http.StatusInternalServerError,err)
		return
	}
	utils.WriteJSON(response,http.StatusOK,map[string]string{"token": token})
}

func (h *Handler) register(writer http.ResponseWriter, request *http.Request) {

	// Get the json payload
	var payload types.AuthUserPayload
	if err := utils.ParseJSON(request, &payload); err != nil {
		utils.WriteError(writer, http.StatusBadRequest, err)
		return
	}

	// validate payload

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(writer, http.StatusBadRequest, err)
		return
	}

	// check if the user already registered or not
	_, err := h.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(writer, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// hash the password and save into database
	hashedPassword, err := auth.HashPasswod(payload.Password)
	if err != nil {
		utils.WriteError(writer, http.StatusInternalServerError, err)
		return
	}
	// if the user is not registered then register the user
	err = h.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(writer, http.StatusCreated, "Account created successfully")
}
