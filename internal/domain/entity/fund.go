package entity

type Fund struct {
	Code             string `json:"code"`
	Name             string `json:"name"`
	FNAV             string `json:"fnav"`               // 单位净值
	TRF              string `json:"trf"`                // 累计净值
	DailyInc         string `json:"daily_inc"`          // 日增长值
	DailyIncRate     string `json:"daily_inc_rate"`     // 日增长率
	InceptionIncRate string `json:"inception_inc_rate"` // 成立以来
	Type             string `json:"type"`
	DataTime         string `json:"data_time"` // 数据日期
}

type FundList []Fund
