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

func TestEc2DescribeSecurityGroups(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeSecurityGroupsOutput{
		SecurityGroups: []*ec2.SecurityGroup{
			{
				VpcId:     aws.String("vpc-cb4914a3"),
				GroupId:   aws.String("sg-ce2c21a5"),
				GroupName: aws.String("default"),
			},
			{
				VpcId:     aws.String("vpc-cb4914a3"),
				GroupId:   aws.String("sg-3b2c2150"),
				GroupName: aws.String("mongo"),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeSecurityGroups(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeSecurityGroups()

	assert.NoError(t, err)

	// Need two in slice
	assert.Equal(t, 2, len(testResp))

	var m []*ec2.SecurityGroup

	b, _ := json.Marshal(testResp)
	json.Unmarshal(b, &m)

	// Compare respons with what you want to get
	assert.Equal(t, "vpc-cb4914a3", *m[0].VpcId)
	assert.Equal(t, "sg-ce2c21a5", *m[0].GroupId)
	assert.Equal(t, "default", *m[0].GroupName)
	assert.Equal(t, "vpc-cb4914a3", *m[1].VpcId)
	assert.Equal(t, "sg-3b2c2150", *m[1].GroupId)
	assert.Equal(t, "mongo", *m[1].GroupName)
}

func TestEc2DescribeSecurityGroupsError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeSecurityGroups(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeSecurityGroups()
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
