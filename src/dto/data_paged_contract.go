package dto

type Pagination struct {
	Current      int      		`json:"current"`
	PerPage      int      		`json:"per_page"`
	TotalPages   int      		`json:"total_page"`
	TotalRecords int      		`json:"total_records"`
	Data         interface{} 	`json:"data"`
}