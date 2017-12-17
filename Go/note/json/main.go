package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	OrdId		int32
	BrandId		uint64
	SkuId		uint64
	SkuCount	int32
	SkuSmlCatId	uint64
}

func main()  {
	//sku := "[{\"ord_id\":1,\"brand_id\":4902330214759531549,\"sku_id\":4902353547747657669,\"sku_count\":2,\"sku_sml_cat_id\":4902330194594890777},{\"ord_id\":2,\"brand_id\":4902330214759531549,\"sku_id\":4902353547135289243,\"sku_count\":2,\"sku_sml_cat_id\":4902330194628445211}]"
	sku := `[{"ord_id":1,"brand_id":4902330214759531549,"sku_id":4902353547747657669,"sku_count":2,"sku_sml_cat_id":4902330194594890777},{"ord_id":2,"brand_id":4902330214759531549,"sku_id":4902353547135289243,"sku_count":2,"sku_sml_cat_id":4902330194628445211}]`

	skuData, err := Unmarshal(sku)
	if err != nil {
		return
	}

	fmt.Printf("====================================\n")
	str, err := Marshal(skuData)
	if err != nil {
		return
	}
	fmt.Printf("str = %s\n", str)
}

func Unmarshal(str string) ([]*Data, error)  {
	sku_data := []*Data{}
	//sku_data := []Data{}

	if err := json.Unmarshal([]byte(str),&sku_data); err != nil {
		fmt.Printf("json.Unmarshal() err: %s\n", err.Error())
		return nil, err
	}

	for _, item := range sku_data {
		fmt.Printf("item = %v\n",item)
	}

	return sku_data,nil
}

func Marshal(sku_data []*Data) (string, error) {
	buf, err := json.Marshal(sku_data)
	if err != nil {
		fmt.Printf("json.Marshal() err: %s\n", err.Error())
		return "", err
	}
		return string(buf), nil
}