package types

type CommonFilters struct {
	Page       int32
	Limit      int32
	Search     string
	SortBy     string
	SortOrder  string
	StartDate  string
	EndDate    string
	Status     string
	Categories []string
	Tags       []string
	MinPrice   float64
	MaxPrice   float64
	StoreId    string
	UserId     string
	IsActive   bool
	Type       string
	Priority   int32
	Location   string
	Radius     float64
	Features   []string
	Brands     []string
	Colors     []string
	Sizes      []string
	Materials  []string
	Rating     float64
	InStock    bool
	OnSale     bool
	IsFeatured bool
}