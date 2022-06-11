package v1

import (
	"context"
	_"fmt"
	"net/http"
	"time"

	pb "github.com/NajmiddinAbdulhakim/api-gateway/genproto"
	l "github.com/NajmiddinAbdulhakim/api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"github.com/NajmiddinAbdulhakim/api-gateway/pkg/utils"
	// "github.com/NajmiddinAbdulhakim/api-gateway/api/hendlers/model"
)

type UserReq struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	FirstName            string     `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name"`
	LastName             string     `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name"`
	UserName             string     `protobuf:"bytes,4,opt,name=user_name,json=userName,proto3" json:"user_name"`
	Email                string     `protobuf:"bytes,5,opt,name=email,proto3" json:"email"`
	PhoneNumber          []string   `protobuf:"bytes,6,rep,name=phone_number,json=phoneNumber,proto3" json:"phone_number"`
	Addresses            []*Address `protobuf:"bytes,7,rep,name=addresses,proto3" json:"addresses"`
	// Posts                []*Post    `protobuf:"bytes,8,rep,name=posts,proto3" json:"posts"`
	Bio                  string     `protobuf:"bytes,9,opt,name=bio,proto3" json:"bio"`
	Status               string     `protobuf:"bytes,10,opt,name=status,proto3" json:"status"`
}

type Address struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Country              string   `protobuf:"bytes,2,opt,name=country,proto3" json:"country"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city"`
	District             string   `protobuf:"bytes,4,opt,name=district,proto3" json:"district"`
	PostalCode           string   `protobuf:"bytes,5,opt,name=postal_code,json=postalCode,proto3" json:"postal_code"`
}



// CreateUser creates user
// @Summary Create user summary
// @Description This api is using create new user
// @Tags User 
// @Accept json
// @Produce json
// @Success 200 {string} Success
// @Param user body UserReq true "user body"
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

	c.JSON(http.StatusCreated, response)
}

// GetUser gets user by id
// route /v1/users/{id} [get]
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

// ListUsers returns list of users
// route /v1/users/ [get]
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
// route /v1/users/{id} [put]
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

// // DeleteUser deletes user by id
// // route /v1/users/{id} [delete]
// func (h *handlerV1) DeleteUser(c *gin.Context) {
// 	var jspbMarshal protojson.MarshalOptions
// 	jspbMarshal.UseProtoNames = true

// 	guid := c.Param("id")
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()

// 	response, err := h.serviceManager.UserService().Delete(
// 		ctx, &pb.ByIdReq{
// 			Id: guid,
// 		})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to delete user", l.Error(err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }
