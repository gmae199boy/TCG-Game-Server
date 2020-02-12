package connect

import (
	"fmt"
	dbserver "github.com/gmae199boy/tcgServer/DBServer"
	"io"
	"log"
	"net"
	"sync"
)

//온라인 유저를 담을 전역 변수들
var (
	//ConnUserArr : 유저 최대 수용 배열 (500명)
	ConnUserArr []ConnUser = make([]ConnUser, MaxUserNum)
	//BLobbyUser :  멀티 플레이를 신청한 유저
	BLobbyUser ConnUser
)

//ConnUser : 접속하는 유저들 정보 저장용
type ConnUser struct {
	ConnInfo       net.Conn
	UserData       dbserver.UserData
	UserCurState   byte
	UserName       string
	UserArrNum     int
	BattleOpponent int
	//나중에 덱정보나 그런 유저 정보까지 풀링해서 집어 넣어야 한다.
}

//읽,쓰기 잠금 변수
var mutex *sync.RWMutex = new(sync.RWMutex)

//Connection : 서버 구동할 객체
//var ConnUserArr = 서버가 수용하는 전체 유저들 정보 배열
type Connection struct {
}

//InitConn : 객체를 만들고 나서 객체 초기화용 함수
func (c *Connection) InitConn() {

}

//ServerLobby : 처음 패킷에 따라 연결을 달리 함
func (c *Connection) ServerLobby(conn net.Conn, b BattleLobby) {
	var arrNum int
	for {
		recvBuf := make([]byte, 16)
		_, err := conn.Read(recvBuf)
		if nil != err {
			if io.EOF == err {
				log.Println(err)
				return
			}
			log.Println(err)
			return
		}
		fmt.Println(recvBuf)
		switch recvBuf[0] {
		case login:
			{

			}
			break
		case lobby:
			{
				arrNum = c.AddUserArr(conn)
				defer c.DeleteUserArr(arrNum)
			}
			break
		case singlePlay:
			{

			}
			break
		case multiLobby:
			{
				b.MultiLobby(ConnUserArr[arrNum])
			}
			break
		case multiBattle:
			{

			}
			break
		case multiCancle:
			{

			}
			break
		case battleConfirm:
			{
				b.Battle(&ConnUserArr[arrNum], &ConnUserArr[ConnUserArr[arrNum].BattleOpponent])
			}
			break
		default:
			{
				panic("recvBuf[0]에서 잘못된 인자 전달. 패닉!")
			}
		}
	}
}

//DeleteUserArr : 접속을 종료하거나 종료 되면 실행. 유저 배열에서 유저를 삭제한다
func (c *Connection) DeleteUserArr(userArrNum int) {
	mutex.Lock()
	fmt.Println("start disconnect user : in DeleteUserArr func")
	fmt.Println(ConnUserArr[userArrNum].ConnInfo)
	ConnUserArr[userArrNum].UserName = ""
	ConnUserArr[userArrNum].UserArrNum = 0
	ConnUserArr[userArrNum].BattleOpponent = -1
	ConnUserArr[userArrNum].ConnInfo.Close()
	ConnUserArr[userArrNum].ConnInfo = nil
	fmt.Println("disconnect complete!!")
	fmt.Println(ConnUserArr[userArrNum].ConnInfo)
	mutex.Unlock()
}

//AddUserArr : 처음 접속 했을 때 유저 배열에 유저 등록
func (c *Connection) AddUserArr(conn net.Conn) int {
	//접속자가 500명이 안되었을 경우 (500명 동접 제한 > 나중에 수정함)
	for i := 0; i < MaxUserNum; i++ {
		/*
			쓰기 락이 애매할 수 있다. 나중에 문제가 생기면 수정 필!
		*/
		mutex.Lock() //여기를
		if ConnUserArr[i].ConnInfo == nil {
			var user ConnUser
			user.ConnInfo = conn
			user.UserArrNum = i
			user.UserCurState = lobby
			user.BattleOpponent = -1
			user.UserData = dbserver.UserData{}
			user.UserData.Inventory.Deck = make([]dbserver.DeckInfo, 5)
			user.UserData.Inventory.Deck[0].BoardArr = make([]byte, 5)
			//나중에 더 추가!!
			ConnUserArr[i] = user
			sendSocket := make([]byte, 8)

			sendSocket[0] = 10
			_, err := conn.Write(sendSocket)
			if err != nil {
				fmt.Println(err)
				return -1
			}
			fmt.Println("유저를 서버에 정상적으로 등록")
			mutex.Unlock()
			return i
		}
		mutex.Unlock()
	}
	panic("AddUserArr : 범위를 벗어남!")
}

//ConInit : 처음 접속자를 소켓에 연결
// func ConInit(conn net.Conn) {
// 	recvBuf := make([]byte, 16)
// 	for {
// 		n, err := conn.Read(recvBuf)
// 		//buf := bytes.NewBuffer(recvBuf)
// 		//d := json.NewDecoder(conn)
// 		//var msg Coor
// 		//err := d.Decode(&msg)
// 		fmt.Println(recvBuf)
// 		one := binary.LittleEndian.Uint16(recvBuf[2:4])
// 		fmt.Println(one)
// 		//fmt.Println(two)
// 		if nil != err {
// 			if io.EOF == err {
// 				log.Println(err)
// 				return
// 			}
// 			log.Println(err)
// 			return
// 		}
// 		if n > 0 {
// 			data := new(bytes.Buffer)
// 			_ = binary.Write(data, binary.LittleEndian, uint8(18))
// 			_ = binary.Write(data, binary.LittleEndian, uint16(50098))
// 			//var data []byte = []byte("123sdvsd")
// 			//var mss []byte
// 			//mss, err := json.Marshal(msg)
// 			//log.Println(mss)
// 			//fmt.Println(data)
// 			_, err = conn.Write(data.Bytes())
// 			if err != nil {
// 				log.Println(err)
// 				return
// 			}
// 		}
// 	}
// }
