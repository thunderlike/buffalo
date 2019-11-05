package models

type AccountLogin struct {
	Model
	Name       string `json:"name"`
	AType      int8   `json:"a_type"`
	Password   string `json:"password"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int8   `json:"state"`
}
