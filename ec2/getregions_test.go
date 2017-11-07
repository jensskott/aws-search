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

func TestEc2GetRegions(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeRegionsOutput{
		Regions: []*ec2.Region{
			{
				RegionName: aws.String("eu-west-1"),
			},
			{
				RegionName: aws.String("ap-southeast-1"),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeRegions(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2GetRegions()

	assert.NoError(t, err)

	// Need two in slice

	assert.Equal(t, 2, len(testResp))

	// Compare respons with what you want to get
	assert.Equal(t, "eu-west-1", testResp[0])
	assert.Equal(t, "ap-southeast-1", testResp[1])

}

func TestEc2GetRegionsError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeRegions(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2GetRegions()
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
