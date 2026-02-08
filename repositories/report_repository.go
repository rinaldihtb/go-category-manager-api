package repositories

import (
	"category-manager-api/models"
	"database/sql"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetSummaryReport(datenow time.Time) (*models.SummaryReport, error) {
	var (
		res *models.SummaryReport
	)

	// #1 Get Today Transactions
	query := "SELECT td.id, p.name, td.quantity, td.subtotal FROM transaction_details td LEFT JOIN products p ON p.id = td.product_id WHERE DATE(td.created_at) = DATE($1)"
	rows, err := repo.db.Query(query, datenow)
	if err != nil {
		return nil, err
	}

	// #1 Count Today Transactions
	query2 := "SELECT COUNT(*) FROM transactions WHERE DATE(created_at) = DATE($1)"
	var totalTransaction int
	err = repo.db.QueryRow(query2, datenow).Scan(&totalTransaction)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	transactionDetails := make([]models.TransactionDetail, 0)
	for rows.Next() {
		var td models.TransactionDetail
		err := rows.Scan(&td.ID, &td.ProductName, &td.Quantity, &td.Subtotal)
		if err != nil {
			return nil, err
		}
		transactionDetails = append(transactionDetails, td)
	}

	// #2 Initiate Total of Revenue, Transaction, and best sellers item
	res = &models.SummaryReport{
		TotalRevenue:     0,
		TotalTransaction: totalTransaction,
		BestSeller: models.BestSellerProduct{
			Name:         "",
			SoldQuantity: 0,
		},
	}

	// #3 Loop each record, then calculate everything
	for _, item := range transactionDetails {
		res.TotalRevenue += item.Subtotal
		if item.Quantity > res.BestSeller.SoldQuantity {
			res.BestSeller = models.BestSellerProduct{
				Name:         item.ProductName,
				SoldQuantity: item.Quantity,
			}
		}
	}

	return res, nil
}
