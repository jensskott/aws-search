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

func TestEc2DescribeNetworkInterfaces(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeNetworkInterfacesOutput{
		NetworkInterfaces: []*ec2.NetworkInterface{
			{
				SubnetId:           aws.String("subnet-98b7bff1"),
				NetworkInterfaceId: aws.String("eni-127d857c"),
			},
			{
				SubnetId:           aws.String("subnet-98b7bff1"),
				NetworkInterfaceId: aws.String("eni-12c9367c"),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeNetworkInterfaces(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeNetworkInterfaces()

	assert.NoError(t, err)

	// Need two in slice
	assert.Equal(t, 2, len(testResp))

	var m []*ec2.NetworkInterface

	b, _ := json.Marshal(testResp)
	json.Unmarshal(b, &m)

	// Compare respons with what you want to get
	assert.Equal(t, "subnet-98b7bff1", *m[0].SubnetId)
	assert.Equal(t, "eni-127d857c", *m[0].NetworkInterfaceId)
	assert.Equal(t, "subnet-98b7bff1", *m[1].SubnetId)
	assert.Equal(t, "eni-12c9367c", *m[1].NetworkInterfaceId)
}

func TestEc2DescribeNetworkInterfacesError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeNetworkInterfaces(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeNetworkInterfaces()
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
