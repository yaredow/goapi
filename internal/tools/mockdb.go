package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"Yared": {
		AuthToken: "1234567890",
		Username:  "yaredow",
	},
	"John": {
		AuthToken: "0987654321",
		Username:  "johnsmith",
	},
	"Jane": {
		AuthToken: "1122334455",
		Username:  "janesmith",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"Yared": {
		Coins:    100,
		username: "yaredow",
	},
	"John": {
		Coins:    50,
		username: "johnsmith",
	},
	"Jane": {
		Coins:    75,
		username: "janesmith",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)
	userData := LoginDetails{}
	userData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &userData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)
	userCoinData := CoinDetails{}
	userCoinData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &userCoinData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
