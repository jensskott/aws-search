package ec2

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/mock/gomock"
	mock "github.com/jensskott/aws-search/_mocks"
	"github.com/stretchr/testify/assert"
)

func TestEc2GetCustomerGateway(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeCustomerGatewaysOutput{
		CustomerGateways: []*ec2.CustomerGateway{
			{
				CustomerGatewayId: aws.String("cgw-34d65f04"),
				IpAddress:         aws.String("165.84.162.17"),
			},
			{
				CustomerGatewayId: aws.String("cgw-c3129ef3"),
				IpAddress:         aws.String("82.99.52.58"),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeCustomerGateways(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	testFilter := []string{"ip-address 82.99.52.58", "ip-address 165.84.162.17"}
	// Run describe describe
	testResp, err := e.Ec2DescribeCustomerGateway(testFilter)

	assert.NoError(t, err)

	// Need two in slice
	assert.Equal(t, 2, len(testResp))

	// Compare respons with what you want to get
	assert.Equal(t, "cgw-34d65f04", *testResp[0].CustomerGatewayId)
	assert.Equal(t, "165.84.162.17", *testResp[0].IpAddress)
	assert.Equal(t, "cgw-c3129ef3", *testResp[1].CustomerGatewayId)
	assert.Equal(t, "82.99.52.58", *testResp[1].IpAddress)

}

func TestEc2DescribeCustomerGatewayError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeCustomerGateways(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeCustomerGateway([]string{})
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
