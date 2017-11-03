package ec2

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type Ec2Implementation struct {
	Session *session.Session
	Svc     ec2iface.EC2API
}
