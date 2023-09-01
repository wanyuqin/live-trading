package entity

import (
	"fmt"
	"log"
	"testing"
)

func TestFoundDetail_ParseTemplate(t *testing.T) {
	detail := FoundDetail{
		FundSummary: FundSummary{
			FdCode:      "0003095",
			FdName:      "中欧医疗健康混合A",
			FoundDate:   "2016-09-29",
			TrupName:    "中国工商银行股份有限公司",
			Totshare:    "252.10亿",
			ManagerName: "葛兰",
			Rating:      "3",
			KeeperName:  "中欧基金管理公司",
		},
		FundPosition: FundPosition{
			StockList: []PositionStock{
				{
					Code:                   "300015",
					Name:                   "爱尔眼科",
					Percent:                8.67,
					CurrentPrice:           18.02,
					ChangePercentage:       -1.26,
					ChangeOfPreQuarter:     "1.42%",
					ChangeOfPreQuarterType: 1,
				},
				{
					Code:                   "603259",
					Name:                   "药明康德",
					Percent:                7.35,
					CurrentPrice:           81.5,
					ChangePercentage:       -0.78,
					ChangeOfPreQuarter:     "1.01%",
					ChangeOfPreQuarterType: 2,
				},
			},
			BondList: []Bond{
				{
					Name:    "22国债14",
					Percent: 0.26,
				},
				{
					Name:    "22国开11",
					Percent: 0.94,
				},
			},
			StockPercent: 92.48,
			CashPercent:  6.23,
			BondPercent:  1.32,
			OtherPercent: 0.17,
		},
		FundManagerList: []FundManager{{
			Name:   "葛兰",
			Resume: "    清华大学本科、美国西北大学生物医学工程专业博士,5年以上证券及基金从业经验。历任国金证券股份有限公司研究所研究员,民生加银基金管理有限公司研究员。2014年10月加入中欧基金管理有限公司,历任任研究员、中欧医疗健康混合型证券投资基金(2016年09月29日)、中欧瑾和灵活配置混合型证券投资基金 (2015年04月13日-2016年04月22日)、中欧瑾源灵活配置混合型证券投资基金 (2015年03月31日-2016年04月22日)、中欧瑾泉灵活配置混合型证券投资基金( 2015年03月16日-2016年04月22日)、中欧明睿新起点混合型证券投资基金( 2015年01月29日-2016年04月22日)、中欧明睿新起点混合型证券投资基金( 2018年07月12日)、中欧医疗创新股票型证券投资基金( 2019年02月28日)、中欧阿尔法混合型证券投资基金 (2020年08月20日)",
		}},

		PositionTypeList: []PositionType{
			{
				TypeDesc: "股票",
				Percent:  92.32,
			},
			{
				TypeDesc: "债券",
				Percent:  1.31,
			},
			{
				TypeDesc: "现金",
				Percent:  6.2,
			},
			{
				TypeDesc: "其他",
				Percent:  0.17,
			},
		},
	}

	content, err := detail.ParseTemplate()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(content)
}
