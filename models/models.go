package models

import "time"

type DogBreed struct {
	ID               int64  `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weightLowLbs"`
	WeightHighLbs    int    `json:"weightHighLbs"`
	AverageWeight    int    `json:"averageWeight"`
	Lifespan         int    `json:"lifespan"`
	Details          string `json:"details"`
	AlternateName    string `json:"alternateName"`
	GeographicOrigin string `json:"geographicOrigin"`
}

type CatBreed struct {
	ID               int64  `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weightLowLbs"`
	WeightHighLbs    int    `json:"weightHighLbs"`
	AverageWeight    int    `json:"averageWeight"`
	Lifespan         int    `json:"lifespan"`
	Details          string `json:"details"`
	AlternateName    string `json:"alternateName"`
	GeographicOrigin string `json:"geographicOrigin"`
}

type Dog struct {
	ID               int64     `json:"id"`
	DogName          string    `json:"dogName"`
	BreedID          int64     `json:"breedId"`
	BreederID        int64     `json:"breederId"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"dateOfBirth"`
	SpayedOrNeutered int       `json:"spayedOrNeutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            DogBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Cat struct {
	ID               int64     `json:"id"`
	CatName          string    `json:"catName"`
	BreedID          int64     `json:"breedId"`
	BreederID        int64     `json:"breederId"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"dateOfBirth"`
	SpayedOrNeutered int       `json:"spayedOrNeutered"`
	Description      string    `json:"description"`
	Weight           int       `json:"weight"`
	Breed            CatBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Breeder struct {
	ID            int64       `json:"id"`
	BreederName   string      `json:"breederName"`
	Address       string      `json:"address"`
	City          string      `json:"city"`
	ProvinceState string      `json:"provinceState"`
	Country       string      `json:"country"`
	Zip           string      `json:"zip"`
	Phone         string      `json:"phone"`
	Email         string      `json:"email"`
	Active        int         `json:"active"`
	DogBreed      []*DogBreed `json:"dogBreeds"`
	CatBreed      []*CatBreed `json:"catBreeds"`
}

type Pet struct {
	Species     string `json:"species"`
	Breed       string `json:"breed"`
	MinWeight   int    `json:"minWeight"`
	MaxWeight   int    `json:"maxWeight"`
	Description string `json:"description"`
	Lifespan    int    `json:"lifespan"`
}
