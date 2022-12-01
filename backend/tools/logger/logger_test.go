package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	log := NewFileLogger("test")

	log.WithField("contractID", 3378).Infoln("InfoLn")
	log.WithField("oper_login", "m.darkzar@gmail.com").Debugln("Debugln")
	log.WithFields(logrus.Fields{"contractID": 3378, "se_id": 200189621, "oper_login": "m.darkzar@gmail.com", "details": "test"}).Errorln("Errorln")

}
