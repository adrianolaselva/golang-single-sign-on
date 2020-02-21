package common

type PaginationCommon struct {
	Current      int      		`json:"current"`
	TotalPages   int      		`json:"total_page"`
	PerPage      int      		`json:"per_page"`
	TotalRecords int      		`json:"total_records"`
	Data         interface{} 	`json:"data"`
}