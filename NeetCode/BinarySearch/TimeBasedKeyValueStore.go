package binarysearch

type TimeStampData struct {
	timeStamp int
	value     string
}

type TimeMap struct {
	keyVal map[string][]TimeStampData
}

func Constructor() TimeMap {
	return TimeMap{
		keyVal: make(map[string][]TimeStampData),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.keyVal[key] = append(this.keyVal[key], TimeStampData{timeStamp: timestamp, value: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, ok := this.keyVal[key]; !ok {
		return ""
	}

	vals := this.keyVal[key]
	leftPtr, rightPtr := 0, len(vals)-1

	var res string
	for leftPtr <= rightPtr {
		midPtr := (leftPtr + rightPtr) / 2
		if vals[midPtr].timeStamp < timestamp {
			res = vals[midPtr].value
			leftPtr = midPtr + 1
		} else if vals[midPtr].timeStamp > timestamp {
			rightPtr = midPtr - 1
		} else {
			return vals[midPtr].value
		}
	}

	return res
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
