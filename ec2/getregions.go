package ec2

func (e *Ec2Implementation) Ec2GetRegions() ([]string, error) {
	var regions []string
	// Get all regions
	resp, err := e.Svc.DescribeRegions(nil)
	if err != nil {
		return nil, err
	}

	for _, r := range resp.Regions {
		regions = append(regions, *r.RegionName)
	}

	return regions, nil
}
