package flyweight

import "fmt"

const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerrroristDressType terrorist dress type
	CounterTerrroristDressType = "ctDress"
)

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	// map中已存在相应的dressType
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}
	// map中没有相应的dressType
	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerrroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

var (
	dressFactorySingleInstance = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}
