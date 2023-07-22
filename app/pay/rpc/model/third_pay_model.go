package genModel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ThirdPayModel = (*customThirdPayModel)(nil)

type (
	// ThirdPayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customThirdPayModel.
	ThirdPayModel interface {
		thirdPayModel
	}

	customThirdPayModel struct {
		*defaultThirdPayModel
	}
)

// NewThirdPayModel returns a model for the database table.
func NewThirdPayModel(conn sqlx.SqlConn) ThirdPayModel {
	return &customThirdPayModel{
		defaultThirdPayModel: newThirdPayModel(conn),
	}
}
