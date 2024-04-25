package entities

type (
	District struct {
		Id         uint   `json:"id" gorm:"primaryKey"`
		NameTh     string `json:"name_th"`
		NameEn     string `json:"name_en"`
		ProvinceId uint   `json:"province_id"`
	}
)
