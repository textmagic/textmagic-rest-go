package textmagic

import (
	"encoding/json"
	"fmt"
)

const (
	STAT_RES       = "stats"
	TOKEN_RES      = "tokens"
	USER_RES       = "user"
	SUBACCOUNT_RES = "subaccounts"
)

type MessagingStat struct {
	ReplyRate             float32 `json:"replyRate"`
	Date                  string  `json:"date"`
	DeliveryRate          float32 `json:"deliveryRate"`
	Costs                 float32 `json:"costs"`
	MessagesReceived      uint32  `json:"messagesReceived"`
	MessagesSentDelivered uint32  `json:"messagesSentDelivered"`
	MessagesSentAccepted  uint32  `json:"messagesSentAccepted"`
	MessagesSentBuffered  uint32  `json:"messagesSentBuffered"`
	MessagesSentFailed    uint32  `json:"messagesSentFailed"`
	MessagesSentRejected  uint32  `json:"messagesSentRejected"`
	MessagesSentParts     uint32  `json:"messagesSentParts"`
}

type MessagingStatList []MessagingStat

type SpendingStat struct {
	Id      uint32  `json:"id"`
	UserId  uint32  `json:"userId"`
	Date    string  `json:"date"`
	Balance float32 `json:"balance"`
	Delta   float32 `json:"delta"`
	Type    string  `json:"type"`
	Value   string  `json:"value"`
	Comment string  `json:"comment"`
}

type SpendingStatList struct {
	Page          uint32         `json:"page"`
	Limit         uint8          `json:"limit"`
	PageCount     uint32         `json:"pageCount"`
	SpendingStats []SpendingStat `json:"resources"`
}

type TimezoneList map[string]string

type NewToken struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Expires  string `json:"string"`
}

type Currency struct {
	Id         string `json:"id"`
	HtmlSymbol string `json:"htmlSymbol"`
}

type Timezone struct {
	Id       uint32 `json:"id"`
	Area     string `json:"area"`
	Dst      uint8  `json:"dst"`
	Offset   int    `json:"offset"`
	Timezone string `json:"timezone"`
}

type User struct {
	Id             uint32   `json:"id"`
	Username       string   `json:"username"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Status         string   `json:"status"`
	Balance        float32  `json:"balance"`
	Company        string   `json:"company"`
	Currency       Currency `json:"currency"`
	Timezone       Timezone `json:"timezone"`
	SubaccountType string   `json:"subaccountType"`
}

type UserList struct {
	Page      uint32 `json:"page"`
	Limit     uint8  `json:"limit"`
	PageCount uint32 `json:"pageCount"`
	Users     []User `json:"resources"`
}

/*
Return messaging statistics.

	Parameters:

		Var `data` may contain next keys:

			by:    Group results by specified period: `off`, `day`, `month` or `year`. Default is `off`.
			start: Start date in unix timestamp format. Default is 7 days ago.
			end:   End date in unix timestamp format. Default is now.
*/
func (client *TextmagicRestClient) GetMessagingStat(data map[string]string) ([]MessagingStat, error) {
	stat := make([]MessagingStat, 0)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/messaging", client.BaseUrl(), STAT_RES)
	params := transformGetParams(data)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return stat, err
	}
	err = json.Unmarshal(response, &stat)

	return stat, err
}

/*
Return account spending statistics.

	Parameters:

		Var `data` may contain next keys:

			page:  Fetch specified results page. Default=1
	        limit: How many results on page. Default=10
	        start: Start date in unix timestamp format. Default is 7 days ago.
	        end:   End date in unix timestamp format. Default is now.
*/
func (client *TextmagicRestClient) GetSpendingStat(data map[string]string) (*SpendingStatList, error) {
	stat := new(SpendingStatList)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/spending", client.BaseUrl(), STAT_RES)

	params := transformGetParams(data)
	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return stat, err
	}

	err = json.Unmarshal(response, stat)

	return stat, err
}

/*
Get current user info.
*/
func (client *TextmagicRestClient) GetUser() (*User, error) {
	user := new(User)

	method := "GET"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), USER_RES)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(response, user)

	return user, err
}

/*
Update an current User via a PUT request.

	Parameters:

		Var `data` may contain next keys:

			firstName: User first name. Required.
			lastName:  User last name. Required.
			company:   User company. Required.
*/
func (client *TextmagicRestClient) UpdateUser(data map[string]string) (map[string]string, error) {
	result := make(map[string]string)

	method := "PUT"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), USER_RES)
	params := preparePostParams(data)

	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(response, &result)

	return result, err
}

/*
Get all subaccounts.

	Parameters:

		id: Subaccount id.
*/
func (client *TextmagicRestClient) GetSubaccount(id uint32) (*User, error) {
	var subaccount = new(User)

	method := "GET"

	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), SUBACCOUNT_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return subaccount, err
	}

	err = json.Unmarshal(response, subaccount)

	return subaccount, err
}

/*
Get all user subaccounts.

	Parameters:

		Var `data` may contain next keys:

			page:  Fetch specified results page.
        	limit: How many results on page.
*/
func (client *TextmagicRestClient) GetSubaccountList(data map[string]string) (*UserList, error) {
	var subaccountList = new(UserList)

	method := "GET"
	params := transformGetParams(data)

	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), SUBACCOUNT_RES)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return subaccountList, err
	}

	err = json.Unmarshal(response, subaccountList)

	return subaccountList, err
}

/*
Invite new subaccount.

	Parameters:

		Var `data` may contain next keys:

			email: Subaccount email. Required.
        	role:  Subaccount role: `A` for administrator or `U` for regular user. Required.
*/
func (client *TextmagicRestClient) SendInvite(data map[string]string) (bool, error) {
	var success bool

	method := "POST"

	params := preparePostParams(data)

	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), SUBACCOUNT_RES)

	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return false, err
	}

	if response[0] == 204 {
		success = true
	}

	return success, err
}

/*
Close subaccount.

	Parameters:

		id: The unique id of the Subaccount to close.
*/
func (client *TextmagicRestClient) CloseSubaccount(id uint32) (bool, error) {
	var success bool

	method := "DELETE"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), SUBACCOUNT_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return false, err
	}
	if response[0] == 204 {
		success = true
	}

	return success, err
}
