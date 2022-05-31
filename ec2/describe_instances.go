package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	axolgoBase "github.com/tchiunam/axolgo-aws/base"
)

// EC2DescribeInstancesAPI defines the interface for the DescribeInstances function.
// We use this interface to test the function using a mocked service.
type EC2DescribeInstancesAPI interface {
	DescribeInstances(ctx context.Context,
		params *ec2.DescribeInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

// GetInstances retrieves information about EC2 instance.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a DescribeInstancesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to DescribeInstances.
func DescribeInstances(c context.Context, api EC2DescribeInstancesAPI, input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return api.DescribeInstances(c, input)
}

// Describe EC2 instances.
func RunDescribeInstances(input *ec2.DescribeInstancesInput, optFns ...func(*axolgoBase.AWSConfigOptions)) (aws.Config, *ec2.DescribeInstancesOutput, error) {
	cfg, err := axolgoBase.LoadAWSConfig(axolgoBase.WithRegion("ap-southeast-1"))
	if err != nil {
		return cfg, nil, err
	}
	client := ec2.NewFromConfig(cfg)

	result, err := DescribeInstances(context.TODO(), client, input)
	if err != nil {
		return cfg, nil, fmt.Errorf("Encountered error when describing EC2 instances: %v", err)
	}

	return cfg, result, nil
}
