package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonData := []byte(`
	{
	  "username" :"blackhat",
	  "shipto":
	    {
	        "street": "Sulphur Springs Rd",
	        "city": "Park City",
	        "state": "VA",
	        "zipcode": 12345
	    },
	  "order":
	    {
	      "paid":false,
	      "orderdetail" :
	        [{
	          "itemname":"A Guide to the World of zeros and ones",
	          "desc": "book",
	          "qty": 3,
	          "price": 50
	        },
			{
	          "itemname":"Second item",
	          "desc": "player",
	          "qty": 2,
	          "price": 90
	        },
			{
	          "itemname":"TV",
	          "desc": "TV sony",
	          "qty": 1,
	          "price": 99
	        }
			]

	    }
	}
	`)
	if !json.Valid(jsonData) {
		fmt.Printf("JSON is not valid: %s", jsonData)
	}
	var c customer
	err := json.Unmarshal(jsonData, &c)
	if err != nil {
		fmt.Println(err)
	}
//	fmt.Println(c)
	c2, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(c2))
}

type address struct {
	Street		string
	City		string
	Zipcode		int
}

type item struct  {
	Name		string `json:"itemname"`
	Description	string `json:"desc,omitempty"`
	Quantity	int `json:"qty"`
	Price		int `json:"price"`
}

type order struct {
	TotalPrice	int `json:"totalprice"`
	IsPaid		bool `json:"paid"`
	Fragile		bool `json:"fragile,omitempty"`
	OrderDetail []item `json:"orderdetail"`
}


type customer struct {
	UserName		string  `json:"username"`
	Password		string  `json:"-"`
	Token			string  `json:"token"`
	ShipTo			address `json:"shipto"`
	PurchaseOrder	order   `json:"order"`
}



