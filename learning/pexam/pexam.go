package main

type patient struct {
	name string
	height int
	weight int
	chest int
	upb int
	hbp int
	lbp int
}

// patientのインスタンスを戻り値として返す
func newpatient(name string) patient {
	return patient{name, 0, 0, 0, 0, 0 ,0}
}