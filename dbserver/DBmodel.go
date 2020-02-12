package dbserver

//직업별 상수 선언
const (
	//Warrior : 전사
	Warrior byte = 0x01
	//Rogue : 도적
	Rogue byte = 0x0B
	//Magician : 마법사
	Magician byte = 0x15
	//Prist : 사제
	Prist byte = 0x20
)

/*
	비트에 따른 유형 ***
	64비트 기준

	0x0000000000000000
	  1111				앞 4자리 공격 유형(타입과 같다)
		  1111			다음 4자리 방어 유형
			  1111		버프 유형
					1111  마지막 4자리 디버프 유형

	랜덤 공격은 같은 대상이 여러번 맞을 수 있다.

	도발은 전체공격이 5명이 맞을 수 있다면 5명 분을 전부 다 맞는다.
*/

//공격 유형 여기에 // 공격,방어,버프,디버프 모두 포함 지정 타입.
const (
	Atk     uint64 = 0x0001000000000000 //일반 공격
	LineAtk uint64 = 0x0002000000000000 //열 공격
	AllAtk  uint64 = 0x0004000000000000 //전체 공격
	Ran2Atk uint64 = 0x0008000000000000 //랜덤 2번 공격
	Ran3Atk uint64 = 0x0010000000000000 //랜덤 3번 공격
	Ran4Atk uint64 = 0x0020000000000000 //랜덤 4번 공격
	Ran5Atk uint64 = 0x0040000000000000 //랜덤 5번 공격
	Enemy   uint64 = 0x1000000000000000
	My      uint64 = 0x2000000000000000
)

//방어 유형 여기에
const (
	Def   uint64 = 0x0000000100000000 //다음 공격을 * 감소 시킨다.
	Taunt uint64 = 0x0000000200000000 //도발
)

//버프 인덱스 여기에
const (
	HollyShield uint64 = 0x0000000000010000 //보호막
	TickHeal    uint64 = 0x0000000000020000 //지속 치유
	Pray        uint64 = 0x0000000000040000 //기도(정화)
)

//디버프 인덱스 여기에
const (
	Poison   uint64 = 0x0000000000000001 //독뎀
	Frost    uint64 = 0x0000000000000002 //빙결
	Stun     uint64 = 0x0000000000000004 //스턴
	Silence  uint64 = 0x0000000000000008 //침묵
	Weakness uint64 = 0x0000000000000010 //약화
	Curse    uint64 = 0x0000000000000020 //저주
)

//카드 타입
const (
	ChaWarrior byte = 1
	ChaRogue   byte = 3
	Inher      byte = 2
	Skill      byte = 0x03
	Equip      byte = 0x04
)

//고유 스킬 여기에
const ()

//장착 스킬 여기에
const ()

//싱글 플레이 챕터 진행 상황 여기에
const ()

//CardInfo : 카드 베이스 정보
type CardInfo struct {
	ID     uint16
	Image  string
	Name   string
	Desc   string
	Job    byte
	EQJob  byte
	Atk    byte
	Hp     byte
	Type   uint64
	Delay  byte
	EImage string
}

//Inventory : 유저 인벤토리
type Inventory struct {
	CardList []CardInfo
	Deck     []DeckInfo //구매목록 확인 대상
}

//DeckInfo : 인벤토리에 들어갈 캐릭터 카드 정보(스킬, 장비 정보 포함)
type DeckInfo struct {
	BoardArr     []byte
	DeckFrontNum byte
}

//UserData : 유저 정보
type UserData struct {
	Email        string
	ID           string
	Inventory    Inventory
	DeckPurchase byte //덱 리스트 최대 크기 증가 구입 갯수
}
