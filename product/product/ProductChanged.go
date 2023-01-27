package product

import (
	"time"
)

type ProductChanged struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	Name string `json:"name" type:"string"` 
	Stock int `json:"stock" type:"int"` 
	
}

func NewProductChanged() *ProductChanged{
	event := &ProductChanged{EventType:"ProductChanged", TimeStamp:time.Now().String()}

	return event
}
