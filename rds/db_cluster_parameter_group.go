package rds

import (
	"context"
	"fmt"

	"k8s.io/klog/v2"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	awsTypes "github.com/aws/aws-sdk-go-v2/service/rds/types"
	axolgoBase "github.com/tchiunam/axolgo-aws/base"
	"github.com/tchiunam/axolgo-lib/types"
)

// ModifyDBClusterParameterGroupAPI defines the interface for the ModifyDBClusterParameterGroup function.
// We use this interface to test the function using a mocked service.
type ModifyDBClusterParameterGroupAPI interface {
	ModifyDBClusterParameterGroup(ctx context.Context,
		params *rds.ModifyDBClusterParameterGroupInput,
		optFns ...func(*rds.Options)) (*rds.ModifyDBClusterParameterGroupOutput, error)
}

// ModifyDBClusterParameterGroup modifies the DB Cluster Parameter Group.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a ModifyDBClusterParameterGroupOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to DescribeDBClusterParameterGroup.
func ModifyDBClusterParameterGroup(c context.Context, api ModifyDBClusterParameterGroupAPI, input *rds.ModifyDBClusterParameterGroupInput) (*rds.ModifyDBClusterParameterGroupOutput, error) {
	return api.ModifyDBClusterParameterGroup(c, input)
}

// Modify the parameters of a DB Cluster Parameter Group.
func RunModifyDBClusterParameterGroup(parameterGroupName string, staticParameters []types.Parameter, dynamicParameters []types.Parameter, optFns ...func(*axolgoBase.AWSConfigOptions)) (aws.Config, error) {
	paramsSize := len(staticParameters) + len(dynamicParameters)
	klog.Infof("Parameter Group name: %v", parameterGroupName)
	klog.Infof("No. of parameters to set: %v", paramsSize)

	cfg, err := axolgoBase.LoadAWSConfig(axolgoBase.WithRegion("ap-southeast-1"))
	if err != nil {
		return cfg, err
	}
	client := rds.NewFromConfig(cfg)

	// Set the Apply Method based on the Apply Type.
	var params []awsTypes.Parameter
	for _, p := range staticParameters {
		klog.V(6).InfoS("static parameter", *p.Name, *p.Value)
		params = append(
			params,
			awsTypes.Parameter{
				ApplyMethod:    "pending-reboot",
				ParameterName:  aws.String(*p.Name),
				ParameterValue: aws.String(*p.Value),
			})
	}
	for _, p := range dynamicParameters {
		klog.V(6).InfoS("dynamic parameter", *p.Name, *p.Value)
		params = append(
			params,
			awsTypes.Parameter{
				ApplyMethod:    "immediate",
				ParameterName:  aws.String(*p.Name),
				ParameterValue: aws.String(*p.Value),
			})
	}

	// AWS has a restriction that a maximum of 20 parameters can be modified in a single request.
	const batchSize int = 20

	// Modify parameters in batches.
	for batchNo := 0; batchSize*batchNo < paramsSize; batchNo++ {
		start := batchSize * batchNo
		end := start + batchSize
		if end >= paramsSize {
			end = paramsSize - 1
		}
		batchParams := params[start:end]
		input := &rds.ModifyDBClusterParameterGroupInput{
			DBClusterParameterGroupName: &parameterGroupName,
			Parameters:                  batchParams,
		}

		klog.Infof("Sending the update of batch %v", batchNo)
		// No useful information in the result metadata.
		_, err := ModifyDBClusterParameterGroup(context.TODO(), client, input)
		if err != nil {
			return cfg, fmt.Errorf("Failed to modify Cluster Parameter Group %q: %v", parameterGroupName, err)
		}
	}

	return cfg, nil
}
