package ec2

func (e *Ec2Implementation) Ec2DescribeKeypairs() ([]interface{}, error) {
	var dataSlice []interface{}
	var data interface{}

	// Describe all describe in the region
	resp, err := e.Svc.DescribeKeyPairs(nil)
	if err != nil {
		return nil, err
	}

	for _, a := range resp.KeyPairs {
		data = a
		dataSlice = append(dataSlice, data)
	}

	return dataSlice, nil
}
