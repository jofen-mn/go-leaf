package dao

import (
	"go-leaf/base/db"
	"go-leaf/conf"
	"testing"
)

func TestUpdateSegments(t *testing.T) {
	segment, err := FindSegments("test")
	if err != nil {
		panic(err)
	}

	segment.MaxId = segment.MaxId + segment.Step

	_, err = UpdateSegmentsTest(segment)
	if err != nil {
		panic(err)
	}
}

func init() {
	conf.ConfInit()
	db.InitMysql()
}
