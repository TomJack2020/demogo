package main

import (
	"demogo/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 正式发布模式
	gin.SetMode(gin.ReleaseMode)
	// 创建一个Gin路由器实例
	router := gin.Default()

	// 定义一个GET路由：响应API状态
	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "API服务运行中",
		})
	})

	// 定义一个Get路由：响应API帮助信息
	router.GET("/pub", func(c *gin.Context) {

		// 从请求参数中获取发布ID
		publishId := c.Query("publishId")
		if publishId == "" {
			c.JSON(http.StatusOK, gin.H{"error": "publishId参数不能为空"})
		}
		// 从数据库中获取发布数据
		am := tools.TestQuery(publishId)

		//创建空的map
		publishData := make(map[string]string)
		if am.PublishId == "" {
			c.JSON(http.StatusOK, gin.H{"error": "未找到该发布信息"})
		}
		//将数据添加到map中
		publishData["publishId"] = am.PublishId
		publishData["title"] = am.Title
		publishData["accountId"] = am.AccountId
		publishData["Sku"] = am.Sku

		//fmt.Println(publishData)
		//// 响应处理结果
		c.JSON(http.StatusOK, gin.H{
			"message": "数据处理成功",
			"result":  publishData,
		})
		//c.String(200, am)

	})

	// 定义一个POST路由：处理数据输入
	router.POST("/process", func(c *gin.Context) {
		var requestData map[string]interface{}

		// 解析JSON请求体
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 假设的数据处理逻辑

		// 响应处理结果
		c.JSON(http.StatusOK, gin.H{
			"message": "数据处理成功",
			"result":  requestData,
		})
	})

	// 启动Gin服务器在8080端口
	err := router.Run(":8082")
	if err != nil {
		return
	}
}
