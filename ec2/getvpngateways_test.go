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

func TestEc2DescribeVpnGateways(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeVpnGatewaysOutput{
		VpnGateways: []*ec2.VpnGateway{
			{
				VpnGatewayId: aws.String("vgw-f211f09b"),
			},
			{
				VpnGatewayId: aws.String("vgw-9a4cacf3"),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeVpnGateways(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeVpnGateways()

	assert.NoError(t, err)

	// Need two in slice
	assert.Equal(t, 2, len(testResp))

	var m []*ec2.VpnGateway

	b, _ := json.Marshal(testResp)
	json.Unmarshal(b, &m)

	// Compare respons with what you want to get
	assert.Equal(t, "vgw-f211f09b", *m[0].VpnGatewayId)
	assert.Equal(t, "vgw-9a4cacf3", *m[1].VpnGatewayId)

}

func TestEc2DescribeVpnGatewaysError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeVpnGateways(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeVpnGateways()
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
