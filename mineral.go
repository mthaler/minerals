package main

type Mineral struct {
	Name          string
	Type          string
	Hardness      string
	Density       string
	Crystalsystem string
}

func (m Mineral) toSlice() []string {
	var result []string
	result = make([]string, 0)
	result = append(result, m.Name)
	result = append(result, m.Type)
	result = append(result, m.Hardness)
	result = append(result, m.Density)
	result = append(result, m.Crystalsystem)
	return result
}
