package sony

import (
	"github.com/sony/sonyflake"
)

func NewIDGenerator(machineID uint16) *sonyflake.Sonyflake {
	settings := sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return machineID, nil
		},
	}

	return sonyflake.NewSonyflake(settings)
}
