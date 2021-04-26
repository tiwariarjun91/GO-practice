package main

import(
	"fmt"
	"encoding/xml"
)

type ChargingContractAllowanceFindRequest struct {
    XMLName 			xml.Name 	`xml:"chargingContractAllowanceFindRequest"`
    SubscriberAccountId string 		`xml:"subscriberAccountId"`
    ChargingContractId 	string 		`xml:"chargingContractId"`
    Ns2 				string 		`xml:"ns2,attr"`
}

func main(){
	str := "<chargingContractAllowanceFindRequest><subscriberAccountId>123456</subscriberAccountId><chargingContractId>78945</chargingContractId><ns2>45</ns2></chargingContractAllowanceFindRequest>"
    var r ChargingContractAllowanceFindRequest
    _ = xml.Unmarshal([]byte(str), &r)
    /*if err != nil {
		panic(err)
	}*/
 fmt.Println(r)
}