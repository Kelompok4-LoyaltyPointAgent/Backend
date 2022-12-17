package response

import "github.com/kelompok4-loyaltypointagent/backend/repositories/analytics_repository"

type AnalyticsResponse struct {
	Year                int                                      `json:"year"`
	VisitorsCount       int                                      `json:"visitors_count"`
	SalesCount          int                                      `json:"sales_count"`
	AppInstallsCount    int                                      `json:"app_installs_count"`
	Income              float64                                  `json:"income"`
	TransactionsByMonth analytics_repository.TransactionsByMonth `json:"transactions_by_month"`
	TransactionsByType  analytics_repository.TransactionsByType  `json:"transactions_by_type"`
	RecentTransactions  []TransactionResponse                    `json:"recent_transactions"`
}

type DataForManageStockAdmin struct {
	TotalProduct   uint `json:"totalProduct"`
	TotalProvider  uint `json:"totalProvider"`
	CashoutBalance uint `json:"cashoutBalance"`
}
