package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	URL     = "https://dummyjson.com/products/1"
	product = Product{}
)

type Product struct {
	Id                 int     `json:"id"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	Price              float32 `json:"price"`
	DiscountPercentage float32 `json:"discountPercentage"`
	Rating             float32 `json:"rating"`
	Stock              int     `json:"stock"`
	Brand              string  `json:"brand"`
	Category           string  `json:"category"`
	Thumbnail          string  `json:"thumbnail"`
}

func (p Product) toString() string {
	return fmt.Sprintf(
		"Id: %d\nTitle: %s\nDescription: %s\nPrice: %.2f\nDiscountPercentage: %.2f\nRating: %.2f\nStock: %d\nBrand: %s\nCategory: %s\nThumbnail: %s",
		p.Id, p.Title, p.Description, p.Price, p.DiscountPercentage, p.Rating, p.Stock, p.Brand, p.Category, p.Thumbnail,
	)
}

func getProduct() Product {
	callResponse, _ := http.Get(URL)
	defer callResponse.Body.Close()

	response, _ := io.ReadAll(callResponse.Body)
	json.Unmarshal(response, &product)
	wg.Done()
	return product
}
