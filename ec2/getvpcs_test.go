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

func TestEc2DescribeVpcs(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeVpcsOutput{
		Vpcs: []*ec2.Vpc{
			{
				VpcId:     aws.String("vpc-a01106c2"),
				IsDefault: aws.Bool(true),
			},
			{
				VpcId:     aws.String("vpc-0s932000"),
				IsDefault: aws.Bool(false),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeVpcs(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeVpcs()

	assert.NoError(t, err)

	// Need two in slice
	assert.Equal(t, 2, len(testResp))

	var m []*ec2.Vpc

	b, _ := json.Marshal(testResp)
	json.Unmarshal(b, &m)

	// Compare respons with what you want to get
	assert.Equal(t, "vpc-a01106c2", *m[0].VpcId)
	assert.Equal(t, true, *m[0].IsDefault)
	assert.Equal(t, "vpc-0s932000", *m[1].VpcId)
	assert.Equal(t, false, *m[1].IsDefault)

}

func TestEc2DescribeVpcsError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeVpcs(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeVpcs()
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
