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

// GetSubscription godoc
// @Summary Get a subscription by ID
// @Description Get details of a subscription by its ID
// @ID get-subscription
// @Produce  json
// @Param id path string true "Subscription ID (UUID)"
// @Success 200 {object} database.Subscription
// @Router /subscription/{id} [get]
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

// CreateSubscription godoc
// @Summary Create a new subscription
// @Description Create a new subscription with the provided details
// @ID create-subscription
// @Accept  json
// @Produce  json
// @Param request body handlers.CreateSubscriptionRequest true "Subscription details"
// @Success 200 {object} database.Subscription
// @Router /subscription [post]
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

// ChangeSubscription godoc
// @Summary Change an existing subscription
// @Description Change an existing subscription with the provided details
// @ID change-subscription
// @Accept  json
// @Produce  json
// @Param id path string true "Subscription ID (UUID)"
// @Param request body handlers.ChangeSubscriptionRequest true "Subscription details"
// @Success 200 {object} database.Subscription
// @Router /subscription/{id} [put]
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

// DeleteSubscription godoc
// @Summary Delete a subscription
// @Description Delete a subscription with the provided ID
// @ID delete-subscription
// @Produce  json
// @Param id path string true "Subscription ID (UUID)"
// @Success 200 {object} handlers.DeleteSubscriptionResponse
// @Router /subscription/{id} [delete]
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

// TotalSum godoc
// @Summary Get total sum of subscriptions
// @Description Get total sum of subscriptions with the provided start date
// @ID total-sum
// @Accept  json
// @Produce  json
// @Param request body handlers.TotalSumRequest true "Total sum details"
// @Success 200 {object} handlers.TotalSumResponse
// @Router /subscription/sum [get]
func (cfg *Config) TotalSum(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		StartDate string `json:"start_date"`
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

type CreateSubscriptionRequest struct {
	ServiceName string    `json:"service_name"`
	PriceRub    int       `json:"price_rub"`
	UserID      uuid.UUID `json:"user_id"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
}

type ChangeSubscriptionRequest struct {
	ServiceName string `json:"service_name"`
}

type DeleteSubscriptionResponse struct {
	Status string `json:"status"`
}

type TotalSumRequest struct {
	StartDate string `json:"start_date"`
}

type TotalSumResponse struct {
	TotalSum int `json:"total_sum"`
}
