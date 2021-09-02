package apis

import (
	"github.com/gin-gonic/gin"
	"leaning.go.firebase/src/repository"
	"net/http"
)

type userApi struct {
	userRepo repository.UserRepository
}

func RegisterUserApi(ginRouterGroup *gin.RouterGroup, repo repository.UserRepository) {
	api := &userApi{userRepo: repo}

	ginRouterGroup.GET("/user/:user_id", api.get)
	ginRouterGroup.POST("/user", api.create)
	ginRouterGroup.PUT("/user/:user_id", api.update)
}

func (ref *userApi) get(c *gin.Context) {
	request := GetUserRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ref.userRepo.Get(request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, userResponseFromDomain(*user))
}

func (ref *userApi) create(c *gin.Context) {
	request := CreateUpdateUserRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := ref.userRepo.Create(request.ToDomain())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userResponseFromDomain(*created))
}

func (ref *userApi) update(c *gin.Context) {
	request := CreateUpdateUserRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := ref.userRepo.Update(request.ToDomain())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userResponseFromDomain(*updated))
}

type GetUserRequest struct {
	ID string `uri:"user_id" binding:"required"`
}

type CreateUpdateUserRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
}

func (req CreateUpdateUserRequest) ToDomain() repository.User {
	return repository.User{
		ID: req.ID, Name: req.Name, DateOfBirth: req.DateOfBirth,
	}
}

type UserResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	Version     int    `json:"version"`
}

func userResponseFromDomain(user repository.User) UserResponse {
	return UserResponse{
		ID: user.ID, Name: user.Name, DateOfBirth: user.DateOfBirth, Version: user.Version,
	}
}
