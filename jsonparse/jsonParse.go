package jasonParse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ID struct {
	ID string `json:"$oid"`
}

type Element struct {
	ID             ID     `json:"_id"`
	Capitalization string `json:"capitalization"`
}

type Category struct {
	Сurrency []string `json:"category"`
}

type ChangePrice struct {
	Change        string `json: "change"`
	ChangePercent string `json: "change_percent"`
	MaxPrice      string `json:"max_price"`
	MinPrice      string `json:"min_price"`
	Period        string `json:"period"`
	UpdateDate    string `json:"update_date"`
}
type CoefIssue struct {
	CoefIssue string `json:"coef issue"`
}
type CoefSale struct {
	CoefSale string `json:"coef_sale"`
}
type CurrentIssue struct {
	CurrentIssue string `json:"current issue"`
}
type CurrenRate struct {
	UpdateAt string `json"update_at"`
	Value    string `"value"`
}
type Dynamic struct {
	Change        string `json:"change"`
	ChangePercent string `json"change_percent"`
	MaxPrice      string `json:"max_price"`
	MinPrice      string `json:"min_price"`
	Period        string `json:"period"`
	UpdateDate    string `json:"update_date"`
}

type Event struct {
	Events struct {
		CountEvent     int `json:"count_event"`
		Index24H       int `json:"index_24_h"`
		IndexOfRating  int `json:"index_of_rating"`
		IndexOfReviews int `json:"index_of_reviews"`
	}
	UpdateAt string `json:"update_at"`
}

//help
type Fund struct {
	Funds    []string `json:"funds"`
	UpdateAt string   `json"update_at"`
}

type Github struct {
	Contributors int    `json:"contributors"`
	Stargazers   int    `json:"stargazers"`
	Team         int    `json:"tem"`
	UpdateAt     string `json:"update_at"`
}

type Name struct {
	Name string `json:"name"`
}
type Platform struct {
	Platform []string `json:"platform"`
}
type Presale struct {
	Presale []string `json:"presale"`
}
type Rank struct {
	Rank int `json:"rank"`
}
type StockExchange struct {
	StockExchange string `json:"stock_exchange"`
}

//???
type TableQuarterlyYield struct {
	Q1   string `json:"Q1"`
	Q2   string `json:"Q2"`
	Q3   string `json:"Q3"`
	Q4   string `json:"Q4"`
	Year int    `json:"year"`
}
type Tags struct {
	Tags []string `json:"tags"`
}
type TotalIssue struct {
	TotalIssue int `json:"total_issue"`
}
type UpdateAt struct {
	UpdateAt string `json"update_at"`
}
type Urls struct {
	GithubUrl  string `json:"github_url"`
	TwitterUrl string `json:"twitter_url"`
}
type VolumePreDay struct {
	VolumePreDay string `json:"volume_pre_day"`
}

func JasonParse(jsonFile string) ([]Urls, error) {

	byteValue, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	date := make([]Urls, 0)

	err = json.Unmarshal(byteValue, &date)
	if err != nil {
		return nil, err
	}

	fmt.Println(date)
	return date, nil

}
