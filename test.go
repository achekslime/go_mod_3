package go_mod_3

import "github.com/sirupsen/logrus"

func GetMessage() string {
	logrus.Info("lol")
	return "go_mod_3"
}
