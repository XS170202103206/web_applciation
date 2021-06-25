package router

import "gin/handler"

func (r *Router) V1Group() {
	h := handler.NewOrderHandler()

	r.V1 = r.Root.Group("/api/v1")
	r.V1.GET("/models", h.GetModels)                   //  get list
	r.V1.GET("/models/:id", h.GetModel)                //  get single
	r.V1.POST("/models", h.CreateModel)                //  create single
	r.V1.POST("/models/count", h.CreateModels)         //  create more
	r.V1.PUT("/models/:id", h.UpdateModel)             //  update single
	r.V1.POST("/models/file/upload/:id", h.Upload)     //  upload file
	r.V1.GET("/models/file/download/:id", h.Download)  //  download file
	r.V1.GET("/models/file/download/xlsx", h.DumpXLSX) //  download xlsx

	//r.V1.POST("/models/upload",h.Upload2) //test upload function
}