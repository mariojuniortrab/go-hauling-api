package protocol_entity

import (
	"strconv"
	"strings"
)

type List struct {
	Page      int
	Limit     int
	OrderBy   string
	OrderType string
	Q         string
}

type ListInputDto struct {
	Limit     string
	Page      string
	OrderBy   string
	OrderType string
	Q         string
}

func FillFromInput(input *ListInputDto, output *List) error {
	page, err := strconv.Atoi(input.Page)
	if err != nil {
		return err
	}

	limit, err := strconv.Atoi(input.Limit)
	if err != nil {
		return err
	}

	output.Page = page
	output.Limit = limit
	output.OrderBy = strings.ToLower(input.OrderBy)
	output.OrderType = strings.ToLower(input.OrderType)
	output.Q = input.Q

	return nil
}
