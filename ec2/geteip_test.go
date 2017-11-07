package ec2

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/mock/gomock"
	mock "github.com/jensskott/aws-search/mocks"
	"github.com/stretchr/testify/assert"
)

func TestEc2DescribeEip(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeAddressesOutput{
		Addresses: []*ec2.Address{
			{
				PublicIp: aws.String("52.52.0.12"),
			},
			{
				PublicIp: aws.String("32.18.22.24"),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeAddresses(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	testFilter := []string{"public-ip 52.52.0.12", "public-ip 32.18.22.24"}

	// Run describe describe
	testResp, err := e.Ec2DescribeEips(testFilter)

	assert.NoError(t, err)

	// Need two in slice
	assert.Equal(t, 2, len(testResp))

	// Compare respons with what you want to get
	assert.Equal(t, "52.52.0.12", *testResp[0].PublicIp)
	assert.Equal(t, "32.18.22.24", *testResp[1].PublicIp)

}

func TestEc2DescribeEipsError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeAddresses(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeEips([]string{})
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
