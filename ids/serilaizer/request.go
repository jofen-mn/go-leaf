package serilaizer

type SegmentAddReq struct {
	BizTag      string `json:"biz_tag"`
	Step        int64  `json:"step"`
	Description string `json:"description"`
}
