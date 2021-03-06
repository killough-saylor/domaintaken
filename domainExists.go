package main

import (
	"errors"
	"strconv"
)

func domainExists(name string) (bool, error) {
	req := &DNSRequest{
		Name: name,
		Type: "NS",
	}

	resp, err := PerformRequest(req)

	if err != nil {
		return false, err
	}

	switch resp.Status {
	case 0:
		return true, nil
	case 3:
		return false, nil
	default:
		return false, errors.New("unexpected DNS status " + strconv.Itoa(resp.Status))
	}
}
