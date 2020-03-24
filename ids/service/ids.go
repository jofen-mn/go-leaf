package service

import (
	"github.com/jinzhu/gorm"
	"go-leaf/ids/dao"
	"go-leaf/ids/entity"
	"go-leaf/ids/serilaizer"
	"log"
	"time"
)

func GetIds(bizTag string) (interface{}, error) {
	result, err := dao.FindSegments(bizTag)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("query record not found: %+v\n", err)
		}
		log.Printf("query db error: %+v\n", err)
		return nil, err
	}
	result.MaxId = result.MaxId + result.Step
	// 更新max_id
	_, err = dao.UpdateSegments(result)
	if err != nil {
		log.Printf("update db error: %+v\n", err)
		return nil, err
	}

	return &struct {
		MaxId int64 `json:"max_id"`
		OffSet int64 `json:"off_set"`
	}{result.MaxId, result.MaxId - result.Step}, nil
}

func AddSegment(segmentAddReq *serilaizer.SegmentAddReq) (bool, error) {
	segment := &entity.Segments{
		BizTag:      segmentAddReq.BizTag,
		MaxId:       0,
		Step:        segmentAddReq.Step,
		Description: segmentAddReq.Description,
		UpdateTime:  time.Now(),
	}

	if result, err := dao.FindSegments(segmentAddReq.BizTag); err != nil {
		return false, err
	} else if result != nil {
		return false, nil
	}

	_, err := dao.InsertSegments(segment)
	if err != nil {
		log.Printf("insert segment error: %+v", err)
		return false, err
	}

	return true, nil
}
