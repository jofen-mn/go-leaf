package controller

import (
	"github.com/gin-gonic/gin"
	"go-leaf/ids/serilaizer"
	"go-leaf/ids/service"
)

func GetIds(c *gin.Context) {
	bizTag := c.Param("biz_tag")

	result, err := service.GetIds(bizTag)
	if err != nil {
		serilaizer.Response(c, err, nil)
		return
	}
	serilaizer.Response(c, nil, result)
}

func AddSegment(c *gin.Context) {
	segmentAddReq := &serilaizer.SegmentAddReq{}
	err := c.BindJSON(segmentAddReq)
	if err != nil {
		serilaizer.Response(c, err, nil)
		return
	}

	if flag, err := service.AddSegment(segmentAddReq); err != nil {
		serilaizer.Response(c, err, nil)
		return
	} else if !flag {
		// TODO fill data
		serilaizer.Response(c, nil, nil)
		return
	}

	serilaizer.Response(c, nil, nil)
}
