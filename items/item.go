package items

type Item struct {
	Id                string    `json:"id"`
	Seller            int64     `json:"seller"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Picture           []Picture `json:"pictures"`
	Video             string    `json:"video"`
	Price             float32   `json:"price"`
	AvailableQuantity int       `json:"available_quantity"`
	SoldQuantity      int       `json:"sold_quantity"`
	Status            int       `json:"status"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}

type Picture struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}
