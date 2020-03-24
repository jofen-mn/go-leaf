package dao

import (
	"go-leaf/base/db"
	"go-leaf/ids/entity"
)

/**
 * 查询db获取对应业务标识的id段
 * bizTag 业务标识
 */
func FindSegments(bizTag string) (*entity.Segments, error) {
	segment := &entity.Segments{}
	if err := db.Mysql.First(segment, &entity.Segments{BizTag: bizTag}).Error; err != nil {
		return nil, err
	}
	return segment, nil
}

/**
 * 插入id分段记录
 */
func InsertSegments(segment *entity.Segments) (bool, error) {
	result := db.Mysql.Create(segment)
	if result != nil {
		if result.Error != nil {
			return false, result.Error
		}
	}
	return true, nil
}

func UpdateSegmentsByQuery(query , updateData map[string]interface{}) (bool, error) {
	err := db.Mysql.Model(&entity.Segments{}).Where(&query).Updates(&updateData).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func UpdateSegments(segments *entity.Segments) (bool, error) {
	err := db.Mysql.Model(segments).Updates(segments).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
