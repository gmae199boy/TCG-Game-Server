package connect

const (
	login         byte = 0x00 //처음 연결 했을 때, 유저 인증
	lobby         byte = 0x01 //유저 인증이 끝난 후, 타이틀 화면일때
	singlePlay    byte = 0x02 //싱글 플레이 시
	multiLobby    byte = 0x03 //멀티 플레이 요청시 멀티플레이 로비로 보낸다
	multiBattle   byte = 0x04 //멀티 플레이 성사시 배틀룸으로 옮긴다.
	multiCancle   byte = 0x05 //멀티 플레이 매칭을 취소 한다.
	battleConfirm byte = 0x06 //멀티 상대가 잡혔을 때 확인을 한다. 확인하면 배틀 성립
	inventory     byte = 0x07 //인벤토리 기능
)
const (
	battleContinue byte = 0x51
	battleGiveup   byte = 0x52
	deckListReq    byte = 0x53
	sendDeckList   byte = 0x54
	mouseCursor    byte = 0x55
)

const (
	//MaxUserNum : 동접 최대 유저 수
	MaxUserNum int = 500
	//LoginPacket : 처음 접속시 패킷의 용량 지정
	LoginPacket int = 16
	//MaxBattleUserNum : 배틀 로비의 수용 인원 수
	//MaxBattleUserNum int = 100
)
