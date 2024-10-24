package bai

type Bai4 struct{}

func (n *Bai4) Main(arr []int, low, high int) {
	if high > low {
		pv := piv(arr, low, high)
		n.Main(arr, low, pv-1)
		n.Main(arr, pv+1, high)
	}
}

func piv(arr []int, low, high int) int {
	i := low - 1
	pv := arr[high]
	for j := low; j < high; j++ {
		if arr[j] < pv {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
