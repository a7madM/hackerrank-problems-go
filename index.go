package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// test()
	// fmt.Println(ClacChoclateCount(43203, 60, 5))
	// fmt.Println(ClacChoclateCount(10, 2, 5))
	// fmt.Println(ClacChoclateCount(12, 4, 4))
	// fmt.Println(ClacChoclateCount(6, 2, 2))
	// fmt.Println(ClacChoclateCount(5483, 40, 17))
	// fmt.Println(ClacChoclateCount(20741, 60, 6676))
	// fmt.Println(ClacChoclateCount(33463, 28870, 18710))
	// fmt.Println(ClacChoclateCount(15643, 153, 13015))
	testChoclate()
}

func testChoclate() {
	var t int
	Result := []int{13, 1, 1, 211, 364, 161, 168, 1, 1850, 109, 432, 108, 322, 1, 290, 190, 338, 116, 118, 221, 866, 423, 1, 1002, 144, 293, 2472,
		185, 183, 1556, 1, 2185, 1, 149, 342, 1, 157, 6, 305, 174, 380, 3, 222, 134, 2, 160, 4, 156, 164, 732, 192, 1380, 144, 1, 194,
		135, 3, 3766, 1, 6, 1, 287, 1, 117, 109, 138, 130, 98, 112, 8, 1, 105, 166, 1253, 327, 9, 222, 516, 168, 1, 1, 7, 120, 336, 687,
		228, 191, 356, 2662, 5, 229, 434, 187, 1, 113, 162, 635, 434, 158, 520, 197, 1, 148, 6, 281, 225, 121, 2, 201, 121, 198, 107,
		1613, 120, 1, 2, 1236, 362, 4, 107, 7, 3100, 445, 2, 346, 137, 96, 6, 3, 119, 162, 131, 2, 105, 171, 1, 2, 30, 466, 119, 1073, 127,
		1, 294, 126, 296, 1, 128, 160, 373, 219, 114, 115, 122, 141, 277, 1, 174, 6, 348, 1610, 102, 2891, 268, 113, 1, 2, 158, 544, 37, 4,
		202, 1, 1862, 126, 1, 147, 2, 245, 315, 1, 1208, 395, 164, 504, 199, 129, 1, 145, 580, 1, 122, 105, 110, 213, 233, 173, 238, 2, 155,
		108, 3, 463, 129, 192, 216, 2, 172, 127, 238, 1513, 108, 165, 117, 206, 3, 100, 146, 1193, 3700, 1, 1, 100, 2, 1092, 373, 103, 442,
		111, 1, 4, 216, 255, 182, 158, 189, 2, 1, 3370, 3906, 870, 17, 1, 2461, 6, 219, 675, 1, 123, 206, 120, 1, 175, 352, 1, 115, 115, 188,
		347, 1, 1, 1, 332, 164, 292, 2, 238, 3, 130, 445, 378, 290, 1, 362, 146, 128, 477, 252, 10, 324, 175, 173, 252, 228, 410, 105, 3, 1111,
		199, 99, 1, 276, 171, 227, 2, 184, 2, 103, 155, 1, 249, 161, 1190, 2896, 1, 94, 3372, 124, 304, 182, 120, 401, 1, 3639, 368, 1, 105, 146,
		112, 105, 370, 130, 210, 6, 3, 214, 6, 1, 8, 120, 119, 2, 338, 3761, 503, 1645, 144, 1, 1, 151, 196, 291, 1, 113, 4, 417, 221, 2, 393, 644,
		124, 97, 145, 139, 912, 213, 123, 224, 245, 1, 293, 1042, 1, 104, 333, 254, 874, 1, 2, 119, 405, 1, 106, 613, 274, 1, 2, 212, 2, 245, 305,
		3, 112, 538, 818, 129, 390, 160, 281, 205, 131, 1, 339, 1358, 130, 117, 107, 1, 140, 105, 637, 210, 100, 410, 398, 175, 606, 249, 22,
		847, 2, 9, 348, 2, 293, 1, 178, 3, 6, 116, 157, 122, 295, 196, 447, 142, 4039, 296, 784, 899, 138, 1, 8372, 147, 1, 11, 762, 101, 5, 1,
		110, 2, 2675, 1, 1417, 4, 331, 2, 1, 285, 198, 174, 102, 19, 323, 371, 114, 119, 4, 28, 203, 141, 120, 1, 1, 114, 138, 2, 10158, 275, 38,
		282, 313, 132, 110, 534, 223, 1, 1213, 2507, 137, 224, 139, 745, 1, 143, 146, 4, 213, 119, 2, 1, 1, 221, 11, 3, 319, 166, 36, 125, 5, 17,
		143, 1788, 775, 178, 257, 111, 1, 295, 105, 1, 302, 781, 1, 1, 204, 158, 723, 4, 155, 1, 166, 205, 1, 132, 1, 410, 8, 1, 13, 128, 2, 194,
		114, 3, 327, 6, 1017, 511, 117, 147, 2, 345, 107, 5, 105, 110, 1, 147, 197, 428, 119, 119, 215, 116, 135, 193, 4, 1, 2, 150, 2, 150, 154,
		103, 108, 507, 137, 2, 1, 175, 1, 1, 107, 107, 107, 171, 184, 286, 258, 1046, 77, 7, 310, 139, 2, 897, 2, 192, 101, 146, 286, 103, 2076,
		124, 164, 1, 147}
	fmt.Scanf("%d", &t)
	var arr [600]int
	// var errors [900]int

	for i := 0; i < t; i++ {
		var money, choclateCount, wrappers int
		fmt.Scanf("%d %d %d", &money, &choclateCount, &wrappers)
		result := ClacChoclateCount(money, choclateCount, wrappers)
		// fmt.Println(result)
		arr[i] = result
	}
	fmt.Println("---------")

	for i := 0; i < 600; i++ {
		if Result[i] != arr[i] {
			fmt.Printf("Error in %d expectd: %d got: %d\n", i, Result[i], arr[i])
		}
	}
}

func ClacChoclateCount(money, price, wrappers int) int {
	by_money := money / price
	by_wrappers := by_money / wrappers
	total := by_money + by_wrappers
	remainder := by_money % wrappers
	by_wrappers = (remainder + by_wrappers) / wrappers
	total += by_wrappers
	return total
}

func strongPassword(password string) int {
	result := 0
	numbers := "0123456789"
	lower_case := "abcdefghijklmnopqrstuvwxyz"
	upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special_characters := "!@#$%^&*()-+"

	if !strings.ContainsAny(password, numbers) {
		result += 1
	}

	if !strings.ContainsAny(password, lower_case) {
		result += 1
	}

	if !strings.ContainsAny(password, upper_case) {
		result += 1
	}

	if !strings.ContainsAny(password, special_characters) {
		result += 1
	}
	if len(password) < 6 {
		if (result + len(password)) < 6 {
			result = (6 - len(password))
		}
	}
	return result
}

func testStrongPass() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	result := strongPassword(text)
	fmt.Println(result)
}
func test() {
	testStrongPass()
	testChoclate()
}
