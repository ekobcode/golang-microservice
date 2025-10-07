package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"golang-microservice/internal/entity"
	"golang-microservice/internal/usecase"
	"golang-microservice/internal/utils"
)

type UserHandler struct {
	usecase usecase.UserUsecase
	logger  *zap.SugaredLogger
}

// NewUserHandler creates a new handler with injected usecase and logger
func NewUserHandler(usecase usecase.UserUsecase, logger *zap.SugaredLogger) *UserHandler {
	return &UserHandler{usecase: usecase, logger: logger}
}

// Create handles POST /users
func (h *UserHandler) Create(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondError(c, h.logger, err, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.usecase.Create(&user); err != nil {
		utils.RespondError(c, h.logger, err, "Failed to create user", http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   user,
	})
}

// GetAll handles GET /users
func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.usecase.GetAll()
	if err != nil {
		utils.RespondError(c, h.logger, err, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
	})
}

// GetByID handles GET /users/:id
func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.RespondError(c, h.logger, err, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.usecase.GetByID(id)
	if err != nil {
		utils.RespondError(c, h.logger, err, "User not found", http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

// Update handles PUT /users/:id
func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.RespondError(c, h.logger, err, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondError(c, h.logger, err, "Invalid request payload", http.StatusBadRequest)
		return
	}
	user.ID = id

	if err := h.usecase.Update(&user); err != nil {
		utils.RespondError(c, h.logger, err, "Failed to update user", http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

// Delete handles DELETE /users/:id
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.RespondError(c, h.logger, err, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// 1️⃣ Check if user exists before deleting
	user, err := h.usecase.GetByID(id)
	if err != nil || user == nil {
		utils.RespondError(c, h.logger, err, "User not found", http.StatusNotFound)
		return
	}

	if err := h.usecase.Delete(id); err != nil {
		utils.RespondError(c, h.logger, err, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User deleted",
	})
}
