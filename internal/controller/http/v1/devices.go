package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nextclan/validator-register/mobile-validator-register-service/internal/entity"
	usecase "nextclan/validator-register/mobile-validator-register-service/internal/usecase"
	"nextclan/validator-register/mobile-validator-register-service/pkg/logger"
)

type deviceRoutes struct {
	cmv usecase.CreateMobileValidatorDevice
	umv usecase.UpdateMobileValidatorDeviceStatus
	fmv usecase.FindByIdMobileValidatorDeviceStatus
	pa  usecase.PusherAuthentication
	l   logger.Interface
}

func newDeviceRoutes(handler *gin.RouterGroup, cmv usecase.CreateMobileValidatorDevice, fmv usecase.FindByIdMobileValidatorDeviceStatus, umv usecase.UpdateMobileValidatorDeviceStatus, pusherAuth usecase.PusherAuthentication, l logger.Interface) {
	r := &deviceRoutes{cmv, umv, fmv, pusherAuth, l}
	handler.POST("/devices", r.doRegisterDevice)
	handler.POST("/devices/:deviceId/status", r.doUpdateDeviceStatus)
	handler.GET("/devices/auth", r.doAuth)
}

type registerDeviceCommand struct {
	UserId    string `json:"userId"`
	DeviceId  string `json:"deviceId"`
	PublicKey string `json:"publicKey"`
}

type FindByIdStatusCommand struct{}
type updateDeviceStatusCommand struct{}

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Param       request body doTranslateRequest true "Set up translation"
// @Success     200 {object} entity.Translation
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /translation/do-translate [post]
func (r *deviceRoutes) doRegisterDevice(c *gin.Context) {
	var request registerDeviceCommand
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - registerDeviceCommand")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	device := &entity.Device{
		UserId:    request.UserId,
		DeviceId:  request.DeviceId,
		PublicKey: request.PublicKey,
	}
	err := r.cmv.Execute(device)
	if err != nil {
		r.l.Error(err, "http - v1 - doRegisterDevice")
		errorResponse(c, http.StatusInternalServerError, "Device service problems")
		return
	}

	c.Writer.WriteHeader(http.StatusAccepted)
}

// //////////////// findUserID(userId, publickey) //////////////////
func (r *deviceRoutes) doFindByIdStatus(c *gin.Context) {
	var request FindByIdStatusCommand
	deviceId := c.Param("deviceId")
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - FindByIdStatusCommand")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	device := &entity.Device{
		DeviceId: deviceId,
	}
	err := r.fmv.Execute(device)
	if err != nil {
		r.l.Error(err, "http - v1 - doFindByIdStatus")
		errorResponse(c, http.StatusInternalServerError, "Device service problems")

		return
	}
	c.Writer.WriteHeader((http.StatusAccepted))
}

func (r *deviceRoutes) doUpdateDeviceStatus(c *gin.Context) {
	var request updateDeviceStatusCommand
	deviceId := c.Param("deviceId")
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - updateDeviceStatusCommand")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	device := &entity.Device{
		DeviceId: deviceId,
	}
	err := r.umv.Execute(device)
	if err != nil {
		r.l.Error(err, "http - v1 - doUpdateDeviceStatus")
		errorResponse(c, http.StatusInternalServerError, "Device service problems")

		return
	}
	c.Writer.WriteHeader(http.StatusAccepted)
}

type authCommand struct {
	UserId string `json:"userId"`
}

func (r *deviceRoutes) doAuth(c *gin.Context) {
	userId := c.Query("user_id")

	token, err := r.pa.Execute(userId)
	if err != nil {
		r.l.Error(err, "http - v1 - doAuth")
		errorResponse(c, http.StatusInternalServerError, "Pusher Beam service problems")
		return
	}
	c.JSON(http.StatusOK, token)
}
