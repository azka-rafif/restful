package handlers

import (
	"encoding/json"
	"go-mini/internal/domains/run"
	"go-mini/internal/response"
	"go-mini/internal/response/failure"
	"go-mini/internal/shared/pagination"
	"net/http"

	"github.com/go-chi/chi"
)

type RunHandler struct {
	Service *run.RunService
}

func NewRunHandler(service run.RunService) RunHandler {
	return RunHandler{Service: &service}
}

// HandlePost creates a new Run.
// @Summary Create a new Run.
// @Description This endpoint creates a new Run.
// @Tags Run
// @Param payload body run.RunPayload true "The Run to be created."
// @Produce json
// @Success 201 {object} response.Base{data=run.RunResponseFormat}
// @Failure 400 {object} response.Base
// @Failure 409 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/runs [post]
func (h *RunHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload run.RunPayload
	err := decoder.Decode(&payload)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	res, err := h.Service.Create(payload)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, res)
}

// HandleGetAll gets all Runs.
// @Summary Gets all Runs.
// @Description This endpoint gets all Runs.
// @Tags Run
// @Param page query int true "page number"
// @Param limit query int true "page limit"
// @Param sort query string false "sort direction"
// @Param field query string false "field to sort"
// @Param city query string false "city to filter"
// @Produce json
// @Success 200 {object} response.Base{data=[]run.RunResponseFormat}
// @Failure 400 {object} response.Base
// @Failure 409 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/runs [get]
func (h *RunHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	page, err := pagination.ConvertToInt(pagination.ParseQueryParams(r, "page"))
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	limit, err := pagination.ConvertToInt(pagination.ParseQueryParams(r, "limit"))
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	sort := pagination.GetSortDirection(pagination.ParseQueryParams(r, "sort"))
	field := pagination.CheckFieldQuery(pagination.ParseQueryParams(r, "field"), "id")
	city := pagination.ParseQueryParams(r, "city")
	offset := (page - 1) * limit
	res, err := h.Service.GetAll(offset, limit, field, sort, city)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}

// HandleUpdate updates a Run.
// @Summary updates a Run.
// @Description This endpoint updates a Run.
// @Tags Run
// @Param id path string true "run id"
// @Param payload body run.RunPayload true "payload"
// @Produce json
// @Success 200 {object} response.Base{data=run.RunResponseFormat}
// @Failure 400 {object} response.Base
// @Failure 409 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/runs/{id} [put]
func (h *RunHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	decoder := json.NewDecoder(r.Body)
	var payload run.RunPayload
	err := decoder.Decode(&payload)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	res, err := h.Service.Update(idString, payload)

	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, res)
}

// HandleDelete deletes a Run.
// @Summary deletes a Run.
// @Description This endpoint deletes a Run.
// @Tags Run
// @Param id path string true "run id"
// @Produce json
// @Success 204
// @Failure 400 {object} response.Base
// @Failure 409 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/runs/{id} [delete]
func (h *RunHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	err := h.Service.Delete(idString)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.NoContent(w)
}
