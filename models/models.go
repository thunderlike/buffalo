package models

type Model struct {
	ID         int32 `gorm:"primary_key" json:"id"`
	CreatedOn  int32 `json:"created_on"`
	ModifiedOn int32 `json:"modified_on"`
	DeletedOn  int32 `json:"deleted_on"`
}
