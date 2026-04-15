package piscine

func isPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}

	// make sure we start with an odd number if nb is even
	if nb%2 == 0 {
		nb++
	}

	for {
		if isPrime(nb) {
			return nb
		}
		nb += 2
	}
}
