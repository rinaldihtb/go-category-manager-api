package services

import (
	"category-manager-api/models"
	"category-manager-api/repositories"
	"time"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetSummaryReport(datenow time.Time) (*models.SummaryReport, error) {
	return s.repo.GetSummaryReport(datenow)
}
