package entity

import (
	"bytes"
	"html/template"
)

const detailContent = `
## 概况
|  基金全称        	{{.FdName}}                       
|  基金代码        	{{.FdCode}}                     
|  成立日期        	{{.FoundDate}}                  
|  基金规模        	{{.Totshare}}                   
|  基金经理        	{{.ManagerName}}                
|  基金公司        	{{.KeeperName}}                 
|  托管银行        	{{.TrupName}}                   
## 持仓
o 股票 {{.FundPosition.StockPercent}}%    o 债券 {{.FundPosition.BondPercent}}%
o 现金 {{.FundPosition.CashPercent}}%     o 其他 {{.FundPosition.OtherPercent}}%
### 股票
重仓股票         日涨幅     持仓占比     较上期变动
{{range .FundPosition.StockList}}
{{.Name}}        {{.ChangePercentage}}%     {{.Percent}}%       {{if gt .ChangeOfPreQuarterType  1}}{{.ChangeOfPreQuarter}}⬆{{else}}{{.ChangeOfPreQuarter}}⬇{{end}}
{{end}}
### 债券
重仓债券                                  持仓占比
{{range .FundPosition.BondList}}
{{.Name}}                                  {{.Percent}}%      
{{end}}
`

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

type PositionStock struct {
	Code                   string  `json:"code"`
	Name                   string  `json:"name"`
	Percent                float64 `json:"percent"`                    // 持仓占比
	CurrentPrice           float64 `json:"current_price"`              // 当前价
	ChangePercentage       float64 `json:"change_percentage"`          // 日涨幅
	ChangeOfPreQuarter     string  `json:"change_of_pre_quarter"`      // 较上期百分比
	ChangeOfPreQuarterType int     `json:"change_of_pre_quarter_type"` // 较上期变化类型
	IndustryLabel          string  `json:"industry_label"`             // 标签
}

type PositionStockList []PositionStock

type FundManager struct {
	Name            string  `json:"name"`
	WorkYear        string  `json:"work_year"`
	PostDate        int64   `json:"post_date"`
	PostStatus      int     `json:"post_status"`
	CpTerm          string  `json:"cp_term"`
	CpRate          float64 `json:"cp_rate"`
	PostName        int     `json:"post_name"`
	PerformanceYear float64 `json:"performance_year"`
	FundTotalNav    float64 `json:"fund_total_nav"`
	Resume          string  `json:"resume"`  // 简介
	College         string  `json:"college"` // 大学
}

type FundManagerList []FundManager

type FoundDetail struct {
	FundSummary `json:"fund_summary"`

	FundCompany      string         `json:"fund_company"`
	FundPosition     FundPosition   `json:"fund_position"`      // 持仓信息
	FundManagerList  []FundManager  `json:"fund_manager_list"`  // 管理人
	PositionTypeList []PositionType `json:"position_type_list"` // 持仓类型

}

type FundSummary struct {
	FdCode           string `json:"fd_code"`
	FdType           string `json:"fd_type"`
	FdName           string `json:"fd_name"`
	FdFullName       string `json:"fd_full_name"`
	FoundDate        string `json:"found_date"`
	FdStatus         string `json:"fd_status"`
	DeclareStatus    string `json:"declare_status"`
	SubscribeStatus  string `json:"subscribe_status"`
	WithdrawStatus   string `json:"withdraw_status"`
	AutoInvestStatus string `json:"auto_invest_status"`
	Totshare         string `json:"totshare"`
	KeeperName       string `json:"keeper_name"`
	ManagerName      string `json:"manager_name"`
	TrupName         string `json:"trup_name"`
	Rating           string `json:"rating"`
	SaleStatus       string `json:"sale_status"`
	RiskLevel        string `json:"risk_level"`
	IpoStartDate     int64  `json:"ipo_start_date"`
	IpoEndDate       int64  `json:"ipo_end_date"`
}

type FundPosition struct {
	Source       string          `json:"source"`
	SourceMark   string          `json:"source_mark"`
	Enddate      int64           `json:"enddate"`
	AssetTot     float64         `json:"asset_tot"`
	AssetVal     float64         `json:"asset_val"`
	StockPercent float64         `json:"stock_percent"`
	CashPercent  float64         `json:"cash_percent"`
	BondPercent  float64         `json:"bond_percent"`
	OtherPercent float64         `json:"other_percent"`
	StockList    []PositionStock `json:"stock_list"`
	BondList     []Bond          `json:"bond_list"`
}

type Bond struct {
	Name          string  `json:"name"`
	Code          string  `json:"code"`
	Percent       float64 `json:"percent"`
	XqSymbol      string  `json:"xq_symbol"`
	Amarket       bool    `json:"amarket"`
	PercentDouble float64 `json:"percent_double"`
}

type PositionTypeList []PositionType

type PositionType struct {
	TypeDesc string  `json:"type_desc"`
	Percent  float64 `json:"percent"`
}

func (fd *FoundDetail) ParseTemplate() (string, error) {
	tmpl, err := template.New("fundDetail").Parse(detailContent)
	if err != nil {
		return "", err
	}
	buffer := bytes.Buffer{}
	err = tmpl.Execute(&buffer, fd)
	if err != nil {
		return "", err

	}

	return buffer.String(), nil

}
