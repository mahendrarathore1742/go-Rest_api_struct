package main

type apierror struct {
	Err    string
	Status int
}

func (e apierror) Error() string {
	return e.Err
}
