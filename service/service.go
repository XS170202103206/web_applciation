package service

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"

	"mime/multipart"
	"os"
	"reflect"
	"time"

	"gin/Dao"
	"gin/models"
)



type OrderService struct {
	orderDao *Dao.OrderDao
}

func NewOrderService() *OrderService {
	return &OrderService{orderDao: Dao.NewOrderDao()}
}
//插入单条数据
func (service *OrderService) CreateModel(m *models.DemoOrder) error {
	return service.orderDao.CreateModel(m)
}

func (service *OrderService) CreateModels(ms *models.DemoOrders) error {
	return service.orderDao.CreateModels(ms.Models)
}

func (service *OrderService) GetModel(id string) (*models.DemoOrder, error) {
	return service.orderDao.GetModel(id)
}

func (service *OrderService) GetModels(query, order, by string) ([]models.DemoOrder, error) {
	if order == "CREATE" {
		order = "created_at"
	}
	return service.orderDao.GetModels(query, order, by)
}

func (service *OrderService) UpdateModel(id string, updateModel *models.DemoOrder) error {
	m, err := service.orderDao.GetModel(id)
	if err != nil {
		return err
	}

	m.Amount = updateModel.Amount
	m.Status = updateModel.Status
	m.FileUrl = updateModel.FileUrl
	return service.orderDao.UpdateModel(m)
}

func (service *OrderService) UploadFile(id string, file *multipart.FileHeader) (*string, error) {
	//  检查id是否存在
	m, err := service.orderDao.GetModel(id)
	if err != nil {
		return nil, err
	}

	//  更新文件路径
	dst := fmt.Sprintf("./file/%s_%s", id, file.Filename)
	m.FileUrl = dst
	if err = service.orderDao.UpdateModel(m); err != nil {
		return nil, err
	}

	return &dst, nil
}

func (service *OrderService) DownloadFile(id string) (*string, error) {
	m, err := service.GetModel(id)
	if err != nil {
		return nil, err
	}

	//  校验文件是否存在
	dst := m.FileUrl
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		return nil, err
	}

	return &dst, nil
}

func (service *OrderService) DumpXLSX() (*string, error) {
	//var Models models.DemoOrder
	ms, err := service.orderDao.GetModels("", "created_at", "DESC")
	if err != nil {
		return nil, err
	}

	wb := xlsx.NewFile()
	sh, err := wb.AddSheet("New sheet") //添加表
	if err != nil {
		return nil, err
	}

	row := sh.AddRow() //添加行
	row.SetHeightCM(1)//设置行高

	//  设置表头
	row.AddCell().Value = "ID"
	row.AddCell().Value = "created_at"
	row.AddCell().Value = "updated_at"
	row.AddCell().Value = "deleted_at"
	v := reflect.ValueOf(models.DemoOrder{})
	for i := 1; i < v.NumField(); i++ {
		row.AddCell().Value = v.Type().Field(i).Tag.Get("json")
	}

	for _, m := range ms {
		row := sh.AddRow()
		row.SetHeightCM(1)
		//fmt.Printf("%#v",m)
		row.AddCell().Value = strconv.Itoa(int(m.ID))
		row.AddCell().Value = m.CreatedAt.String()
		row.AddCell().Value = m.UpdatedAt.String()
		row.AddCell().Value = m.DeletedAt.Time.String()

		//  使用反射获取字段值
		v := reflect.ValueOf(m)
		for i := 1; i < v.NumField(); i++ {
			filed := v.Field(i)
			var value string
			switch filed.Interface().(type) {
			case int:
				value = strconv.Itoa(filed.Interface().(int))
			case float64:
				value = strconv.FormatFloat(filed.Interface().(float64), 'f', -1, 64)
			case string:
				value = filed.String()
			case time.Time:
				value = filed.Interface().(time.Time).String()
			}

			row.AddCell().Value = value
		}
	}

	dst := fmt.Sprintf("./dump_%d.xlsx", time.Now().Unix())
	if err = wb.Save(dst); err != nil {
		return nil, err
	}

	return &dst, nil
}