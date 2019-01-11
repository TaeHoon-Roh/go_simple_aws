package aws_ec2_delete

import (
	"aws_test/aws_data_struct"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func Delete_DeepLearning_Plyaer_Instances(svc *ec2.EC2) {
	for index := range aws_data_struct.Player {
		p := aws_data_struct.Player[index]
		for i := 0; i < 3; i++ {
			Terminate_Instance(svc, p.Ec2Id[i])
		}
		Delete_Key(svc, p.KeyName)
	}
	aws_data_struct.Player = nil
}

func Terminate_Instance(svc *ec2.EC2, instanceId string) {
	result, err := svc.TerminateInstances(&ec2.TerminateInstancesInput{
		DryRun: nil,
		InstanceIds: []*string{
			aws.String(instanceId),
		},
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Ec2 Terminate ", result.TerminatingInstances)
}

func Delete_Key(svc *ec2.EC2, keyname string) {
	result, err := svc.DeleteKeyPair(&ec2.DeleteKeyPairInput{
		DryRun:  nil,
		KeyName: aws.String(keyname),
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Keypair Delete ", result.String())
}
