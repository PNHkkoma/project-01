package models

type StoreData struct {
	StoreName   string `form:"store_name" json:"store_name"`
	SubCategory string `form:"sub_category" json:"sub_category"`
	Category    string `form:"category" json:"category"`
	Location    string `form:"location" json:"location"`
	ImageURL    string `form:"image_url" json:"image_url"`
	ImageName   string `form:"image_name" json:"image_name"`
	MallName    string `form:"mall_name" json:"mall_name"`
}

type StoreDataByCategory struct {
	Category string `form:"category" json:"category"`
}

type StoreDataSearchKey struct {
	QueryKey string `form:"query" json:"query"`
}
