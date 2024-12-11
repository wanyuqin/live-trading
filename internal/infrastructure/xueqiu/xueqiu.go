package xueqiu

import (
	"encoding/json"
	"github.com/jinzhu/copier"
	"live-trading/internal/domain/entity"
)

var (
	positionStockUrl = "https://danjuanfunds.com/djapi/fundx/base/fund/record/asset/percent"
	fundManagerUrl   = "https://danjuanfunds.com/djapi/fundx/base/fund/record/manager/list"
	fundDetailUrl    = "https://danjuanfunds.com/djapi/fund/detail"
	fundSummaryUrl   = "https://danjuanfunds.com/djapi/fund"
	listMarketUrl    = "https://stock.xueqiu.com/v5/stock/batch/quote.json?symbol=SH000001,SZ399001,SZ399006,SH000688"
)

type PositionStockResponse struct {
	Data struct {
		Source       string  `json:"source"`
		SourceMark   string  `json:"source_mark"`
		StockPercent float64 `json:"stock_percent"`
		CashPercent  float64 `json:"cash_percent"`
		BondPercent  float64 `json:"bond_percent"`
		OtherPercent float64 `json:"other_percent"`
		StockList    []struct {
			Name                   string  `json:"name"`
			Code                   string  `json:"code"`
			Percent                float64 `json:"percent"`
			CurrentPrice           float64 `json:"current_price"`
			ChangePercentage       float64 `json:"change_percentage"`
			XqSymbol               string  `json:"xq_symbol"`
			XqUrl                  string  `json:"xq_url"`
			ChangeOfPreQuarter     string  `json:"change_of_pre_quarter"`
			ChangeOfPreQuarterType int     `json:"change_of_pre_quarter_type"`
			IndustryLabel          string  `json:"industry_label"`
			Amarket                bool    `json:"amarket"`
			PercentDouble          float64 `json:"percent_double"`
		} `json:"stock_list"`
		BondList []struct {
			Name          string  `json:"name"`
			Code          string  `json:"code"`
			Percent       float64 `json:"percent"`
			XqSymbol      string  `json:"xq_symbol"`
			Amarket       bool    `json:"amarket"`
			PercentDouble float64 `json:"percent_double"`
		} `json:"bond_list"`
		ChartList []struct {
			TypeDesc string  `json:"type_desc"`
			Type     string  `json:"type"`
			Percent  float64 `json:"percent"`
			Color    string  `json:"color"`
		} `json:"chart_list"`
		ImportantListShow int `json:"important_list_show"`
		IndustryList      []struct {
			IndustryCode string  `json:"industry_code"`
			IndustryName string  `json:"industry_name"`
			Percent      float64 `json:"percent"`
			Color        string  `json:"color"`
		} `json:"industry_list"`
		IndustryTip []struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		} `json:"industry_tip"`
	} `json:"data"`
	ResultCode int `json:"result_code"`
}

func parseFundPosition(body []byte) (entity.FundPosition, error) {
	response := PositionStockResponse{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return entity.FundPosition{}, err

	}

	fundPosition := entity.FundPosition{}
	copier.Copy(&fundPosition, response.Data)

	return fundPosition, nil

}

type FundManagerResponse struct {
	Data struct {
		Items []struct {
			IndiId          string  `json:"indi_id"`
			Name            string  `json:"name"`
			WorkYear        string  `json:"work_year"`
			PostDate        int64   `json:"post_date"`
			PostStatus      int     `json:"post_status"`
			CpTerm          string  `json:"cp_term"`
			CpRate          float64 `json:"cp_rate"`
			PostName        int     `json:"post_name"`
			PerformanceYear float64 `json:"performance_year"`
			FundTotalNav    float64 `json:"fund_total_nav"`
		} `json:"items"`
		CurrentPage int `json:"current_page"`
		Size        int `json:"size"`
		TotalItems  int `json:"total_items"`
		TotalPages  int `json:"total_pages"`
	} `json:"data"`
	ResultCode int `json:"result_code"`
}

