# Go Makeup!

A simple wrapper for the http://makeup-api.herokuapp.com/ api!

### Structures

```go
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
```

### Methods

`makeup.GetMakeUp(query ...map[string]string) returns []Product`

The `GetMakeUp` method accepts a map that is made up of the params for the api. Even though it is variadic, the method only really accepts the one map.

Example:
```go
makeupRes :=  makeup.GetMakeup(map[string]string{
	"product_type": "blush",
	"product_category": "powder",
})
```
