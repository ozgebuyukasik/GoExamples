package main

func GetMostFreqKey(freqMap map[int32]int32) int32 {
	var mostFreqCount int32
	var mostFreqType int32
	for key := range freqMap {
		if freqMap[key] > mostFreqCount {
			mostFreqCount = freqMap[key]
			mostFreqType = key
		}
	}
	return mostFreqType
}

func MigratoryBirds(arr []int32) int32 {
	frequenceOfTypes := make(map[int32]int32)
	for birdType := range arr {
		count, ok := frequenceOfTypes[arr[birdType]]
		if ok {
			frequenceOfTypes[arr[birdType]] = count + 1
		} else {
			frequenceOfTypes[arr[birdType]] = 1
		}
	}
	return GetMostFreqKey(frequenceOfTypes)

}
