package product

import (
	"github.com/mitchellh/mapstructure"
)

func wheneverProductChanged_ItIsNotCommon(data map[string]interface{}){
	
	event := NewProductChanged()
	mapstructure.Decode(data,&event)

	ItIsNotCommon(event);
}

