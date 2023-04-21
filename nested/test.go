package nested

import (
	"github.com/sirupsen/logrus"
)

func GetNestedMessage() string {
	logrus.Info("lol")
	return "nested message go_mod_3"
}
