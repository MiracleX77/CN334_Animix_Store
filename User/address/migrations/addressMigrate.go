package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MiracleX77/CN334_Animix_Store/address/entities"
	"github.com/MiracleX77/CN334_Animix_Store/configs"
	"github.com/MiracleX77/CN334_Animix_Store/database"
)

func main() {
	cfg := configs.GetConfig()
	db := database.NewPostgresDatabase(&cfg)
	//addressMigrate(db)
	//fetchAndSaveDistrict(db)
	fetchAndSaveSubDistrict(db)
}

func addressMigrate(db database.Database) {
	db.GetDb().AutoMigrate(&entities.Address{})
	db.GetDb().AutoMigrate(&entities.Province{})
	db.GetDb().AutoMigrate(&entities.District{})
	db.GetDb().AutoMigrate(&entities.SubDistrict{})
}

func fetchAndSaveDistrict(db database.Database) {
	//url := "https://raw.githubusercontent.com/kongvut/thai-province-data/master/api_province.json" // Replace with the actual API URL
	url := "https://raw.githubusercontent.com/kongvut/thai-province-data/master/api_amphure.json"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)
	var provinces []entities.District
	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	// // Optional: log the results to verify
	// for _, province := range provinces {
	// 	log.Printf("Read province ID: %d, Name: %s", province.Id, province.NameTh)
	// }

	saveProvinces(db, provinces)
}

func fetchAndSaveSubDistrict(db database.Database) {
	type apiSubDistrict struct {
		Id        uint    `json:"id"`
		ZipCode   uint    `json:"zip_code"`
		NameTh    string  `json:"name_th"`
		NameEn    string  `json:"name_en"`
		AmphureId uint    `json:"amphure_id"`
		CreatedAt string  `json:"created_at"`
		UpdatedAt string  `json:"updated_at"`
		DeletedAt *string `json:"deleted_at"`
	}

	url := "https://raw.githubusercontent.com/kongvut/thai-province-data/master/api_tambon.json" // Replace with the actual API URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()
	var apiSubDistricts []apiSubDistrict
	if err := json.NewDecoder(resp.Body).Decode(&apiSubDistricts); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	var subDistricts []entities.SubDistrict
	for _, apiSubDistrict := range apiSubDistricts {
		//convert unit to string
		zipcode := fmt.Sprintf("%d", apiSubDistrict.ZipCode)
		subDistrict := entities.SubDistrict{
			Id:         apiSubDistrict.Id,
			NameTh:     apiSubDistrict.NameTh,
			NameEn:     apiSubDistrict.NameEn,
			PostCode:   zipcode,
			DistrictId: apiSubDistrict.AmphureId,
		}
		subDistricts = append(subDistricts, subDistrict)
	}
	saveSubDistricts(db, subDistricts)
}

func saveSubDistricts(db database.Database, subDistricts []entities.SubDistrict) {
	for _, subDistrict := range subDistricts {
		result := db.GetDb().Create(&subDistrict)
		if result.Error != nil {
			log.Fatalf("Error saving sub-district ID %d to database: %v", subDistrict.Id, result.Error)
		}
	}
	log.Println("All sub-districts saved successfully")
}

func saveProvinces(db database.Database, provinces []entities.District) {
	for _, province := range provinces {
		result := db.GetDb().Create(&province)
		if result.Error != nil {
			log.Fatalf("Error saving province ID %d to database: %v", province.Id, result.Error)
		}
	}
	log.Println("All provinces saved successfully")
}
