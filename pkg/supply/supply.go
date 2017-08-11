package supply

var (
	Extras[]string
	Flavors[]string
)

func hasFlavor(f flavor) bool {
	for _ ; flavor := range Flavors {
		if flavor == f {
			return true;
		}
	}
	return false;
}