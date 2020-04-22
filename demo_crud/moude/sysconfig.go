package moude

import "demo_crud/pkgs/enums"

type SysConfig struct {
	ID         int                 `gorm:"column:id;type:int" json:"id"`
	ParamKey   string              `gorm:"column:param_key;type:varchar(45)" json:"param_key" validate:"required"`
	ParamValue string              `gorm:"column:param_value;type:varchar(100)" json:"param_value" validate:"required"`
	ParamType  enums.Config        `gorm:"column:param_type;type:int" json:"param_type" validate:"required"`
}
type GetType struct {
	ParamType []enums.Config        `json:"param_type"`
}
