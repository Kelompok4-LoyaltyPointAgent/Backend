package analytics_service

import (
	"math/rand"
	"time"

	"github.com/kelompok4-loyaltypointagent/backend/dto/response"
	"github.com/kelompok4-loyaltypointagent/backend/repositories/analytics_repository"
)

type AnalyticsService interface {
	Analytics() (*response.AnalyticsResponse, error)
}

type analyticsService struct {
	analyticsRepository analytics_repository.AnalyticsRepository
}

func NewAnalyticsService(analyticsRepository analytics_repository.AnalyticsRepository) AnalyticsService {
	return &analyticsService{analyticsRepository}
}

func (s *analyticsService) Analytics() (*response.AnalyticsResponse, error) {
	year := time.Now().Year()

	transactions, err := s.analyticsRepository.RecentTransactions(year, 5)
	if err != nil {
		return nil, err
	}

	analyticsResponse := response.AnalyticsResponse{
		Year:                year,
		VisitorsCount:       rand.Intn(69420), // dummy data
		SalesCount:          s.analyticsRepository.SalesCount(year),
		AppInstallsCount:    rand.Intn(69420), // dummy data
		Income:              s.analyticsRepository.Income(year),
		TransactionsByMonth: s.analyticsRepository.TransactionsByMonth(year),
		TransactionsByType:  s.analyticsRepository.TransactionsByType(year),
		RecentTransactions:  *response.NewTransactionsResponse(transactions),
	}

	return &analyticsResponse, nil
}
