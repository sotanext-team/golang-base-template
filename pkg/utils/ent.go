package utils

import (
	"reflect"

	"app-api/ent"

	"github.com/sirupsen/logrus"
)

type EntClientWrapper struct {
	Client *ent.Client
	Tx     *ent.Tx
}

func ParseEntClient(client interface{}) EntClientWrapper {
	varType := reflect.TypeOf(client).String()

	logrus.Info(varType)
	switch varType {
	case "*ent.Client":
		return EntClientWrapper{
			Client: client.(*ent.Client),
			Tx:     nil,
		}
	case "*ent.Tx":
		return EntClientWrapper{
			Client: nil,
			Tx:     client.(*ent.Tx),
		}
	default:
		return EntClientWrapper{
			Client: client.(*ent.Client),
			Tx:     nil,
		}
	}
}
