package time_based_key_value_store

import "sort"

type Store struct {
	timestamps []int
	data       map[int]string
}

type TimeMap struct {
	data map[string]*Store
}

func Constructor() TimeMap {
	return TimeMap{data: make(map[string]*Store)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	var data *Store
	if _, ok := this.data[key]; !ok {
		data = &Store{data: make(map[int]string)}
		this.data[key] = data
	} else {
		data = this.data[key]
	}

	data.data[timestamp] = value
	data.timestamps = append(data.timestamps, timestamp)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	data, ok := this.data[key]
	if !ok {
		return ""
	}

	i := sort.Search(len(data.timestamps), func(i int) bool { return data.timestamps[i] > timestamp }) - 1

	if i < 0 {
		return ""
	}

	return data.data[data.timestamps[i]]
}