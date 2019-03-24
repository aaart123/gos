/*
 * Generated by tools/gen_protocol
 */

package pt

var NameToId = map[string]uint16{
	"SessionAuthParams":   uint16(1),
	"SessionAuthResponse": uint16(2),
	"Ok":                  uint16(3),
	"Fail":                uint16(4),
	"EquipLoadParams":     uint16(5),
	"EquipLoadResponse":   uint16(6),
	"EquipUnLoadParams":   uint16(7),
	"EquipUnLoadResponse": uint16(8),
	"LoginResponse":       uint16(9),
	"RoomJoinParams":      uint16(10),
	"RoomJoinResponse":    uint16(11),
	"RoomJoinNotice":      uint16(12),
}

var IdToName = map[uint16]string{
	uint16(1):  "SessionAuthParams",
	uint16(2):  "SessionAuthResponse",
	uint16(3):  "Ok",
	uint16(4):  "Fail",
	uint16(5):  "EquipLoadParams",
	uint16(6):  "EquipLoadResponse",
	uint16(7):  "EquipUnLoadParams",
	uint16(8):  "EquipUnLoadResponse",
	uint16(9):  "LoginResponse",
	uint16(10): "RoomJoinParams",
	uint16(11): "RoomJoinResponse",
	uint16(12): "RoomJoinNotice",
}

var IdToType = map[uint16]int{
	uint16(1):  PT_TYPE_GS,
	uint16(2):  PT_TYPE_GS,
	uint16(3):  PT_TYPE_GS,
	uint16(4):  PT_TYPE_GS,
	uint16(5):  PT_TYPE_GS,
	uint16(6):  PT_TYPE_GS,
	uint16(7):  PT_TYPE_GS,
	uint16(8):  PT_TYPE_GS,
	uint16(9):  PT_TYPE_GS,
	uint16(10): PT_TYPE_ROOM,
	uint16(11): PT_TYPE_ROOM,
	uint16(12): PT_TYPE_ROOM,
}

const (
	PT_SessionAuthParams   = "SessionAuthParams"
	PT_SessionAuthResponse = "SessionAuthResponse"
	PT_Ok                  = "Ok"
	PT_Fail                = "Fail"
	PT_EquipLoadParams     = "EquipLoadParams"
	PT_EquipLoadResponse   = "EquipLoadResponse"
	PT_EquipUnLoadParams   = "EquipUnLoadParams"
	PT_EquipUnLoadResponse = "EquipUnLoadResponse"
	PT_LoginResponse       = "LoginResponse"
	PT_RoomJoinParams      = "RoomJoinParams"
	PT_RoomJoinResponse    = "RoomJoinResponse"
	PT_RoomJoinNotice      = "RoomJoinNotice"
)

const (
	PT_TYPE_GS   = 1
	PT_TYPE_ROOM = 2
)