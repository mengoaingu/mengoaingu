package dbutils

type Model struct {
	CreatedAt int64 `gorm:"column:created_at;autoCreateTime:milli;index"`
	UpdatedAt int64 `gorm:"column:updated_at;autoUpdateTime:milli"`
}
