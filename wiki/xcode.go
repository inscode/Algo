package main

import "fmt"

//yiw
func main() {
	var a = [][]int{
		{0, 1, 2, 3},   /* 第一行索引为 0 */
		{4, 5, 6, 7},   /* 第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
		{18, 19, 20, 21}, /* 第三行索引为 2 */
	}
	fmt.Printf("%p\n",(&a))
	fmt.Println()

	fmt.Printf("%p---base\n",&a[0]) //c060  //6*16+0 = 96  = 96+8*0【0 = 0*4+0】
	fmt.Println(&a[0][0]) //c060  //6*16+0 = 96  = 96+8*0【0 = 0*4+0】
	fmt.Println(&a[0][1]) //c068  //6*16+8 = 104 = 96+8*1【1 = 0*4+1】
	fmt.Println(&a[0][2]) //c070  //7*16+0 = 112 = 96+8*2【2 = 0*4+2】
	fmt.Println(&a[0][3]) //c078  //7*16+8 = 120 = 96+8*3【3 = 0*4+3】
	fmt.Println()

	fmt.Printf("%p---base\n",&a[1]) //c060  //6*16+0 = 96  = 96+8*0【0 = 0*4+0】
	fmt.Println(&a[1][0]) //c080  //8*16+0 = 128 = 96+8*4【4 = 1*4+0】
	fmt.Println(&a[1][1]) //c088  //8*16+8 = 136 = 96+8*5【5 = 1*4+1】
	fmt.Println(&a[1][2]) //c090  //9*16+0 = 144 = 96+8*6【6 = 1*4+2】
	fmt.Println(&a[1][3]) //c098  //9*16+8 = 152 = 96+8*7【7 = 1*4+3】
	fmt.Println()

	fmt.Printf("%p---base\n",&a[2]) //c060  //6*16+0 = 96  = 96+8*0【0 = 0*4+0】
	fmt.Println(&a[2][0]) //c0a0  //10*16+0 = 160 = 96+8*8
	fmt.Println(&a[2][1]) //c0a8  //10*16+8 = 168 = 96+8*9
	fmt.Println(&a[2][2]) //c0b0  //11*16+0 = 176 = 96+8*10
	fmt.Println(&a[2][3]) //c0b8  //11*16+8 = 184 = 96+8*11

	fmt.Println()

	fmt.Printf("%p,%s\n", &a[3], "---base_address")
	fmt.Println(&a[3][0]) //c0a0  //10*16+0 = 160 = 96+8*8
	fmt.Println(&a[3][1]) //c0a8  //10*16+8 = 168 = 96+8*9
	fmt.Println(&a[3][2]) //c0b0  //11*16+0 = 176 = 96+8*10
	fmt.Println(&a[3][3]) //c0b8  //11*16+8 = 184 = 96+8*11

	//a[m][n]_address = base_address + (i*n + j) * data_type_size
	fmt.Println()
}
