package ec2

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/mock/gomock"
	mock "github.com/jensskott/aws-search/_mocks"
	"github.com/stretchr/testify/assert"
)

func TestEc2DescribeReservedInstances(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeReservedInstancesOutput{
		ReservedInstances: []*ec2.ReservedInstances{
			{
				InstanceCount: aws.Int64(20),
				InstanceType:  aws.String("t2"),
			},
			{
				InstanceCount: aws.Int64(2),
				InstanceType:  aws.String("c4"),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeReservedInstances(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeReservedInstances()

	assert.NoError(t, err)

	// Need two in slice
	assert.Equal(t, 2, len(testResp))

	var m []*ec2.ReservedInstances

	b, _ := json.Marshal(testResp)
	json.Unmarshal(b, &m)

	// Compare respons with what you want to get
	assert.Equal(t, int64(20), *m[0].InstanceCount)
	assert.Equal(t, "t2", *m[0].InstanceType)
	assert.Equal(t, int64(2), *m[1].InstanceCount)
	assert.Equal(t, "c4", *m[1].InstanceType)

}

func TestEc2DescribeReservedInstancesError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeReservedInstances(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeReservedInstances()
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
