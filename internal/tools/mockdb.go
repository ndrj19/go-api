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

var mockBitcoinDetails = map[string]BitcoinDetails{
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

var mockEthereumDetails = map[string]EthereumDetails{
	"gwaggli": {
		Coins:    1000,
		Username: "gwaggli",
	},
	"urs": {
		Coins:    2000,
		Username: "urs",
	},
	"gigi": {
		Coins:    3000,
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

func (d *mockDB) GetUserBitcoin(username string) *BitcoinDetails {
	// Simulate db call
	time.Sleep(time.Second * 1)

	var clientData = BitcoinDetails{}
	clientData, ok := mockBitcoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserEthereum(username string) *EthereumDetails {
	// Simulate db call
	time.Sleep(time.Second * 1)

	var clientData = EthereumDetails{}
	clientData, ok := mockEthereumDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
