package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"todolist-api/helpers"
	"todolist-api/services"
)

type TodoController struct {
	service *services.TodoService
}

func NewTodoController(service *services.TodoService) *TodoController {
	return &TodoController{service: service}
}

func (c *TodoController) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos, err := c.service.GetAll()
	if err != nil {
		helpers.Error(w, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	helpers.Success(w, http.StatusOK, "Success", todos)
}

func (c *TodoController) GetByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid ID", []string{"id must be a valid integer"})
		return
	}

	todo, err := c.service.GetByID(id)
	if err != nil {
		if err.Error() == "todo not found" {
			helpers.Error(w, http.StatusNotFound, "Not Found", []string{err.Error()})
			return
		}
		helpers.Error(w, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	helpers.Success(w, http.StatusOK, "Success", todo)
}

func (c *TodoController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid JSON", []string{"invalid request body"})
		return
	}

	todo, err := c.service.Create(req.Title, req.Description)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Validation Error", []string{err.Error()})
		return
	}

	helpers.Success(w, http.StatusCreated, "Created", todo)
}

func (c *TodoController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid ID", []string{"id must be a valid integer"})
		return
	}

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid JSON", []string{"invalid request body"})
		return
	}

	todo, err := c.service.Update(id, req.Title, req.Description, req.Completed)
	if err != nil {
		if err.Error() == "todo not found" {
			helpers.Error(w, http.StatusNotFound, "Not Found", []string{err.Error()})
			return
		}
		helpers.Error(w, http.StatusBadRequest, "Validation Error", []string{err.Error()})
		return
	}

	helpers.Success(w, http.StatusOK, "Updated", todo)
}

func (c *TodoController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid ID", []string{"id must be a valid integer"})
		return
	}

	if err := c.service.Delete(id); err != nil {
		if err.Error() == "todo not found" {
			helpers.Error(w, http.StatusNotFound, "Not Found", []string{err.Error()})
			return
		}
		helpers.Error(w, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *TodoController) MarkComplete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid ID", []string{"id must be a valid integer"})
		return
	}

	todo, err := c.service.MarkComplete(id)
	if err != nil {
		if err.Error() == "todo not found" {
			helpers.Error(w, http.StatusNotFound, "Not Found", []string{err.Error()})
			return
		}
		helpers.Error(w, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	helpers.Success(w, http.StatusOK, "Completed", todo)
}

func (c *TodoController) MarkIncomplete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		helpers.Error(w, http.StatusBadRequest, "Invalid ID", []string{"id must be a valid integer"})
		return
	}

	todo, err := c.service.MarkIncomplete(id)
	if err != nil {
		if err.Error() == "todo not found" {
			helpers.Error(w, http.StatusNotFound, "Not Found", []string{err.Error()})
			return
		}
		helpers.Error(w, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	helpers.Success(w, http.StatusOK, "Uncompleted", todo)
}