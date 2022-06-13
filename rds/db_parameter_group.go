/*
Copyright © 2022 tchiunam

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package rds

import (
	"context"
	"fmt"

	"k8s.io/klog/v2"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	awsTypes "github.com/aws/aws-sdk-go-v2/service/rds/types"
	axolgoawsutil "github.com/tchiunam/axolgo-aws/util"
	"github.com/tchiunam/axolgo-lib/types"
)

// ModifyDBParameterGroupAPI defines the interface for the ModifyDBParameterGroup function.
// We use this interface to test the function using a mocked service.
type ModifyDBParameterGroupAPI interface {
	ModifyDBParameterGroup(ctx context.Context,
		params *rds.ModifyDBParameterGroupInput,
		optFns ...func(*rds.Options)) (*rds.ModifyDBParameterGroupOutput, error)
}

// ModifyDBParameterGroup modifies the DB Parameter Group.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a ModifyDBParameterGroupOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to DescribeDBParameterGroup.
func ModifyDBParameterGroup(c context.Context, api ModifyDBParameterGroupAPI, input *rds.ModifyDBParameterGroupInput) (*rds.ModifyDBParameterGroupOutput, error) {
	return api.ModifyDBParameterGroup(c, input)
}

// Modify the parameters of a DB Parameter Group
func RunModifyDBParameterGroup(parameterGroupName string, staticParameters []types.Parameter, dynamicParameters []types.Parameter, optFns ...func(*axolgoawsutil.AWSConfigOptions) error) (aws.Config, error) {
	paramsSize := len(staticParameters) + len(dynamicParameters)
	klog.Infof("Parameter Group name: %v", parameterGroupName)
	klog.Infof("No. of parameters to set: %v", paramsSize)

	cfg, err := axolgoawsutil.LoadAWSConfig(optFns...)
	if err != nil {
		return cfg, err
	}
	client := rds.NewFromConfig(cfg)

	// Set the Apply Method based on the Apply Type
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

	// AWS has a restriction that a maximum of 20 parameters can be modified in a single request
	const batchSize int = 20

	// Modify parameters in batches.
	for batchNo := 0; batchSize*batchNo < paramsSize; batchNo++ {
		start := batchSize * batchNo
		end := start + batchSize
		if end >= paramsSize {
			end = paramsSize - 1
		}
		batchParams := params[start:end]
		input := &rds.ModifyDBParameterGroupInput{
			DBParameterGroupName: &parameterGroupName,
			Parameters:           batchParams,
		}

		klog.Infof("Sending the update of batch %v", batchNo)
		// No useful information in the result metadata
		_, err := ModifyDBParameterGroup(context.TODO(), client, input)
		if err != nil {
			return cfg, fmt.Errorf("Failed to modify Parameter Group %q: %v", parameterGroupName, err)
		}
	}

	return cfg, nil
}
