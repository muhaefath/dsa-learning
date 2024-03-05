package main

var mapPhoneNumber = map[string]string{
	"1": "1",
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
	"0": "0",
}

// 1. initiate resp []string
// 2. split phoneNumber into []rune
// 3. loop array
// 4. create function: phoneNumberMnemonics(target string, resp)[]string
// 4.1 if string == 1 or 0,
// 4.1.1 then loop resp and append target
// 4.2 else, covert target to 3 char
// 4.3 loop resp
// 4.4 append for every result char
// 5. return resp
func PhoneNumberMnemonics(phoneNumber string) []string {
	resp := []string{}

	for _, v := range phoneNumber {
		resp = phoneNumberMnemonics(string(v), resp)
	}

	return resp
}

func phoneNumberMnemonics(target string, resp []string) []string {
	phoneMap := mapPhoneNumber[target]

	tempResp := []string{}
	for _, phone := range phoneMap {
		if len(resp) == 0 {
			tempResp = append(tempResp, string(phone))
			continue
		}

		for _, value := range resp {
			value += string(phone)
			tempResp = append(tempResp, value)
		}

	}

	return tempResp
}
