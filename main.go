package main

import (
	"encoding/json"
	"fmt"

	"net/http"
)

type User struct{
	Cardno int `json:"c_n"`
}

var user User
var err error

func module(card_no int) (int, int) {

	var sli []int
	var x int
	len := 0

	for card_no > 0 {

		x = card_no % 10
		sli = append(sli, x)
		card_no = card_no / 10
		len++

	}

	sum := 0
	j := false
	for i := 0; i < len; i++ {

		if j {
			sli[i] = 2 * sli[i]
			if sli[i] > 9 {
				x = sli[i] % 10
				sli[i] = sli[i] / 10
				sum = sum + x + sli[i]
			} else {
				sum = sum + sli[i]
			}

			// x=sli[i]%10
			// sli[i]=sli[i]/10
			// sum=sum+x+sli[i]
			j = false

		} else {
			sum = sum + sli[i]
			j = true

		}

	}
	return sum, len

}

func isvalidcard(w http.ResponseWriter, r *http.Request) {

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error in decoding data")
	}

	sum, len := module(user.Cardno)

	if len == 16 || len == 19 {
		if sum%10 == 0 {
			json.NewEncoder(w).Encode("valid card no")
		} else {
			json.NewEncoder(w).Encode("not valid card number")
		}
	} else {
		json.NewEncoder(w).Encode("please enter valid card number")
	}

}

func checksumdigit(w http.ResponseWriter, r *http.Request) {

	json.NewDecoder(r.Body).Decode(&user)
	var req_no int
	_, len := module(user.Cardno)
	if len == 15 {
		user.Cardno = user.Cardno * 10
		sum, _ := module(user.Cardno)
		y := sum % 10
		if y != 0 {
			req_no = 10 - y
		} else {
			req_no = 0
		}
		user.Cardno = user.Cardno + req_no

		// type user1 struct{
		// 	digit_to_append int `json:"d_t_a"`
		// 	resultant_card_no int `json:"r_c_no"`
		// }
		// ap := user1{
		// 	digit_to_append: req_no,
		// 	resultant_card_no: user.card_no + req_no,
		// }
		json.NewEncoder(w).Encode(req_no)
	}else{
		json.NewEncoder(w).Encode("card no is not of 15 digits")
	}


}

func main() {
	fmt.Println("we are good to go")

	http.HandleFunc("/isvalidcard", isvalidcard)
	http.HandleFunc("/checksumdigit", checksumdigit)

	http.ListenAndServe(":8080", nil)
}
