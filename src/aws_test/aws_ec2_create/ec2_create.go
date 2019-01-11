package aws_ec2_create

import (
	"aws_test/aws_data_struct"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
	"os"
	"time"
)

func Create_DeepLearning_Player_Instances(svc *ec2.EC2, TeamName string) {
	info := aws_data_struct.PlayerEc2Status{}

	info.TeamName = TeamName
	info.KeyName, info.KeyFingerPrint, info.KeyMaterial = Create_Key(svc, TeamName)
	info.Ec2Id[0] = Create_Instance(svc, info.KeyName, "t3.medium", 20, TeamName)
	info.Ec2Id[1] = Create_Instance(svc, info.KeyName, "t3.medium", 20, TeamName)
	info.Ec2Id[1] = Create_Instance(svc, info.KeyName, "t3.medium", 20, TeamName)

	//info.Ec2Id[1] = Create_Instance(svc, info.KeyName, "p2.xlarge", 100)
	//info.Ec2Id[2] = Create_Instance(svc, info.KeyName, "p2.xlarge", 100)

	info.StartTime = time.Now().UTC()
	aws_data_struct.Player = append(aws_data_struct.Player, info)
}

func Create_Instance(svc *ec2.EC2, keyName string, instanceType string, ebsVolume int64, TeamName string) string {
	// Specify the details of the instance that you want to create.

	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		AdditionalInfo: nil,
		BlockDeviceMappings: []*ec2.BlockDeviceMapping{
			{
				DeviceName: aws.String("/dev/sda1"),
				Ebs: &ec2.EbsBlockDevice{
					DeleteOnTermination: aws.Bool(true),
					Encrypted:           nil,
					Iops:                nil,
					KmsKeyId:            nil,
					SnapshotId:          nil,
					VolumeSize:          aws.Int64(ebsVolume),
					VolumeType:          aws.String("gp2"),
				},
				NoDevice:    nil,
				VirtualName: nil,
			},
		},
		//BlockDeviceMappings: nil,
		CapacityReservationSpecification: nil,
		ClientToken:                      nil,
		CpuOptions:                       nil,
		CreditSpecification:              nil,
		DisableApiTermination:            nil,
		DryRun:                           nil,
		EbsOptimized:                     nil,
		ElasticGpuSpecification:          nil,
		ElasticInferenceAccelerators:     nil,
		HibernationOptions:               nil,
		IamInstanceProfile:               nil,
		// An Amazon Linux AMI ID for t2.micro instances in the us-west-2 region
		ImageId:                           aws.String("ami-0e0f4ff1154834540"),
		InstanceInitiatedShutdownBehavior: nil,
		InstanceMarketOptions:             nil,
		InstanceType:                      aws.String(instanceType),
		Ipv6AddressCount:                  nil,
		Ipv6Addresses:                     nil,
		KernelId:                          nil,
		KeyName:                           aws.String(keyName),
		//KeyName: nil,
		LaunchTemplate:        nil,
		LicenseSpecifications: nil,
		MaxCount:              aws.Int64(1),
		MinCount:              aws.Int64(1),
		Monitoring:            nil,
		NetworkInterfaces:     nil,
		Placement:             nil,
		PrivateIpAddress:      nil,
		RamdiskId:             nil,
		SecurityGroupIds:      nil,
		SecurityGroups:        nil,
		SubnetId:              nil,
		TagSpecifications:     nil,
		UserData:              nil,
	})

	if err != nil {
		fmt.Println("Could not create instance", err)
		return ""
	}

	// Add tags to the created instance
	_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{runResult.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("uxfac_test" + TeamName),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", runResult.Instances[0].InstanceId, errtag)
		return ""
	}

	return *runResult.Instances[0].InstanceId
}

func Create_Key(svc *ec2.EC2, team_name string) (string, string, string) {
	// Create the key
	result, err := svc.CreateKeyPair(&ec2.CreateKeyPairInput{
		DryRun:  nil,
		KeyName: aws.String(team_name),
	})

	if err != nil {
		fmt.Println("Got error creating key: ", err)
		os.Exit(1)
	}

	return *result.KeyName, *result.KeyFingerprint, *result.KeyMaterial
}

func Create_Image(svc *ec2.EC2, keyName string, blockDeviceId string) {
	opts := &ec2.CreateImageInput{
		Description: aws.String("image description"),
		InstanceId:  aws.String("i-abcdef12"),
		Name:        aws.String("image name"),
		BlockDeviceMappings: []*ec2.BlockDeviceMapping{
			{
				DeviceName: aws.String("/dev/sda1"),
				NoDevice:   aws.String(""),
			},
			{
				DeviceName: aws.String("/dev/sdb"),
				NoDevice:   aws.String(""),
			},
			{
				DeviceName: aws.String("/dev/sdc"),
				NoDevice:   aws.String(""),
			},
		},
	}
	resp, err := svc.CreateImage(opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ID: ", resp.ImageId)
}

func Create_BlockDevice(svc *ec2.EC2) {
	input := &ec2.CreateVolumeInput{
		AvailabilityZone:  aws.String("ap-northeast-2a"),
		DryRun:            nil,
		Encrypted:         nil,
		Iops:              nil,
		KmsKeyId:          nil,
		Size:              aws.Int64(20),
		SnapshotId:        nil,
		TagSpecifications: nil,
		VolumeType:        aws.String("gp2"),
	}
	result, err := svc.CreateVolume(input)
	if err != nil {
		fmt.Println("Could not create BlockDevice", err)
		return
	}

	fmt.Println("Create BlockDevice", *result.VolumeId)
}
