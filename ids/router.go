package ids

import (
	"github.com/gin-gonic/gin"
	"go-leaf/ids/controller"
)

func NewIds(baseGroup *gin.RouterGroup) {
	ids := baseGroup.Group("/ids")
	ids.GET("/:biz_tag", controller.GetIds)
	ids.POST("", controller.AddSegment)
}
