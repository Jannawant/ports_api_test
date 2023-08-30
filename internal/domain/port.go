package domain

import "fmt"

type Port struct {
	id          string
	name        string
	code        string
	city        string
	country     string
	alias       []string
	regions     []string
	coordinates []float64
	province    string
	timezone    string
	unlocs      []string
}

func NewPort(id, name, code, city, country string, alias, regions []string, coords []float64,
	province, tz string, unlocs []string) (*Port, error) {
	if id == "" {
		return nil, fmt.Errorf("#{ErrRequired}: port id is required")
	}
	if name == "" {
		return nil, fmt.Errorf("#{ErrRequired}: port name is required")
	}
	if city == "" {
		return nil, fmt.Errorf("#{ErrRequired}: port city is required")
	}
	if country == "" {
		return nil, fmt.Errorf("#{ErrRequired}: port country is required")
	}
	return &Port{
		id:          id,
		name:        name,
		code:        code,
		city:        city,
		country:     country,
		alias:       alias,
		regions:     regions,
		coordinates: coords,
		province:    province,
		timezone:    tz,
		unlocs:      unlocs,
	}, nil
}

func (p *Port) ID() string {
	return p.id
}

func (p *Port) Name() string {
	return p.name
}

func (p *Port) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("#{ErrRequired}: port name is required")
	}
	p.name = name
	return nil
}

func (p *Port) Code() string {
	return p.code
}

func (p *Port) City() string {
	return p.city
}

func (p *Port) Country() string {
	return p.country
}

func (p *Port) Alias() []string {
	return p.alias
}

func (p *Port) Regions() []string {
	return p.regions
}

func (p *Port) Coordinates() []float64 {
	return p.coordinates
}

func (p *Port) Province() string {
	return p.province
}

func (p *Port) Timezone() string {
	return p.timezone
}

func (p *Port) Unlocs() []string {
	return p.unlocs
}

