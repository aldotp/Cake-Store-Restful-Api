package test

import (
	"github.com/aldotp/cakestore/config"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("***************** Testing *****************")
	config.LoadEnv("../.env.test")
	config.ConnectToDatabase()
	m.Run()
	log.Println("**************** Test Ended ***************")
}
