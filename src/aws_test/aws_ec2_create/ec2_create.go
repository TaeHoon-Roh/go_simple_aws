package aws_ec2_create

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

func Create_Instance(svc *ec2.EC2){
	// Specify the details of the instance that you want to create.
	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		AdditionalInfo:                   nil,
		BlockDeviceMappings:              nil,
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
		InstanceType:                      aws.String("t2.micro"),
		Ipv6AddressCount:                  nil,
		Ipv6Addresses:                     nil,
		KernelId:                          nil,
		KeyName:                           nil,
		LaunchTemplate:                    nil,
		LicenseSpecifications:             nil,
		MaxCount:                          aws.Int64(1),
		MinCount:                          aws.Int64(1),
		Monitoring:                        nil,
		NetworkInterfaces:                 nil,
		Placement:                         nil,
		PrivateIpAddress:                  nil,
		RamdiskId:                         nil,
		SecurityGroupIds:                  nil,
		SecurityGroups:                    nil,
		SubnetId:                          nil,
		TagSpecifications:                 nil,
		UserData:                          nil,
	})

	if err != nil {
		fmt.Println("Could not create instance", err)
		return
	}

	fmt.Println("Created instance", *runResult.Instances[0].InstanceId)

	// Add tags to the created instance
	_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{runResult.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("UxfacTestInstance"),
			},
		},
	})
	if errtag != nil {
		log.Println("Could not create tags for instance", runResult.Instances[0].InstanceId, errtag)
		return
	}

	fmt.Println("Successfully tagged instance")
}
func Create_BlockDevice(svc *ec2.EC2){
	input := &ec2.CreateVolumeInput{
		AvailabilityZone:  nil,
		DryRun:            nil,
		Encrypted:         nil,
		Iops:              nil,
		KmsKeyId:          nil,
		Size:              nil,
		SnapshotId:        nil,
		TagSpecifications: nil,
		VolumeType:        nil,
	}
	result, err := svc.CreateVolume(input)

}
