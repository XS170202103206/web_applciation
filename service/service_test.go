package service

import (
	"strconv"
	"testing"

	"gin/models"
)

func TestModelService_CRU(t *testing.T) {
	testCases := []struct {
		name        string
		createModel *models.DemoOrder
		updateModel *models.DemoOrder
		readModel   *models.DemoOrder
	}{
		{
			name:        "success 1",
			createModel: &models.DemoOrder{OrderNo: "0", UserName: "user1", Amount: 0, Status: "0", FileUrl: ""},
			updateModel: &models.DemoOrder{OrderNo: "1", UserName: "user1", Amount: 0.99, Status: "1", FileUrl: "./1.png"},
			readModel:   &models.DemoOrder{},
		},
		{
			name:        "success 2",
			createModel: &models.DemoOrder{OrderNo: "0", UserName: "user2", Amount: 0, Status: "0", FileUrl: ""},
			updateModel: &models.DemoOrder{OrderNo: "100", UserName: "user2", Amount: 100, Status: "1", FileUrl: "./2.png"},
			readModel:   &models.DemoOrder{},
		},
		{
			name:        "success 3",
			createModel: &models.DemoOrder{OrderNo: "0", UserName: "user3", Amount: 0, Status: "1", FileUrl: ""},
			updateModel: &models.DemoOrder{OrderNo: "0", UserName: "user3", Amount: 100000.1, Status: "0", FileUrl: "./3.png"},
			readModel:   &models.DemoOrder{},
		},
	}

	var service OrderService
	//  test create update read
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//  Create Model
			err := service.CreateModel(testCase.createModel)
			if err != nil {
				t.Errorf("Create model error: %v", err)
			}

			//  Update Model
			testCase.updateModel.Model = testCase.createModel.Model
			err = service.UpdateModel(testCase.updateModel)
			if err != nil {
				t.Errorf("Update model error: %v", err)
			}

			//  Get Model
			err = service.GetModel(strconv.Itoa(int(testCase.createModel.ID)), testCase.readModel)
			if err != nil {
				t.Errorf("Get model error: %v", err)
			}
			//  Check model if they're equal
			if *testCase.updateModel != *testCase.readModel {
				t.Errorf("Read model is not equal the update model")
			}
		})
	}
}

func TestModelService_GetModels(t *testing.T) {
	testCases := []struct {
		name  string
		query string
		order string
		by    string
	}{
		{name: "success 1", query: "", order: "CREATE", by: "ASC"},
		{name: "success 2", query: "123", order: "CREATE", by: "DESC"},
		{name: "success 3", query: "123", order: "AMOUNT", by: "ASC"},
	}

	var modelss []models.DemoOrder
	var service OrderService
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if err := service.GetModels(&modelss, testCase.query, testCase.order, testCase.by); err != nil {
				t.Errorf("Get all model error: %v\n", err)
			}
		})
	}
}

func TestModelService_CreateModels(t *testing.T) {
	testCases := []struct {
		name         string
		createModels *models.DemoOrders
	}{
		{
			name: "success 1",
			createModels: &models.DemoOrders{
				Models: []models.DemoOrder{
					{OrderNo: "0", UserName: "user1", Amount: 0, Status: "0", FileUrl: ""},
					{OrderNo: "0", UserName: "user2", Amount: 0, Status: "0", FileUrl: ""},
					{OrderNo: "0", UserName: "user3", Amount: 0, Status: "0", FileUrl: ""},
					{OrderNo: "0", UserName: "user4", Amount: 0, Status: "0", FileUrl: ""},
					{OrderNo: "0", UserName: "user5", Amount: 0, Status: "0", FileUrl: ""},
					{OrderNo: "0", UserName: "user6", Amount: 0, Status: "0", FileUrl: ""},
					{OrderNo: "0", UserName: "user7", Amount: 0, Status: "0", FileUrl: ""},
				},
			},
		},
		{
			name: "success 2",
			createModels: &models.DemoOrders{
				Models: []models.DemoOrder{
				},
			},
		},
	}

	var service OrderService
	//  test create update read
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//  Create Models
			err := service.CreateModels(&testCase.createModels.Models)
			if err != nil {
				t.Errorf("Create models error: %v", err)
			}
		})
	}
}