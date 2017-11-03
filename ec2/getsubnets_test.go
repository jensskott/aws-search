package ec2

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/mock/gomock"
	mock "github.com/jensskott/aws-search/mocks"
	"github.com/stretchr/testify/assert"
)

func TestEc2DescribeSubnets(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeSubnetsOutput{
		Subnets: []*ec2.Subnet{
			{
				CidrBlock: aws.String("10.10.0.1/24"),
				VpcId:     aws.String("vpc-a01106c2"),
			},
			{
				CidrBlock: aws.String("10.11.0.1/24"),
				VpcId:     aws.String("vpc-a01106c2"),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeSubnets(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeSubnets()

	assert.NoError(t, err)

	// Need two in slice
	assert.Equal(t, 2, len(testResp))

	var m []*ec2.Subnet

	b, _ := json.Marshal(testResp)
	json.Unmarshal(b, &m)

	// Compare respons with what you want to get
	assert.Equal(t, "10.10.0.1/24", *m[0].CidrBlock)
	assert.Equal(t, "vpc-a01106c2", *m[0].VpcId)
	assert.Equal(t, "10.11.0.1/24", *m[1].CidrBlock)
	assert.Equal(t, "vpc-a01106c2", *m[1].VpcId)

}

func TestEc2DescribeSubnetsError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeSubnets(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeSubnets()
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