func parseFundManager(body []byte) (entity.FundManagerList, error) {
	response := FundManagerResponse{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	fundManagers := make(entity.FundManagerList, len(response.Data.Items))

	for i := range response.Data.Items {
		manager := response.Data.Items[i]
		fundManager := entity.FundManager{}
		copier.Copy(&fundManager, manager)
		fundManagers = append(fundManagers, fundManager)
	}

	return fundManagers, nil
}

type FundDetailResponse struct {
	Data struct {
		FundCompany  string `json:"fund_company"`
		FundPosition struct {
			Source       string  `json:"source"`
			SourceMark   string  `json:"source_mark"`
			Enddate      int64   `json:"enddate"`
			AssetTot     float64 `json:"asset_tot"`
			AssetVal     float64 `json:"asset_val"`
			StockPercent float64 `json:"stock_percent"`
			CashPercent  float64 `json:"cash_percent"`
			BondPercent  float64 `json:"bond_percent"`
			OtherPercent float64 `json:"other_percent"`
			StockList    []struct {
				Name                   string  `json:"name"`
				Code                   string  `json:"code"`
				Percent                float64 `json:"percent"`
				CurrentPrice           float64 `json:"current_price"`
				ChangePercentage       float64 `json:"change_percentage"`
				XqSymbol               string  `json:"xq_symbol"`
				XqUrl                  string  `json:"xq_url"`
				ChangeOfPreQuarter     string  `json:"change_of_pre_quarter"`
				ChangeOfPreQuarterType int     `json:"change_of_pre_quarter_type"`
				IndustryLabel          string  `json:"industry_label"`
				Amarket                bool    `json:"amarket"`
				PercentDouble          float64 `json:"percent_double"`
			} `json:"stock_list"`
			BondList []struct {
				Name          string  `json:"name"`
				Code          string  `json:"code"`
				Percent       float64 `json:"percent"`
				XqSymbol      string  `json:"xq_symbol"`
				Amarket       bool    `json:"amarket"`
				PercentDouble float64 `json:"percent_double"`
			} `json:"bond_list"`
			ChartList []struct {
				TypeDesc string  `json:"type_desc"`
				Type     string  `json:"type"`
				Percent  float64 `json:"percent"`
				Color    string  `json:"color"`
			} `json:"chart_list"`
			ImportantListShow int `json:"important_list_show"`
			IndustryList      []struct {
				IndustryCode string  `json:"industry_code"`
				IndustryName string  `json:"industry_name"`
				Percent      float64 `json:"percent"`
				Color        string  `json:"color"`
			} `json:"industry_list"`
			IndustryTip []struct {
				Title   string `json:"title"`
				Content string `json:"content"`
			} `json:"industry_tip"`
			EndDateStr string `json:"end_date_str"`
		} `json:"fund_position"`
		FundRates struct {
			FdCode            string `json:"fd_code"`
			SubscribeRate     string `json:"subscribe_rate"`
			DeclareRate       string `json:"declare_rate"`
			WithdrawRate      string `json:"withdraw_rate"`
			Discount          string `json:"discount"`
			SubscribeDiscount string `json:"subscribe_discount"`
			DeclareDiscount   string `json:"declare_discount"`
			DeclareRateTable  []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"declare_rate_table"`
			WithdrawRateTable []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"withdraw_rate_table"`
			OtherRateTable []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"other_rate_table"`
		} `json:"fund_rates"`
		ManagerList []struct {
			IndiId          string `json:"indi_id"`
			Name            string `json:"name"`
			Resume          string `json:"resume"`
			College         string `json:"college"`
			WorkYear        string `json:"work_year"`
			AchievementList []struct {
				FundCode  string  `json:"fund_code"`
				Fundsname string  `json:"fundsname"`
				PostDate  string  `json:"post_date"`
				CpRate    float64 `json:"cp_rate"`
				ResiDate  string  `json:"resi_date,omitempty"`
			} `json:"achievement_list"`
		} `json:"manager_list"`
		FundDateConf struct {
			FdCode          string `json:"fd_code"`
			BuyConfirmDate  int    `json:"buy_confirm_date"`
			BuyQueryDate    int    `json:"buy_query_date"`
			SaleConfirmDate int    `json:"sale_confirm_date"`
			SaleQueryDate   int    `json:"sale_query_date"`
			AllBuyDays      int    `json:"all_buy_days"`
			AllSaleDays     int    `json:"all_sale_days"`
		} `json:"fund_date_conf"`
		PensionFund bool `json:"pension_fund"`
	} `json:"data"`
	ResultCode int `json:"result_code"`
}

func parseFundDetail(body []byte) (entity.FoundDetail, error) {
	fundDetail := entity.FoundDetail{}
	response := FundDetailResponse{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return fundDetail, err
	}

	fundDetail.FundCompany = response.Data.FundCompany
	copier.Copy(&fundDetail.FundManagerList, response.Data.ManagerList)
	return fundDetail, nil
}

type FundSummaryResponse struct {
	Data struct {
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
		FundDerived      struct {
			EndDate               string `json:"end_date"`
			UnitNav               string `json:"unit_nav"`
			NavGrtd               string `json:"nav_grtd"`
			NavGrl1M              string `json:"nav_grl1m"`
			NavGrl3M              string `json:"nav_grl3m"`
			NavGrl6M              string `json:"nav_grl6m"`
			NavGrlty              string `json:"nav_grlty"`
			NavGrl1Y              string `json:"nav_grl1y"`
			NavGrl3Y              string `json:"nav_grl3y"`
			NavGrl5Y              string `json:"nav_grl5y"`
			NavGrbase             string `json:"nav_grbase"`
			SrankL1M              string `json:"srank_l1m"`
			SrankL3M              string `json:"srank_l3m"`
			SrankL6M              string `json:"srank_l6m"`
			SrankLty              string `json:"srank_lty"`
			SrankL1Y              string `json:"srank_l1y"`
			SrankL3Y              string `json:"srank_l3y"`
			SrankL5Y              string `json:"srank_l5y"`
			NavGrowth             string `json:"nav_growth"`
			AnnualPerformanceList []struct {
				Period string `json:"period"`
				Nav    string `json:"nav"`
				Rank   string `json:"rank"`
			} `json:"annual_performance_list"`
			YieldHistory []struct {
				Yield string `json:"yield"`
				Name  string `json:"name"`
			} `json:"yield_history"`
		} `json:"fund_derived"`
		FundRates struct {
			SubscribeRate     string `json:"subscribe_rate"`
			DeclareRate       string `json:"declare_rate"`
			Discount          string `json:"discount"`
			SubscribeDiscount string `json:"subscribe_discount"`
			DeclareDiscount   string `json:"declare_discount"`
		} `json:"fund_rates"`
		OpFund struct {
			BannerImg string `json:"banner_img"`
			Tips      string `json:"tips"`
			FundTags  []struct {
				Category string `json:"category"`
				Name     string `json:"name"`
			} `json:"fund_tags"`
		} `json:"op_fund"`
		Yield              string `json:"yield"`
		YieldName          string `json:"yield_name"`
		GrowthDay          string `json:"growth_day"`
		Sales              string `json:"sales"`
		Tips               string `json:"tips"`
		TypeDesc           string `json:"type_desc"`
		RatingSource       string `json:"rating_source"`
		RatingDesc         string `json:"rating_desc"`
		FollowerCount      int    `json:"follower_count"`
		StatusCount        int    `json:"status_count"`
		StockPositionNames string `json:"stock_position_names"`
		AgentSell          bool   `json:"agent_sell"`
		TradeReason        struct {
			ShowButton      bool `json:"show_button"`
			WithdrawDisplay bool `json:"withdraw_display"`
		} `json:"trade_reason"`
		GrowthDefaultPeriod string `json:"growth_default_period"`
		FirHeaderBaseData   []struct {
			DataName        string  `json:"data_name"`
			DataValueStr    string  `json:"data_value_str"`
			DataValueNumber float64 `json:"data_value_number"`
			DataHaveColour  bool    `json:"data_have_colour"`
		} `json:"fir_header_base_data"`
		SecHeaderBaseData []struct {
			DataName        string  `json:"data_name"`
			DataValueStr    string  `json:"data_value_str"`
			DataValueNumber float64 `json:"data_value_number,omitempty"`
			DataHaveColour  bool    `json:"data_have_colour"`
			DataExtend      string  `json:"data_extend,omitempty"`
		} `json:"sec_header_base_data"`
		BaseDataTip []struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		} `json:"base_data_tip"`
		NavTabList []struct {
			NavTabName   string `json:"nav_tab_name"`
			NavTabValue  string `json:"nav_tab_value"`
			NavGrowth    string `json:"nav_growth"`
			NavTabExtent string `json:"nav_tab_extent,omitempty"`
		} `json:"nav_tab_list"`
		DisplayAnnualPerformance bool `json:"display_annual_performance"`
		BenchmarkIndex           []struct {
			Symbol     string `json:"symbol"`
			SymbolName string `json:"symbol_name"`
		} `json:"benchmark_index"`
		InvestOrientation    string `json:"invest_orientation"`
		InvestTarget         string `json:"invest_target"`
		PerformanceBenchMark string `json:"performance_bench_mark"`
		RecordVersion        string `json:"record_version"`
		PensionFund          bool   `json:"pension_fund"`
	} `json:"data"`
	ResultCode int `json:"result_code"`
}

func parseFundSummary(body []byte) (entity.FundSummary, error) {
	fundSummary := entity.FundSummary{}
	response := FundSummaryResponse{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return fundSummary, err
	}

	err = copier.Copy(&fundSummary, response.Data)
	return fundSummary, err
}

type ListMarketResponse struct {
	Data struct {
		Items []struct {
			Market struct {
				StatusId     int         `json:"status_id"`
				Region       string      `json:"region"`
				Status       string      `json:"status"`
				TimeZone     string      `json:"time_zone"`
				TimeZoneDesc interface{} `json:"time_zone_desc"`
				DelayTag     int         `json:"delay_tag"`
			} `json:"market"`
			Quote struct {
				Symbol             string   `json:"symbol"`
				Code               string   `json:"code"`
				Exchange           string   `json:"exchange"`
				Name               string   `json:"name"`
				Type               int      `json:"type"`
				SubType            *string  `json:"sub_type"`
				Status             int      `json:"status"`
				Current            float64  `json:"current"`
				Currency           string   `json:"currency"`
				Percent            float64  `json:"percent"`
				Chg                float64  `json:"chg"`
				Timestamp          int64    `json:"timestamp"`
				Time               int64    `json:"time"`
				LotSize            int      `json:"lot_size"`
				TickSize           float64  `json:"tick_size"`
				Open               float64  `json:"open"`
				LastClose          float64  `json:"last_close"`
				High               float64  `json:"high"`
				Low                float64  `json:"low"`
				AvgPrice           float64  `json:"avg_price"`
				Volume             int64    `json:"volume"`
				Amount             float64  `json:"amount"`
				TurnoverRate       float64  `json:"turnover_rate"`
				Amplitude          float64  `json:"amplitude"`
				MarketCapital      float64  `json:"market_capital"`
				FloatMarketCapital *float64 `json:"float_market_capital"`
				TotalShares        int64    `json:"total_shares"`
				FloatShares        int64    `json:"float_shares"`
				IssueDate          int64    `json:"issue_date"`
				LockSet            *int     `json:"lock_set"`
				CurrentYearPercent float64  `json:"current_year_percent"`
			} `json:"quote"`
			Others struct {
				CybSwitch bool `json:"cyb_switch"`
			} `json:"others"`
			Tags []interface{} `json:"tags"`
		} `json:"items"`
		ItemsSize int `json:"items_size"`
	} `json:"data"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

func parseListMarket(body []byte) (entity.Markets, error) {
	response := ListMarketResponse{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	markets := make(entity.Markets, 0, len(response.Data.Items))

	for i := range response.Data.Items {
		item := response.Data.Items[i]
		market := entity.Market{
			Name:         item.Quote.Name,
			Current:      item.Quote.Current,
			Float:        item.Quote.Chg,
			FloatPercent: item.Quote.Percent,
		}

		markets = append(markets, market)
	}

	return markets, nil
}
