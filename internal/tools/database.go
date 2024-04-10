package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database collections
type LoginDetails struct {
	AuthToken string
	Username  string
}

type BitcoinDetails struct {
	Coins    int64
	Username string
}
type EthereumDetails struct {
	Coins    int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserBitcoin(username string) *BitcoinDetails
	GetUserEthereum(username string) *EthereumDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {

	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
