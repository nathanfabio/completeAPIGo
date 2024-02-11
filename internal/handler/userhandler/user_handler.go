package userhandler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nathanfabio/completeAPIGo/internal/dto"
	"github.com/nathanfabio/completeAPIGo/internal/handler/httperr"
	"github.com/nathanfabio/completeAPIGo/internal/handler/validation"
	"modernc.org/libc/uuid/uuid"
)

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserDto

	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}

	httperr := validation.ValidateHttpData(req)
	if httperr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httperr), slog.String("package", "handler_user"))
		w.WriteHeader(httperr.Code)
		json.NewEncoder(w).Encode(httperr)
		return
	}

	err = h.service.CreateUser(r.Context(), req)
	if err != nil {
		slog.Error(fmt.Sprintf("error to create user: %v", err), slog.String("package", "handler_user"))
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to create user")
		json.NewEncoder(w).Encode(msg)
		return
	}
}


func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateUserDto

	id := chi.URLParam(r, "id")
	if id == "" {
		slog.Error("id is empty", slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("id is required")
		json.NewEncoder(w).Encode(msg)
		return
	}

	_, err := uuid.Parse(id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to parse id: %v", err), slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to parse id")
		json.NewEncoder(w).Encode(msg)
		return
	}

	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}	

	httperr := validation.ValidateHttpData(req)
	if httperr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httperr), slog.String("package", "handler_user"))
		w.WriteHeader(httperr.Code)
		json.NewEncoder(w).Encode(httperr)
		return
	}

	err = h.service.UpdateUser(r.Context(), req, id)
	if err != nil {
		slog.Error(fmt.Sprintf("error to update user: %v", err), slog.String("package", "handler_user"))
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to update user")
		json.NewEncoder(w).Encode(msg)
		return
	}
}