package v1

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/mail.v2"

	"github.com/NajmiddinAbdulhakim/api-gateway/api/auth"
	"github.com/NajmiddinAbdulhakim/api-gateway/api/model"
	pbu "github.com/NajmiddinAbdulhakim/api-gateway/genproto"
	l "github.com/NajmiddinAbdulhakim/api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *handlerV1) Register(c *gin.Context) {
	var (
		body        model.RegisterUser
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
	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	
	exists, err := h.serviceManager.UserService().CheckUnique(ctx, &pbu.CheckUniqueReq{
		Field: "email", 
		Value:  body.Email, 
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed to check email unique", l.Error(err))
			return
		}
	if exists.IsExists {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This email already in use, please use another email",
		})
		h.log.Error("failed to check email unique", l.Error(err))
		return
	}

	existss, err := h.serviceManager.UserService().CheckUnique(ctx, &pbu.CheckUniqueReq{
		Field: "username", 
		Value:  body.UserName, 
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed to check username unique", l.Error(err))
			return
		}
	if existss.IsExists {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This username already in use, please use another username",
		})
		h.log.Error("failed to check username unique", l.Error(err))
		return
	}

	rand.Seed(time.Now().UnixNano())
    min := 100000
    max := 999999
    code := int64(rand.Intn(max - min + 1) + min)
	SendEmail(body.Email,strconv.FormatInt(int64(code),10))
	body.Code = code



	byteUser, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while marshaling user data to redis", l.Error(err))
		return
	}

	err = h.redisStorage.SetWithTTL(body.Email, string(byteUser), int64(time.Second*300))
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while seting user data to redis", l.Error(err))
		return
	}
}

func SendEmail(email,message string) {
	m := gomail.NewMessage()
  
	m.SetHeader("From", "Suqrothakim@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("subject", "Verification code")
	m.SetBody("text/plain", message)
	d := gomail.NewDialer("smtp.gmail.com", 587, "Suqrothakim@gmail.com", "vyzoqqrpwapfpgbf")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
  }

func (h * handlerV1) Verify(c *gin.Context){
	var userData model.RegisterUser
	email := c.Query("Email")
	code := c.Query("Code")

	cod,err := strconv.ParseInt(code, 10, 64)

	data, err := redis.String(h.redisStorage.Get(email))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting data from redis", l.Error(err))
		return
	}
	err = json.Unmarshal([]byte(data),&userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while unmarshiling user data", l.Error(err))
		return
	}

	if cod != userData.Code {
		c.JSON(http.StatusConflict, gin.H{
			"error": `Your code is invalid`,
		})
		h.log.Error("Code is invalid", l.Error(err))
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
    if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while hashing user data", l.Error(err))
		return
	}
	id, err := uuid.NewV4()
	if  err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while generating uuid", l.Error(err))
		return
	}
	userData.Id = id.String()
	h.jwtHendler = auth.JWTHendler{
		Sub: userData.Id,
		Iss: "client",
		Role: "authorized",
		Log: h.log,
	}
	access, refresh, err := h.jwtHendler.GenerateAuthJWT()
	if  err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while generating jwt tokens", l.Error(err))
		return
	}
	userData.RefreshToken = refresh

	//create logik
	
	c.JSON(http.StatusOK, &model.RegisterUserRes{
		Id: userData.Id,
		RefreshToken: refresh,
		AccessToken: access,
	})
} 