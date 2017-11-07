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

func TestEc2DescribeRouteTables(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeRouteTablesOutput{
		RouteTables: []*ec2.RouteTable{
			{
				VpcId:        aws.String("vpc-bd2febd4"),
				RouteTableId: aws.String("rtb-f23bfd9b"),
			},
			{
				VpcId:        aws.String("vpc-d253e3bb"),
				RouteTableId: aws.String("rtb-7c6ed015"),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeRouteTables(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeRouteTables()

	assert.NoError(t, err)

	// Need two in slice
	assert.Equal(t, 2, len(testResp))

	var m []*ec2.RouteTable

	b, _ := json.Marshal(testResp)
	json.Unmarshal(b, &m)

	// Compare respons with what you want to get
	assert.Equal(t, "vpc-bd2febd4", *m[0].VpcId)
	assert.Equal(t, "rtb-f23bfd9b", *m[0].RouteTableId)
	assert.Equal(t, "vpc-d253e3bb", *m[1].VpcId)
	assert.Equal(t, "rtb-7c6ed015", *m[1].RouteTableId)

}

func TestEc2DescribeRouteTablesError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeRouteTables(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeRouteTables()
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
