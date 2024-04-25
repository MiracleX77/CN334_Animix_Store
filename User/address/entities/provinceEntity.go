package entities

type (
	Province struct {
		Id     uint   `json:"id" gorm:"primaryKey"`
		NameTh string `json:"name_th"`
		NameEn string `json:"name_en"`
	}
)
