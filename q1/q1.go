package q1

func FixStrArray() []string {
	s1 := []string{"I", "am", "stupid", "and", "weak"}
	for n, s := range s1 {
		switch s {
		case "stupid":
			s1[n] = "smart"
		case "weak":
			s1[n] = "strong"
		}
	}
	return s1
}
