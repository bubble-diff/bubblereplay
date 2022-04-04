package models

type Task struct {
	ID            int64          `json:"id" bson:"id"`
	TrafficConfig *TrafficConfig `json:"traffic_config" bson:"traffic_config"`
	FilterConfig  *FilterConfig  `json:"filter_config" bson:"filter_config"`
	AdvanceConfig *AdvanceConfig `json:"advance_config" bson:"advance_config"`
	TotalRecord   int64          `json:"total_record" bson:"total_record"`
	SuccessRecord int64          `json:"success_record" bson:"success_record"`
}
