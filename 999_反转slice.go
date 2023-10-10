package main

func xr(sub []string) {
	for i := 0; i < len(sub)/2; i++ { //反转sub
		sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
	}
}

func main() {
	xr([]string{"1", "2", "3", "4", "5"}) // 0 4; 1 3; 2
}
