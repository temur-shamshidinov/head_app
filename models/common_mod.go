package models

type Common struct {
	TableName  string `json:"table_name"`
	ColumnName string `json:"column_name"`
	ExpValue   any    `json:"exp_value"`
}

type GetListReq struct {
	Page   int32  `json:"page"`
	Limit  int32  `json:"limit"`
	Search string `json:"search"`
}

type CheckExistsResp struct {
	IsExists bool   `json:"is_exists"`
	Status   string `json:"status"`
}
