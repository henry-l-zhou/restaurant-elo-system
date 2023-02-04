package nearbysearch

type Nearby struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	NextPageToken    string        `json:"next_page_token"`
	Results          []Results     `json:"results"`
	Status           string        `json:"status"`
}
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type Northeast struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type Southwest struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type Viewport struct {
	Northeast Northeast `json:"northeast"`
	Southwest Southwest `json:"southwest"`
}
type Geometry struct {
	Location Location `json:"location"`
	Viewport Viewport `json:"viewport"`
}
type Photos struct {
	Height           int      `json:"height"`
	HTMLAttributions []string `json:"html_attributions"`
	PhotoReference   string   `json:"photo_reference"`
	Width            int      `json:"width"`
}
type OpeningHours struct {
	OpenNow bool `json:"open_now"`
}
type PlusCode struct {
	CompoundCode string `json:"compound_code"`
	GlobalCode   string `json:"global_code"`
}
type Results struct {
	Geometry            Geometry     `json:"geometry"`
	Icon                string       `json:"icon"`
	IconBackgroundColor string       `json:"icon_background_color"`
	IconMaskBaseURI     string       `json:"icon_mask_base_uri"`
	Name                string       `json:"name"`
	Photos              []Photos     `json:"photos,omitempty"`
	PlaceID             string       `json:"place_id"`
	Reference           string       `json:"reference"`
	Scope               string       `json:"scope"`
	Types               []string     `json:"types"`
	Vicinity            string       `json:"vicinity"`
	BusinessStatus      string       `json:"business_status,omitempty"`
	OpeningHours        OpeningHours `json:"opening_hours,omitempty"`
	PlusCode            PlusCode     `json:"plus_code,omitempty"`
	PriceLevel          int          `json:"price_level,omitempty"`
	Rating              float64      `json:"rating,omitempty"`
	UserRatingsTotal    int          `json:"user_ratings_total,omitempty"`
}
