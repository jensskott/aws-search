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

func TestEc2DescribeImages(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeImagesOutput{
		Images: []*ec2.Image{
			{
				ImageId: aws.String("ami-8bf746e4"),
			},
			{
				ImageId: aws.String("ami-dcfe02b3"),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeImages(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeImages()

	assert.NoError(t, err)

	// Need two in slice
	assert.Equal(t, 2, len(testResp))

	var m []*ec2.Image

	b, _ := json.Marshal(testResp)
	json.Unmarshal(b, &m)

	// Compare respons with what you want to get
	assert.Equal(t, "ami-8bf746e4", *m[0].ImageId)
	assert.Equal(t, "ami-dcfe02b3", *m[1].ImageId)

}

func TestEc2DescribeImagesError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeImages(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeImages()
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
