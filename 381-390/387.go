package _381_390

func firstUniqChar(s string) int {
	hashMap:=make(map[rune]int)
	for _,v:=range s{
		_,found:=hashMap[v]
		if found{
			hashMap[v]=hashMap[v]+1
		}else{
			hashMap[v]=1
		}
	}
	for i,v:=range s{
		_,found:=hashMap[v]
		if found&&hashMap[v]==1{
			return i
		}
	}
	return -1
}
