type MyHashMap struct {
	MyMap [][]int
}

func Constructor() MyHashMap {
    return MyHashMap{
		MyMap: [][]int{},
	}
}

func (this *MyHashMap) Put(key int, value int) {
	isFound := false
	for _, pair := range this.MyMap{
		if pair[0] == key{
			pair[1] = value
			isFound = true
		}
	}
    if !isFound{
		this.MyMap = append(this.MyMap, []int{key, value})
	}
}

func (this *MyHashMap) Get(key int) int {
    for _, pair := range this.MyMap{
		if pair[0] == key{
			return pair[1]
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
    for i, pair := range this.MyMap{
		if pair[0] == key{
			this.MyMap = append(this.MyMap[:i], this.MyMap[i+1:]...)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */