package model

// Pesanan is representation to make a menu was ordered
type Pesanan struct {
	MenuID int
	Qty    int
}

// Transaksi is representation to make a new one transaction
type Transaksi struct {
	MejaID int
	Notes  string
	Pesan  []Pesanan
}

// MenuOrdered is to representation menu was ordered by customers
type MenuOrdered struct {
	Nama  string
	Qty   int
	Harga int
	Total int
}

// Bill is representation for billing after transaction
type Bill struct {
	MejaID     int
	Menus      []MenuOrdered
	GrandTotal int
}
