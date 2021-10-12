package main

import "fmt"

func prime(n int) []int {

	// チェック用の配列作成
	arr := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = true
	}

	// ルートnまでに存在する全約数で消込み
	for i := 2; i*i <= n; i++ {
		for j := 1; i*j <= n; j++ {
			if j != 1 {
				arr[i*j] = false
			}
		}
	}

	// 有効な整数だけ列挙
	var ans []int
	for i := 1; i <= n; i++ {
		if arr[i] {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	fmt.Println(prime(10000))
}
