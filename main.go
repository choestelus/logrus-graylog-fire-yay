package main

import (
	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	hook := graylog.NewGraylogHook("127.0.0.1:12201", map[string]interface{}{"environment": "yay"})
	defer hook.Flush()

	logrus.AddHook(hook)
	logrus.Infof("yayyy")
}
