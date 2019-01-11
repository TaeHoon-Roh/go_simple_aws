package aws_data_struct

import "time"

type PlayerEc2Status struct {
	TeamName       string
	Type           string
	Ec2Id          [3]string
	BlockDeviceId  [3]string
	KeyName        string
	KeyFingerPrint string
	KeyMaterial    string
	StartTime      time.Time
}

var Player = make([]PlayerEc2Status, 0)
