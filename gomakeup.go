package gomakeup

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

//Product struct
type Product struct {
	ID            int            `json:"id"`
	Brand         string         `json:"brand"`
	Name          string         `json:"name"`
	Price         string         `json:"price"`
	ImageLink     string         `json:"image_link"`
	ProductLink   string         `json:"product_link"`
	WebsiteLink   string         `json:"website_link"`
	Description   string         `json:"description"`
	Rating        float32        `json:"rating"`
	Category      string         `json:"category"`
	ProductType   string         `json:"product_type"`
	TagList       []string       `json:"tag_list"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     string         `json:"updated_at"`
	ProductAPIURL string         `json:"product_api_url"`
	ProductColors []ProductColor `json:"product_colors"`
}

//ProductColor struct
type ProductColor struct {
	HexValue   string `json:"hex_value"`
	ColourName string `json:"colour_name"`
}

const apiURL string = "http://makeup-api.herokuapp.com/api/v1/products.json"

// GetMakeup used to get makeup from API
func GetMakeup(query ...map[string]string) []Product {

	queryParams := url.Values{}

	for key, value := range query[0] {
		queryParams.Set(key, value)
	}

	makeupRes, _ := http.Get(apiURL + "?" + queryParams.Encode())
	defer makeupRes.Body.Close()

	makeBytes, _ := ioutil.ReadAll(makeupRes.Body)

	var MakeupProducts []Product

	err := json.Unmarshal(makeBytes, &MakeupProducts)

	if err != nil {
		panic(err)
	}

	return MakeupProducts
}
