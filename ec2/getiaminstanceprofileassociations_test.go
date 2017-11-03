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

func TestEc2IamInstanceProfileAssociations(t *testing.T) {
	// Create a mock respons for ec2 describe
	resp := &ec2.DescribeIamInstanceProfileAssociationsOutput{
		IamInstanceProfileAssociations: []*ec2.IamInstanceProfileAssociation{
			{
				InstanceId: aws.String("i-06baca28edca29ce9"),
				State:      aws.String("associated"),
			},
			{
				InstanceId: aws.String("i-321afd87"),
				State:      aws.String("disassociated"),
			},
		},
	}

	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeIamInstanceProfileAssociations(gomock.Any()).Return(resp, nil)

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeIamInstanceProfileAssociations()

	assert.NoError(t, err)

	// Need two in slice

	assert.Equal(t, 2, len(testResp))

	var m []*ec2.IamInstanceProfileAssociation

	b, _ := json.Marshal(testResp)
	json.Unmarshal(b, &m)

	// Compare respons with what you want to get
	assert.Equal(t, "i-06baca28edca29ce9", *m[0].InstanceId)
	assert.Equal(t, "associated", *m[0].State)
	assert.Equal(t, "i-321afd87", *m[1].InstanceId)
	assert.Equal(t, "disassociated", *m[1].State)

}

func TestEc2IamInstanceProfileAssociationsError(t *testing.T) {
	// Setup gomock controller with data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := mock.NewMockEC2API(ctrl)
	mockSvc.EXPECT().DescribeIamInstanceProfileAssociations(gomock.Any()).Return(nil, errors.New("I got a booboo"))

	// Create client
	e := Ec2Implementation{
		Svc: mockSvc,
	}

	// Run describe describe
	testResp, err := e.Ec2DescribeIamInstanceProfileAssociations()
	assert.Error(t, err)

	assert.Nil(t, testResp)

}
