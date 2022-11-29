package global

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel 数据表基本属性
type BaseModel struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

// BaseAudit 审计相关属性
type BaseAudit struct {
	CreatedBy uint `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint `gorm:"column:deleted_by;comment:删除者"`
}
