package v1

import (
	"context"
	_"fmt"
	"net/http"
	"time"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	pb "github.com/NajmiddinAbdulhakim/api-gateway/genproto"
	l "github.com/NajmiddinAbdulhakim/api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"github.com/NajmiddinAbdulhakim/api-gateway/pkg/utils"
	_"github.com/NajmiddinAbdulhakim/api-gateway/api/model"
)



// CreateUser creates user
// @Summary Create user summary
// @Description This api is using create new user
// @Tags User 
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Param user body model.User true "user body"
// @Router /v1/users [post]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		body        pb.User
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().CreateUser(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}


	bodyByte,err := json.Marshal(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed marshal set to redis", l.Error(err))
		return
	}
	
	err = h.redisStorage.Set(body.FirstName,string(bodyByte))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed set to redis", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}


// Login returns list of users
// @Summary Login User summary
// @Description This api is using for login user
// @Tegs Login
// @Accept json
// @Produce json
// @Param email query int true "email"
// @Param password query int true "password"
// @Success 200 {string} model.LoginUser
// @Router /v1/users/login [get]
func (h * handlerV1) LoginUser(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	email := c.Query("email")
	password := c.Query("password")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resp, err := h.serviceManager.UserService().LoginUser(
		ctx,
		&pb.LoginUserReq{
			Email: email,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to email", l.Error(err))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(resp.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to password", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}


// GetUser gets user by id
// @Summary Get User By Id With Posts summary
// @Description This api is using getting by id with posts
// @Tegs User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} model.User
// @Router /v1/users/{id} [get]
func (h *handlerV1) GetUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetUserByIdWithPosts(
		ctx, &pb.UserByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetListUsers returns list of users
// @Summary Get User list summary
// @Description This api is using for getting users list
// @Tegs User
// @Accept json
// @Produce json
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Success 200 {string} model.User
// @Router /v1/users [get]
func (h *handlerV1) GetListUsers(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetListUsers(
		ctx, &pb.GetUserListReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUser updates user by id
// @Summary Update user summary
// @Description This api is using update user by id
// @Tags User 
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Param id path string true "id"
// @Param user body model.UpdateUserReq true "user body"
// @Router /v1/users/{id} [put]
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var (
		body        pb.UpdateUserReq
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().UpdateUser(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteUser deletes user by id
// @Summary Delete user summary
// @Description This api is using delete user by id
// @Tags User 
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Param id path string true "id"
// @Router /v1/users/{id} [delete]
func (h *handlerV1) DeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().DeleteUser(
		ctx, &pb.UserByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
