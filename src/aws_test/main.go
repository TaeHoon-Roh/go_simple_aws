package main

import (
	"aws_test/aws_data_struct"
	"fmt"
	"strconv"
	"time"
)

func makeTestPlayer(pnum int) {
	for i := 0; i < pnum; i++ {
		p := aws_data_struct.PlayerEc2Status{
			TeamName:       "uxfac" + strconv.Itoa(i),
			Type:           "Rule" + strconv.Itoa(i),
			Ec2Id:          [3]string{"1", "2", "3"},
			BlockDeviceId:  [3]string{"11", "22", "33"},
			KeyName:        "mykey" + strconv.Itoa(i),
			KeyFingerPrint: "mykey" + strconv.Itoa(i),
			KeyMaterial:    "mykey" + strconv.Itoa(i),
			StartTime:      time.Now(),
		}

		aws_data_struct.Player = append(aws_data_struct.Player, p)
	}

}

func main() {
	makeTestPlayer(3)
	fmt.Println(aws_data_struct.Player)

	aws_data_struct.Player = nil
	fmt.Println(aws_data_struct.Player)

	makeTestPlayer(2)
	fmt.Println(aws_data_struct.Player)
	//server := aws_web.CreateWebServer()
	//server.Run()
}

func testValue(val *int) int {
	*val = 10
	return *val
}
