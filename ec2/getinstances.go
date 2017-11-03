package ec2

func (e *Ec2Implementation) Ec2DescribeInstances() ([]interface{}, error) {
	var dataSlice []interface{}
	var data interface{}

	// Describe all describe in the region
	resp, err := e.Svc.DescribeInstances(nil)
	if err != nil {
		return nil, err
	}

	for idx := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			data = inst
			dataSlice = append(dataSlice, data)
		}
	}
	return dataSlice, nil
}
