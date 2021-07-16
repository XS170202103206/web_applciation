package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"gin/models"
	"gin/response"
	"gin/service"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{service: service.NewOrderService()}
}
//插入单条数据
func (handler *OrderHandler) CreateModel(c *gin.Context) {
	var m models.DemoOrder

	if err := c.BindJSON(&m); err != nil {
		response.ErrorRes(c, http.StatusBadRequest, "Bind json error", err.Error())
		return
	}

	if err := handler.service.CreateModel(&m); err != nil {
		response.ErrorRes(c, http.StatusBadRequest, "Create model error", err.Error())
		return
	}

	response.SuccessRes(c, "Create success", &m)
}
// create order more
func (handler *OrderHandler) CreateModels(c *gin.Context) {
	var ms models.DemoOrders

	if err := c.BindJSON(&ms); err != nil {
		response.ErrorRes(c, http.StatusBadRequest, "Bind json error", err.Error())
		return
	}

	if err := handler.service.CreateModels(&ms); err != nil {
		response.ErrorRes(c, http.StatusInternalServerError, "Create models error", err.Error())
		return
	}

	response.SuccessRes(c, "Create models success", gin.H{"count": len(ms.Models)})
}
//get order list
func (handler *OrderHandler) GetModel(c *gin.Context) {
	ids := c.Params.ByName("id")
	id, _ := strconv.Atoi(ids)
	m, err := handler.service.GetModel(id)
	if err != nil {
		response.ErrorRes(c, http.StatusBadRequest, "Model not exists", err.Error())
		return
	}

	response.SuccessRes(c, "Get model success", m)
}

func (handler *OrderHandler) GetListModels(c *gin.Context) {
	m, err := handler.service.GetListModels()
	if err != nil {
		response.ErrorRes(c, http.StatusBadRequest, "models not exist", err.Error())
		return
	}
	response.SuccessRes(c, "Get models successful", m)
}


//get order single
func (handler *OrderHandler) GetModels(c *gin.Context) {
	var ms []models.DemoOrder

	query := c.DefaultQuery("query", "")
	order := c.DefaultQuery("order", "CREATE")
	by := c.DefaultQuery("by", "ASC")

	if order != "AMOUNT" && order != "CREATE" {
		response.ErrorRes(c, http.StatusBadRequest, "Param order must be AMOUNT o CREATE", errors.New("bad request"))
		return
	} else if by != "ASC" && by != "DESC" {
		response.ErrorRes(c, http.StatusBadRequest, "Param order must be ASC or DESC", errors.New("bad request"))
		return
	}

	ms, err := handler.service.GetModels(query, order, by)
	if err != nil {
		response.ErrorRes(c, http.StatusBadRequest, "Get models error", err.Error())
		return
	}

	response.SuccessRes(c, "Get models success", ms)
}
//update single
func (handler *OrderHandler) UpdateModel(c *gin.Context) {
	var m models.DemoOrder

	ids := c.Params.ByName("id")
	id, _ := strconv.Atoi(ids)
	if err := c.BindJSON(&m); err != nil {
		response.ErrorRes(c, http.StatusBadRequest, "Bind json error", err.Error())
		return
	}

	if err := handler.service.UpdateModel(id, &m); err != nil {
		response.ErrorRes(c, http.StatusBadRequest, "Update failed", err.Error())
		return
	}

	response.SuccessRes(c, "Update success", "Update success")
}
//上传文件
func (handler *OrderHandler) Upload(c *gin.Context) {
	ids := c.Params.ByName("id")
	id, _ := strconv.Atoi(ids)
	file, err := c.FormFile("file")
	if err != nil {
		response.ErrorRes(c, http.StatusBadRequest, "Read file error", err.Error())
		return
	}

	dst, err := handler.service.UploadFile(id, file)
	if err != nil {
		response.ErrorRes(c, http.StatusBadRequest, "File upload failed", err.Error())
	}

	if err = c.SaveUploadedFile(file, *dst); err != nil {
		response.ErrorRes(c, http.StatusBadRequest, "File upload failed", err.Error())
	}

	response.SuccessRes(c, "File upload success", "File upload success")
}

//下载文件
func (handler *OrderHandler) Download(c *gin.Context) {
	ids := c.Params.ByName("id")
	id, _ := strconv.Atoi(ids)
	dst, err := handler.service.DownloadFile(id)
	if err != nil {
		response.ErrorRes(c, http.StatusBadRequest, "Download file error", err.Error())
		return
	}

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", *dst))
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.File(*dst)
}
//导出excel文件
func (handler *OrderHandler) DumpXLSX(c *gin.Context) {
	dst, err := handler.service.DumpXLSX()
	if err != nil {
		response.ErrorRes(c, http.StatusBadRequest, "Dump xlsx error", err.Error())
		return
	}

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", *dst))
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.File(*dst)
}