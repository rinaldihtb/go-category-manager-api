package models

type SummaryReport struct {
	TotalRevenue     int               `json:"total_revenue"`
	TotalTransaction int               `json:"total_transaksi"`
	BestSeller       BestSellerProduct `json:"produk_terlaris"`
}

type BestSellerProduct struct {
	Name         string `json:"nama"`
	SoldQuantity int    `json:"qty_terjual"`
}
