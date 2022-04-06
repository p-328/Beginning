package main

import (
	"fmt"
	"os"
	"sort"
)
func find_max(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
func median(nums []int) float64 {
	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j];
	});
	if len(nums) % 2 != 0 {
		midpoint := len(nums) - 1;
		midpoint = midpoint/2;
		return float64(nums[midpoint]);
	} else {
		sec1End := len(nums) / 2;
		sec1End -= 1;
		sec2End := sec1End + 1;
		median := float64(nums[sec1End]) + float64(nums[sec2End]);
		median /= 2;
		return median;
	}
}
func find_min(arr []int) int {
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
func find_outliers(nums []int) ([]int, bool) {
	var outliers []int
	var sec1 []int
	var sec2 []int
	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] < nums[j];
	});
	if len(nums) % 2 != 0 {
		midpoint := len(nums) - 1;
		midpoint = midpoint/2;
		sec1 = nums[:midpoint];
		sec2 = nums[midpoint+1:];
	} else {
		sec1End := len(nums) / 2;
		sec1End -= 1;
		sec2End := sec1End + 1;
		sec1 = nums[:sec2End];
		sec2 = nums[sec2End:];
	}
	q1 := median(sec1);
	q3 := median(sec2);
	iqr := q3 - q1;
	for i := 0; i < len(nums); i++ {
		if float64(nums[i]) < q1 - (1.5*iqr) {
			outliers = append(outliers, nums[i]);
		}
		if float64(nums[i]) > q3 + (1.5*iqr) {
			outliers = append(outliers, nums[i]);
		}
	}
	if (len(outliers) > 0) {
		return outliers, true;
	}
	return nil, false;
}
func main() {
	var num int;
	fmt.Println("Enter length of numbers you want to generate");
	fmt.Scanf("%d\n", &num);
	if num <= 1 {
		fmt.Println("Please enter a valid amount of numbers");
		os.Exit(-1);
	}
	var nums []int = make([]int, num);
	fmt.Println("Enter numbers: ");
	for i := 0; i < len(nums); i++ {
		fmt.Scanf("%d\n", &nums[i]);
	}
	var conglomerateSum int = 0;
	for i := 0; i < len(nums); i++ {
		conglomerateSum += nums[i];
	}
	var mean float64 = float64(conglomerateSum)/float64(len(nums));
	fmt.Println("Mean:", mean);
	fmt.Println("Range:", find_max(nums)-find_min(nums));
	outliers, any := find_outliers(nums);
	if any == false {
		fmt.Println("No outliers found in dataset.")
	} else {
		fmt.Println("Outliers found");
		fmt.Println("Outliers: ");
		for i := 0; i < len(outliers); i++ {
			fmt.Println(outliers[i]);
		}
	}
	fmt.Println("Median:", median(nums));
}
