package analytics_service

import (
	"log"
	"math/rand"
	"time"

	"github.com/kelompok4-loyaltypointagent/backend/cachedrepositories/cached_analytics_repository"
	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/analytics_repository"
)

type AnalyticsService interface {
	Analytics() (*response.AnalyticsResponse, error)
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
			log.Printf("Caching failed: %s", err)
		}
	}

	var income float64
	if s.cachedAnalyticsRepository.CheckIncome(year) {
		income = s.cachedAnalyticsRepository.Income(year)
	} else {
		income = s.analyticsRepository.Income(year)
		if err := s.cachedAnalyticsRepository.SetIncome(year, income); err != nil {
			log.Printf("Caching failed: %s", err)
		}
	}

	transactions, err := s.analyticsRepository.RecentTransactions(year, 5)
	if err != nil {
		return nil, err
	}

	analyticsResponse := response.AnalyticsResponse{
		Year:                year,
		VisitorsCount:       rand.Intn(69420), // dummy data
		SalesCount:          salesCount,
		AppInstallsCount:    rand.Intn(69420), // dummy data
		Income:              income,
		TransactionsByMonth: s.analyticsRepository.TransactionsByMonth(year),
		TransactionsByType:  s.analyticsRepository.TransactionsByType(year),
		RecentTransactions:  *response.NewTransactionsResponse(transactions),
	}

	return &analyticsResponse, nil
}
