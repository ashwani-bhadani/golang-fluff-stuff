package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to processing jsons")
	//encoding struct to JSON
	// EncodeJson()
	DecodeJson()
}

type course struct {
	Name     string   `json:"course name"` //this way to correctly define a json payload, called aliases
	Price    float32  `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`               //omit from JSON like @JsonIgnore
	Tags     []string ` json:"tags,omitempty"` //omit if empty/nil
	//be very mindful of this space here between tags & omitempty. will give error "tags, omitempty"
}

func EncodeJson() {
	lcocourses := []course{ //slice of stuct course type
		//slice of courses
		{"react js bootcamp", 299.0, "learncodeonline.com", "abc123", []string{"web-dev", "js"}},
		{"angular bootcamp", 499.0, "learncodeonline.com", "aaa23", []string{"web-dev", "js", "angular"}},
		{"mern bootcamp", 499.0, "learncodeonline.com", "aaa23", []string{"web-dev", "js", "express.js"}},
	}
	//marshall is impl of using jsons, another option is to beautify json too
	// resultJson, err := json.Marshal(lcocourses)
	// resultJson, err := json.MarshalIndent(lcocourses, "", " ") //prefix ="" , indent = space
	resultJson, err := json.MarshalIndent(lcocourses, "", "\t") //prefix ="" , indent = tab
	//MarshalIndent takes 2 params, the any-interface & 2nd based on what to indent values
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", resultJson)
}

func DecodeJson() {
	jsonDataRecieved := []byte(`{
                "course name": "mern bootcamp",
                "price": 499,
                "platform": "learncodeonline.com",
                "tags": [
                        "web-dev",
                        "js",
                        "express.js"
                ]
        }`) //any data from web is byte stream, convert string to byte here for example

	var lcoCourse course
	isValidJson := json.Valid(jsonDataRecieved)
	if isValidJson {
		fmt.Println("JSON is valid") //note password was not printed
		//we do not want to pass just an interface but actually store in lcoCourse, so pass a reference
		json.Unmarshal(jsonDataRecieved, &lcoCourse) //want to make sure copy is not passed along
		fmt.Printf("%#v\n", lcoCourse)               //#=interface %#v prints interface value
	} else {
		println("Not a valid JSON")
	}

	//-----------------some case you just want to add data to K-V pair--------------------
	//we do not know data format coming from web so declare it as an interface
	//map for json, the 1st value which is key will always be string but value of it can be of ANY type, so use interface
	var onlineDataFetch map[string]interface{}
	json.Unmarshal(jsonDataRecieved, &onlineDataFetch) //this map is useful to process now
	// fmt.Printf("%#v\n", onlineDataFetch)

	for k, v := range onlineDataFetch {
		fmt.Printf("key is %v and value is %v & type is %T\n", k, v, v)
		/*		OUTPUT
				key is course name and value is mern bootcamp & type is string
				key is price and value is 499 & type is float64
				key is platform and value is learncodeonline.com & type is string
				key is tags and value is [web-dev js express.js] & type is []interface {}
		*/
	}

}
