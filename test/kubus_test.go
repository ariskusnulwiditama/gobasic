package test

type Kubus struct {
	Sisi float64
}

type KubusInterface interface {
	Volume() float64
	LuasPermukaan() float64
}

func NewKubus(sisi float64) *Kubus {
	return &Kubus{Sisi: sisi}
}

func (k *Kubus) Volume() float64 {
	return k.Sisi * k.Sisi * k.Sisi
}

func (k *Kubus) LuasPermukaan() float64 {
	return 6 * k.Sisi * k.Sisi
}

