package weather

// Forecaster can query for the conditions in a given
// location
type Forecaster interface {
	Forecast(location string) (*Conditions, error)
}

// ForecasterFunc implements Forecaster calling itself
type ForecasterFunc func(string) (*Conditions, error)

// Forecast returns the current conditions for the given location
func (f ForecasterFunc) Forecast(location string) (forecast *Conditions, err error) {
	return f(location)
}

// Conditions describes a set of info about the
// weather in a location on a single point in turn
type Conditions struct {
	Location    string
	Celsius     int
	Description string
}
