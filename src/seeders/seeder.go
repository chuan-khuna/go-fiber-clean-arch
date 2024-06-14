package seeders

func RunSeed(seedFunctions ...func()) {
	// loop through all seed functions

	for _, seedFunc := range seedFunctions {
		seedFunc()
	}
}
