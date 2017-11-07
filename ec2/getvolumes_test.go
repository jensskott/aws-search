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

func TestEc2DescribeVolumes(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeVolumesOutput{
		Volumes: []*ec2.Volume{
			{
				VolumeId: aws.String("vol-049df61146c4d7901"),
				Size:     aws.Int64(20),
			},
			{
				VolumeId: aws.String("vol-a5290j09921c09219"),
				Size:     aws.Int64(1000),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeVolumes(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeVolumes()

	assert.NoError(t, err)

	// Need two in slice
	assert.Equal(t, 2, len(testResp))

	var m []*ec2.Volume

	b, _ := json.Marshal(testResp)
	json.Unmarshal(b, &m)

	// Compare respons with what you want to get
	assert.Equal(t, "vol-049df61146c4d7901", *m[0].VolumeId)
	assert.Equal(t, int64(20), *m[0].Size)
	assert.Equal(t, "vol-a5290j09921c09219", *m[1].VolumeId)
	assert.Equal(t, int64(1000), *m[1].Size)

}

func TestEc2DescribeVolumesError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeVolumes(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeVolumes()
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
