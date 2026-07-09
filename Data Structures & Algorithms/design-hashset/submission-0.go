type MyHashSet struct {
	hs []int
}

func Constructor() MyHashSet {
    return MyHashSet{
		hs: []int{},
	}
}

func (this *MyHashSet) Add(key int) {
	isFound := false
    for _, item := range this.hs{
		if item == key{
			isFound = true
			break
		}
	}
	if !isFound{
		this.hs = append(this.hs, key)
	}
}

func (this *MyHashSet) Remove(key int) {
    for i, item := range this.hs{
		if item == key{
			this.hs = append(this.hs[:i], this.hs[i+1:]...)
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
    for _, item := range this.hs{
		if item == key{
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 