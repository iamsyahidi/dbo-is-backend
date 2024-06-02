package middleware

import (
	"dbo/dbo-is-backend/models"
	"dbo/dbo-is-backend/pkg/logger"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// Response setting gin.JSON
func Response(c *gin.Context, req interface{}, res models.Response) {
	// LOGGER
	reqByte, _ := json.Marshal(req)
	resByte, _ := json.Marshal(res)
	logger.Infof("[dbo-is-backend:log] [RequestURL] : %s, [RequestMethod] : %s, [RequestBody] : %s, [ResponseData] : %s", c.Request.RequestURI, c.Request.Method, string(reqByte), string(resByte))

	c.JSON(res.Code, res)
}
