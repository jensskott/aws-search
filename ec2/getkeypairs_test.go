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

func TestEc2DescribeKeypairs(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeKeyPairsOutput{
		KeyPairs: []*ec2.KeyPairInfo{
			{
				KeyName: aws.String("aws_key"),
			},
			{
				KeyName: aws.String("masterkey"),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeKeyPairs(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	testFilter := []string{"key-name aws_key", "key-name masterkey"}
	// Run describe describe
	testResp, err := e.Ec2DescribeKeypairs(testFilter)

	assert.NoError(t, err)

	// Need two in slice
	assert.Equal(t, 2, len(testResp))

	var m []*ec2.KeyPairInfo

	b, _ := json.Marshal(testResp)
	json.Unmarshal(b, &m)

	// Compare respons with what you want to get
	assert.Equal(t, "aws_key", *testResp[0].KeyName)
	assert.Equal(t, "masterkey", *testResp[1].KeyName)

}

func TestEc2DescribeKeypairsError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeKeyPairs(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeKeypairs([]string{})
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
