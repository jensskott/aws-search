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

func TestEc2DescribeSnapshots(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeSnapshotsOutput{
		Snapshots: []*ec2.Snapshot{
			{
				SnapshotId: aws.String("snap-0040519d2a9330694"),
				VolumeSize: aws.Int64(100),
			},
			{
				SnapshotId: aws.String("snap-024c8f9c10a2210ec"),
				VolumeSize: aws.Int64(20),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeSnapshots(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeSnapshots()

	assert.NoError(t, err)

	// Need two in slice
	assert.Equal(t, 2, len(testResp))

	var m []*ec2.Snapshot

	b, _ := json.Marshal(testResp)
	json.Unmarshal(b, &m)

	// Compare respons with what you want to get
	assert.Equal(t, "snap-0040519d2a9330694", *m[0].SnapshotId)
	assert.Equal(t, int64(100), *m[0].VolumeSize)
	assert.Equal(t, "snap-024c8f9c10a2210ec", *m[1].SnapshotId)
	assert.Equal(t, int64(20), *m[1].VolumeSize)

}

func TestEc2DescribeSnapshotsError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeSnapshots(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeSnapshots()
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
