package tests

import (
	"testing"

	"github.com/goshx/leanote/app/db"
	//	. "github.com/goshx/leanote/app/lea"
	//	"github.com/goshx/leanote/app/service"
	//	"gopkg.in/mgo.v2"
	//	"fmt"
)

func TestDBConnect(t *testing.T) {
	db.Init("mongodb://localhost:27017/leanote", "leanote")
}
