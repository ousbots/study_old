package time_based_key_value_store

import "sort"

type TimeMap struct {
	timestamps map[string][]int
	values     map[string][]string
}

func Constructor() TimeMap {
	return TimeMap{timestamps: make(map[string][]int), values: make(map[string][]string)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.timestamps[key] = append(this.timestamps[key], timestamp)
	this.values[key] = append(this.values[key], value)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	times := this.timestamps[key]
	vals := this.values[key]

	i := sort.Search(len(times), func(i int) bool { return times[i] > timestamp }) - 1

	if i < 0 {
		return ""
	}

	return vals[i]
}