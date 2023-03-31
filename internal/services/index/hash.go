package index

func getKeyHash(key string) int {
	total := 0
	i := 0

	for i = 0; i < len(key); i++ {
		total += int(key[i])
	}

	if i == 0 {
		return 0
	}

	return (total / i) % maxBuckets
}
