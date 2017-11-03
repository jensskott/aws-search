package ec2

func (e *Ec2Implementation) Ec2DescribeReservedInstances() ([]interface{}, error) {
	var dataSlice []interface{}
	var data interface{}

	// Describe all describe in the region
	resp, err := e.Svc.DescribeReservedInstances(nil)
	if err != nil {
		return nil, err
	}

	for _, a := range resp.ReservedInstances {
		data = a
		dataSlice = append(dataSlice, data)
	}

	return dataSlice, nil
}
