package Dao

import (
	"gin/db"
	"gorm.io/gorm"

   "gin/models"
)

type OrderDao struct {
	db *gorm.DB
}

type ModelDao struct{}

func NewOrderDao() *OrderDao {
	return &OrderDao{db: db.NewDb()}
}

//插入单条数据
func (orderDao *OrderDao) CreateModel(m *models.DemoOrder) error {
	if err := orderDao.db.Create(m).Error; err != nil {
		return err
	}
	return nil
}
//插入多条数据
func (orderDao *OrderDao) CreateModels(ms []*models.DemoOrder) error {
	tx := orderDao.db.Begin()
	defer tx.Rollback()

	for _, m := range ms {
		if err := tx.Create(&m).Error; err != nil {
			return err
		}
	}

	return tx.Commit().Error
}
//更新order
func (orderDao *OrderDao) UpdateModel(m *models.DemoOrder) error {
	if err := orderDao.db.Save(m).Error; err != nil {
		return err
	}
	return nil
}
//获取单条order信息
func (orderDao *OrderDao) GetModel(id int) (*models.DemoOrder, error) {
	var m models.DemoOrder
	if err := orderDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}
//获取order列表
func (orderDao *OrderDao) GetModels(query, order, by string) ([]models.DemoOrder, error) {
	var ms []models.DemoOrder
	err := orderDao.db.Where("user_name LIKE ?", "%"+query+"%").
		Order(order + " " + by).
		Find(&ms).Error

	if err != nil {
		return nil, err
	}

	return ms, nil
}
//上传文件
func (orderDao *OrderDao)UploadFile(id string,FileUrl string) (*models.DemoOrder,error){
	//var m models.DemoOrder
	err := orderDao.db.Where("id = ?",id).
		Update("file_url",FileUrl).Error
	if err != nil {
		return nil, err
	}
	return nil,err
}