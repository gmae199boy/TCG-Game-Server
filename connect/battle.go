package connect

import (
	"fmt"
)

//BattleLobby : 유저가 멀티플레이를 할 때, 여기서 대기하면서 상대를 찾는다
type BattleLobby struct {
	/*
		지금은 랭킹이 없어서 있으면 바로바로 매칭이 되겠지만
		나중에 만약 랭킹제가 생긴다면 랭킹별로 매칭을 해야 하기 때문에
		꼭 로비 수용량을 늘리자
	*/
}

//InitBattleLobby : 배틀 로비 초기화
func (b *BattleLobby) InitBattleLobby() {

}

//MultiLobby : 멀티 플레이를 하기 위해 상대방을 기다림
func (b *BattleLobby) MultiLobby(connUser ConnUser) {
	mutex.Lock()
	if BLobbyUser.ConnInfo == nil {
		BLobbyUser = connUser

		fmt.Println(BLobbyUser.ConnInfo)
		fmt.Println("배틀 로비에 유저 정상적으로 등록!")
	} else {
		//매치 성사! 배틀 로비에서 배틀 룸으로 이동 >> 구현
		var sendMessage []byte = []byte{battleConfirm}
		_, err := BLobbyUser.ConnInfo.Write(sendMessage)
		_, err2 := connUser.ConnInfo.Write(sendMessage)
		if err != nil {
			fmt.Println("MultiLobby func : conn.Write ERR!!")
			return
		}
		if err2 != nil {
			fmt.Println("MultiLobby func : conn.Write ERR!!")
			return
		}
		fmt.Println("배틀이 정상적으로 성립됨")
		//배틀 상대를 자기 정보에 입력한다.
		var opponentNum int = BLobbyUser.UserArrNum
		ConnUserArr[connUser.UserArrNum].BattleOpponent = opponentNum
		ConnUserArr[BLobbyUser.UserArrNum].BattleOpponent = connUser.UserArrNum
		b.DeleteBLobby(BLobbyUser.UserArrNum)
	}
	mutex.Unlock()
}

//DeleteBLobby : 배틀이 끝났거나 강제종료, 오류가 났을 떄 유저를 로비에서 제거.
func (b *BattleLobby) DeleteBLobby(userArrNum int) {
	fmt.Println("start BLobby user delete : in DeleteBLobby func")
	fmt.Println(BLobbyUser.ConnInfo)
	BLobbyUser.UserName = ""
	BLobbyUser.UserArrNum = 0
	BLobbyUser.ConnInfo = nil
	BLobbyUser.BattleOpponent = -1
	fmt.Println("delete complete!!")
}

//Battle : 배틀이 시작되면 여기에서 로직 처리
//로그 남기는 로직도 작성할 것!!!!!!!!!!!!
func (b *BattleLobby) Battle(me *ConnUser, opponent *ConnUser) {
	recvBuf := make([]byte, 32)
	fmt.Println("배틀에 입장하였습니다.")
	deckReq := []byte{deckListReq}
	//덱 리스트 요청
	me.ConnInfo.Write(deckReq)
	//배틀 메인 루프
	for {
		_, err := me.ConnInfo.Read(recvBuf)
		if err != nil {
			fmt.Println("배틀 중에 패킷 수신 오류 : ", err)
			return
		}
		switch recvBuf[0] {
		case battleContinue:
			{
				_, err = opponent.ConnInfo.Write(recvBuf)
				if err != nil {
					fmt.Println("배틀 중에 패킷 전송 오류 : ", err)
					return
				}
				fmt.Println("배틀을 속행한다!!")
			}
			break
		case battleGiveup:
			{

			}
			break
		case sendDeckList:
			{
				fmt.Println(recvBuf)
				//나중에 덱정보 교환 알고리즘 절대 변경할 것!!!!!!!!!!!!!!!!!!!!!!!
				//지금은 프로토타입이라 이렇게 교환하지만 나중엔 인벤토리에서 캐싱해서 로딩할 것!!
				// deckList := make([]byte, 2)
				// var deckArr [5]uint16
				// i := 1
				// ind := 0
				// for i < 31 {
				// 	deckList[0] = recvBuf[i]
				// 	deckList[1] = recvBuf[i+1]
				// 	deckArr[ind] = binary.LittleEndian.Uint16(deckList)
				// 	bytesArr := byte(ind)
				// 	me.UserData.Inventory.Deck[0].BoardArr[ind] = bytesArr
				// 	ind++
				// 	i += 2
				// }
				// me.UserData.Inventory.Deck[0].DeckFrontNum = recvBuf[31]
				recvBuf[0] = sendDeckList
				_, err = opponent.ConnInfo.Write(recvBuf)
				if err != nil {
					fmt.Println("덱교환 중에 패킷 전송 오류 : ", err)
					return
				}
				fmt.Println("덱정보를 교환 했다!")
			}
			break
		case mouseCursor:
			{
				recvBuf[0] = mouseCursor
				var sendCursor []byte = []byte{recvBuf[0], recvBuf[1], recvBuf[2]}
				_, err = opponent.ConnInfo.Write(sendCursor)
				if err != nil {
					fmt.Println("커서 위치 전송 오류 : ", err)
					return
				}
			}
			break
		default:
			{

			}
		}
	}
}
