package analytics_service

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/kelompok4-loyaltypointagent/backend/cachedrepositories/cached_analytics_repository"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/helper"
	"github.com/kelompok4-loyaltypointagent/backend/models"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/analytics_repository"
)

type AnalyticsService interface {
	Analytics() (*response.AnalyticsResponse, error)
	DataForManageStockAdmin() (*response.DataForManageStockAdmin, error)
}

type analyticsService struct {
	analyticsRepository       analytics_repository.AnalyticsRepository
	cachedAnalyticsRepository cached_analytics_repository.CachedAnalyticsRepository
}

func NewAnalyticsService(
	analyticsRepository analytics_repository.AnalyticsRepository,
	cachedAnalyticsRepository cached_analytics_repository.CachedAnalyticsRepository,
) AnalyticsService {
	return &analyticsService{analyticsRepository, cachedAnalyticsRepository}
}

func (s *analyticsService) Analytics() (*response.AnalyticsResponse, error) {
	year := time.Now().Year()

	var salesCount int
	if s.cachedAnalyticsRepository.CheckSalesCount(year) {
		salesCount = s.cachedAnalyticsRepository.SalesCount(year)
	} else {
		salesCount = s.analyticsRepository.SalesCount(year)
		if err := s.cachedAnalyticsRepository.SetSalesCount(year, salesCount); err != nil {
			log.Printf("Redis error: %s", err)
		}
	}

	var income float64
	if s.cachedAnalyticsRepository.CheckIncome(year) {
		income = s.cachedAnalyticsRepository.Income(year)
	} else {
		income = s.analyticsRepository.Income(year)
		if err := s.cachedAnalyticsRepository.SetIncome(year, income); err != nil {
			log.Printf("Redis error: %s", err)
		}
	}

	var transactionsByMonth analytics_repository.TransactionsByMonth
	if s.cachedAnalyticsRepository.CheckTransactionsByMonth(year) {
		transactionsByMonth = s.cachedAnalyticsRepository.TransactionsByMonth(year)
	} else {
		transactionsByMonth = s.analyticsRepository.TransactionsByMonth(year)

		data, err := json.Marshal(transactionsByMonth)
		if err != nil {
			log.Printf("JSON error: %s", err)
		}

		if err := s.cachedAnalyticsRepository.SetTransactionsByMonth(year, string(data)); err != nil {
			log.Printf("Redis error: %s", err)
		}
	}

	var transactionsByType analytics_repository.TransactionsByType
	if s.cachedAnalyticsRepository.CheckTransactionsByType(year) {
		transactionsByType = s.cachedAnalyticsRepository.TransactionsByType(year)
	} else {
		transactionsByType = s.analyticsRepository.TransactionsByType(year)

		data, err := json.Marshal(transactionsByType)
		if err != nil {
			log.Printf("JSON error: %s", err)
		}

		if err := s.cachedAnalyticsRepository.SetTransactionsByType(year, string(data)); err != nil {
			log.Printf("Redis error: %s", err)
		}
	}

	var transactions []models.Transaction
	var err error
	if s.cachedAnalyticsRepository.CheckRecentTransactions(year) {
		transactions, err = s.cachedAnalyticsRepository.RecentTransactions(year, 5)
		if err != nil {
			return nil, err
		}
	} else {
		transactions, err = s.analyticsRepository.RecentTransactions(year, 5)
		if err != nil {
			return nil, err
		}

		data, err := json.Marshal(transactions)
		if err != nil {
			log.Printf("JSON error: %s", err)
		}

		if err := s.cachedAnalyticsRepository.SetRecentTransactions(year, string(data)); err != nil {
			log.Printf("Redis error: %s", err)
		}
	}

	analyticsResponse := response.AnalyticsResponse{
		Year:                year,
		VisitorsCount:       rand.Intn(69420), // dummy data
		SalesCount:          salesCount,
		AppInstallsCount:    rand.Intn(69420), // dummy data
		Income:              income,
		TransactionsByMonth: transactionsByMonth,
		TransactionsByType:  transactionsByType,
		RecentTransactions:  *response.NewTransactionsResponse(transactions),
	}

	return &analyticsResponse, nil
}

func (s *analyticsService) DataForManageStockAdmin() (*response.DataForManageStockAdmin, error) {
	var data *response.DataForManageStockAdmin = &response.DataForManageStockAdmin{}

	if s.cachedAnalyticsRepository.CheckDataInStock() {
		dataRedis, err := s.cachedAnalyticsRepository.GetDataInStock()
		if err != nil {
			return nil, err
		}

		log.Println(*dataRedis)

		err = json.Unmarshal([]byte(*dataRedis), data)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

	} else {
		totalProduct, err := s.analyticsRepository.ProductCount()
		if err != nil {
			return nil, err
		}

		cashoutBalance, err := helper.GetBalance()
		if err != nil {
			return nil, err
		}

		dataResponse := response.DataForManageStockAdmin{
			CashoutBalance: *cashoutBalance,
			TotalProduct:   uint(totalProduct),
			TotalProvider:  6,
		}

		data = &dataResponse

		dataMarshal, err := json.Marshal(dataResponse)
		if err != nil {
			log.Printf("JSON error: %s", err)
		}

		if err := s.cachedAnalyticsRepository.SetProductCount(string(dataMarshal)); err != nil {
			log.Printf("Redis error: %s", err)
		}
	}

	return data, nil
}
