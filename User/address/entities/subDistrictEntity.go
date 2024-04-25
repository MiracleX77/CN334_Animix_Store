package entities

type (
	SubDistrict struct {
		Id         uint   `json:"id" gorm:"primaryKey"`
		NameTh     string `json:"name_th"`
		NameEn     string `json:"name_en"`
		PostCode   string `json:"post_code"`
		DistrictId uint   `json:"district_id"`
	}
)
