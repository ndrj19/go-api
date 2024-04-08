package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"gwaggli": {
		AuthToken: "123ABC",
		Username:  "gwaggli",
	},
	"urs": {
		AuthToken: "456DEF",
		Username:  "urs",
	},
	"gigi": {
		AuthToken: "789GHI",
		Username:  "gigi",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"gwaggli": {
		Coins:    100,
		Username: "gwaggli",
	},
	"urs": {
		Coins:    200,
		Username: "urs",
	},
	"gigi": {
		Coins:    300,
		Username: "gigi",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate db call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate db call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
