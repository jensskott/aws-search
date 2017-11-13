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

func TestEc2DescribeInternetGateways(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeInternetGatewaysOutput{
		InternetGateways: []*ec2.InternetGateway{
			{
				InternetGatewayId: aws.String("igw-992fb9f0"),
			},
			{
				InternetGatewayId: aws.String("igw-1f9f7376"),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeInternetGateways(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	testFilter := []string{"internet-gateway-id igw-992fb9f0", "internet-gateway-id igw-1f9f7376"}
	// Run describe describe
	testResp, err := e.Ec2DescribeInternetGateways(testFilter)

	assert.NoError(t, err)

	// Need two in slice
	assert.Equal(t, 2, len(testResp))

	var m []*ec2.InternetGateway

	b, _ := json.Marshal(testResp)
	json.Unmarshal(b, &m)

	// Compare respons with what you want to get
	assert.Equal(t, "igw-992fb9f0", *testResp[0].InternetGatewayId)
	assert.Equal(t, "igw-1f9f7376", *testResp[1].InternetGatewayId)

}

func TestEc2DescribeInternetGatewaysError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeInternetGateways(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeInternetGateways([]string{})
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
