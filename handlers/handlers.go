package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/imhasandl/subscription-manager/database"
	"github.com/imhasandl/subscription-manager/utils"
)

type Config struct {
	db *database.DB
}

func NewConfig(db *database.DB) *Config {
	return &Config{
		db: db,
	}
}

func (cfg *Config) GetSubscription(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "can't parse id", err)
		return
	}

	sub, err := cfg.db.GetSubscription(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "can't get subscription", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, sub)
}

func (cfg *Config) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		ServiceName string    `json:"service_name"`
		PriceRub    int       `json:"price_rub"`
		UserID      uuid.UUID `json:"user_id"`
		StartDate   string    `json:"start_date"`
		EndDate     string    `json:"end_date"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "can't decode request", err)
		return
	}

	subscription, err := cfg.db.SaveSubscription(database.SaveSubscriptionParams{
		ID:          uuid.New(),
		ServiceName: requestData.ServiceName,
		PriceRub:    requestData.PriceRub,
		UserID:      requestData.UserID,
		StartDate:   requestData.StartDate,
		EndDate:     requestData.EndDate,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "can't save subscription", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, subscription)
}

func (cfg *Config) ChangeSubscription(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		ServiceName string `json:"service_name"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "can't decode request", err)
		return
	}

	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "can't parse id", err)
		return
	}

	sub, err := cfg.db.ChangeSubscription(id, requestData.ServiceName)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "can't change subscription", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, sub)
}

func (cfg *Config) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "can't parse id", err)
		return
	}

	err = cfg.db.DeleteSubscription(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "can't delete subscription", err)
		return
	}

	response := map[string]interface{}{
		"status": "success",
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

func (cfg *Config) TotalSum(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		StartDate   string    `json:"start_date"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "can't decode request", err)
		return
	}

	sum, err := cfg.db.TotalSumSubscriptions(requestData.StartDate)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "can't get total sum  of subscriptions", err)
		return
	}

	response := map[string]interface{}{
		"total_sum": sum,
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}
