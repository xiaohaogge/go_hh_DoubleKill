package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type WxUser struct{
	Name 	string 		`boson:"name"`
	Password 	string 	`boson:"password"`
}

func InitMongo(){
	mongo, err := mgo.Dial("127.0,0,1")
	defer mongo.Close()
	if err != nil {
		fmt.Println(err)
	}
	client := mongo.DB("hh").C("user")

	// 创建数据
	data := WxUser{
		Name:"xhaoge",
		Password:"123456",
	}
	// 插入数据
	cErr := client.Insert(&data)
	if cErr != nil {
		fmt.Println(cErr)
	}
}

// import "honnef.co/go/tools/structlayout"

// import "golang.org/x/tools/go/ssa/interp"

type RoomBaseResponse struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg`
	Data map[string]interface{} `json:data,omitempty`
}

type UserLoginAdd struct {
	Code string `json:"code`
}

type NewRoomAddReq struct {
	Title        string   `json:title`
	PicIdList    []string `json:picIdList`
	Position     string   `json:position`
	Address      string   `json:address`
	RoomType     []int    `json:roomType`
	IsElevator   int      `json:isElevator`
	Price        int      `json:price`
	NearSubway   string   `json:nearSubway`
	Area         float64  `josn:area`
	Floor        int      `json:floor`
	Plot         string   `json:plot`
	Supporting   []string `json:supporting`
	Description  string   `json:description`
	ReleaseTime  string   `json:releaseTime`
	PayType      string   `json:payType`
	ContactPhone string   `json:contactPhone`
	ContactWx    string   `json:contactWx`
	CreatorId    int      `json:creatorId`
}

type NewRoom struct {
	TraceId      string
	RoomInfo     *RoomInfo
	PayType      string
	ContactPhone string
	ContactWx    string
	CreatorId    int
}

type RoomInfo struct {
	Title       string
	RoomType    string
	PicList     string
	Position    string
	Address     string
	IsElevator  bool
	Price       int
	NearSubway  string
	Area        string
	Floor       int
	Plot        string
	Supporting  []string
	Description string
	ReleaseTime string
}

type UserInfo struct {
	Id          string   `json:id`
	OpenId      string   `json:openId`
	Username    string   `json:username`
	Password    string   `json:password`
	Phone       string   `json:phone`
	IsAdmin     string   `json:isAdmin`
	ContainId   []string `json:containId`
	Description string   `json:description`
}

type UserUpdate struct {
	Id          string `json:id`
	Username    string `json:username`
	Password    string `json:password`
	Phone       string `json:phone`
	IsAdmin     string `json:isAdmin`
	Description string `json:description`
}
