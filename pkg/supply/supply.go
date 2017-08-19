package supply

var (
	Extras  = make(map[string]int)
	Flavors = make(map[string]int)

	// Tastes  = []string{
	// 	"bastani",
	// 	"bacon",
	// 	"chocolate",
	// 	"coconut",
	// 	"pistacho",
	// 	"vanilla",
	// }
)

func hasFlavor(f string) bool {
	for i := 0; i < len(Flavors); i++ {
		if Flavors[f] != 0 {
			return true
		}
	}
	return false
}

func HasExtra(e string) bool {
	for i := 0; i < len(Extras); i++ {
		if Extras[e] != 0 {
			return true
		}
	}
	return false
}

func FillFlavors(f map[string]int) {
	for key := range f {
		f[key] += 10
	}
}

func FillExtras(e map[string]bool) {
	for key := range e {
		e[key] = true
	}
}
