package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	rootLogsDirpath = "/oasa-telematics"
)

var Logger *logrus.Logger

func InitLogger(applicationName string) {
	Logger = logrus.New()
	directoryPath := path.Join(rootLogsDirpath, applicationName)
	err := os.Mkdir(directoryPath, 0777)
	if err != nil {
		fmt.Printf("error create directory file: %v\n", err)
	}
	fileName := path.Join(rootLogsDirpath, applicationName, "opswlog.log")
	// open a file
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
	}
	Logger.SetOutput(f)
}

func INFO(str string) {
	Logger.Println(str)
}
