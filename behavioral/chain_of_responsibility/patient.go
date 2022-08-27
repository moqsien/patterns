package responsibility

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}
