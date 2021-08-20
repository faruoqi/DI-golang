package services

import (
	"DI/ds"
	"DI/utils"
	"errors"
)

type SimpleLogic struct {
	l  utils.Logger
	ds ds.DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodBye(userID string) (string, error) {
	sl.l.Log("in SayGoodBye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}

	return "Goodbye, " + name, nil

}

func NewSimpleLogic(l utils.Logger, ds ds.DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}
