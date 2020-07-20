package main

func solution(s string) string {
	if len(s)%2 == 0 {
		return s[len(s)/2-1 : len(s)/2+1]
	} else {
		return s[len(s)/2 : len(s)/2+1]
	}
}

func solution1(s string) string {
	return s[int(float64(len(s))/2-0.5):int(float64(len(s))/2+1)]
}
