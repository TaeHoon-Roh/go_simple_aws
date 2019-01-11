package aws_ec2_delete

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func Delete_DeepLearning_Plyaer_Instances(svc *ec2.EC2, TeamName string) {

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
