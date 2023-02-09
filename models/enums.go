package models

type (
	AcilDurum    string
	YardimDurumu string
)

const (
	Normal AcilDurum = "normal"
	Orta   AcilDurum = "orta"
	Kritik AcilDurum = "kritik"

	Bekleniyor YardimDurumu = "bekleniyor"
	Yolda      YardimDurumu = "yolda"
	Yapildi    YardimDurumu = "yapildi"
)
