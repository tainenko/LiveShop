package tests

import (
	"github.com/LiveShop/utils"
	"testing"
)
var myEnv = utils.LoadConfig()

func TestLoadConfig(t *testing.T) {
	if myEnv["DB_USER"] != "postgresql"{
		t.Error(myEnv["DB_USER"])
	}else{
		t.Log("Test LoadConfig Passed!")
	}
}

func Test_Key_not_exist(t *testing.T){
	if myEnv["Account"] !=""{
		t.Error("'Account' should be empty.")
	}else{
		t.Log("Test Key_not_exist Passed!")
	}
}

