package quiz

// // 0 - 255 regex
// var tffr = `\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])`

// // 16 - 31 regex
// var sttr = `\.(31|[1-2][0-9][0-9]|[1-9]?[0-9])`

// func TestGetRandomIP(t *testing.T) {

// 	rgx := `\b(10|172|192)` + tffr + tffr + tffr + `\b`

// 	for i := 0; i < 1000; i++ {
// 		got := randomIP()
// 		if match, _ := regexp.MatchString(rgx, got); !match {
// 			t.Errorf("RandomIP() = %v, does not match regex\n", got)
// 		}
// 	}
// }

// func TestAddressInRange(t *testing.T) {

// 	tenZeroZeroZero := `\b10` + tffr + tffr + tffr + `\b`
// 	tzzz := addressInRange(10, 0, 0)
// 	if match, _ := regexp.MatchString(tenZeroZeroZero, tzzz); !match {
// 		t.Errorf("RandomIP() = %v, does not match regex\n", tzzz)
// 	}

// 	oneSevenTwo := `\b172` + sttr + tffr + tffr + `\b`
// 	ost := addressInRange(172, 16, 31)
// 	if match, _ := regexp.MatchString(oneSevenTwo, ost); !match {
// 		t.Errorf("RandomIP() = %v, does not match regex\n", ost)
// 	}

// 	oneNineTwo := `\b192` + `\.168` + tffr + tffr + `\b`
// 	ont := addressInRange(192, 168, 168)
// 	if match, _ := regexp.MatchString(oneNineTwo, ont); !match {
// 		t.Errorf("RandomIP() = %v, does not match regex\n", ont)
// 	}
// }

// func TestGetrandomNetwork(t *testing.T) {
// 	var allPossNets []string
// 	for _, v := range networks {
// 		allPossNets = append(allPossNets, v.netmask, v.prefix)
// 	}

// 	iterations := 2048
// 	if isRandStructError(randomNetwork, allPossNets, iterations) {
// 		t.Error("getrandomNetwork() not producing all possible random networks.")
// 	}
// }

// func TestGetrandomQuestionKind(t *testing.T) {

// 	qkinds := make([]string, 0)
// 	for k := range questions {
// 		qkinds = append(qkinds, k)
// 	}

// 	iterations := 128
// 	if isRandStructError(randomQuestionKind, qkinds, iterations) {
// 		t.Error("getQuestionKind() not producing all possible random question kinds.")
// 	}
// }

// func isRandStructError(genFunc func() string, allPermutations []string, iter int) bool {

// 	generatedOutputs := make([]string, iter)
// 	for i := 0; i < iter; i++ {
// 		generatedOutputs[i] = genFunc()
// 	}

// 	var totalUnique int
// 	for _, v1 := range allPermutations {
// 		for _, v2 := range generatedOutputs {
// 			if v1 == v2 {
// 				totalUnique++
// 				break
// 			}
// 		}
// 	}
// 	return totalUnique != len(allPermutations)
// }
