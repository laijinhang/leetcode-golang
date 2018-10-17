/*
题目描述
请编写一个程序，按升序对栈进行排序（即最大元素位于栈顶），要求最多只能使用一个额外的栈存放临时数据，但不得将元素复制到别的数据结构中。

给定一个int[] numbers(C++中为vector&ltint>)，其中第一个元素为栈顶，请返回排序后的栈。请注意这是一个栈，意味着排序过程中你只能访问到最后一个元素。

测试样例：
[1,2,3,4,5]
返回：[5,4,3,2,1]


思路来源：[https://www.nowcoder.com/questionTerminal/d0d0cddc1489476da6b782a6301e7dec](https://www.nowcoder.com/questionTerminal/d0d0cddc1489476da6b782a6301e7dec)

思路：
**1\.** 先看下图
![图片说明](http://upload-images.jianshu.io/upload_images/5022741-254576f871c6be16?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240 "图片标题")

初始栈initStack中存放的数组中待排序的数；临时栈tempStack中存放的是已经排好序的数。
现在继续对初始栈中的数进行排序，5应当插入到临时栈哪个位置？

**2.** 5应该插入到8下，3上。
具体如何操作呢？
首先初始栈initStack弹出待排序的数5，存入变量tmp；而临时栈tempStack弹出比5大的数，并存入初始化栈initStack中。如下图：
![图片说明](http://upload-images.jianshu.io/upload_images/5022741-e1036e40eb797875?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240 "图片标题")

**3.** 将变量tmp保存的数插入到临时栈tempStack中去，由于初始化栈initStack中8，12是排好序的，可以再直接弹入临时栈中，再对下一个数10进行如上操作。
![图片说明](http://upload-images.jianshu.io/upload_images/5022741-d803d0b40980ba69?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240 "图片标题")
 */
package main

import "fmt"

type TwoStatcks struct {}

func (TwoStatcks) twoStackSort(numbers []int) []int {
	len1 := len(numbers)

	tempStack := make([]int, 0, len1)
	for i := 0;i < len1;i++ {
		temp := numbers[i]
		j := i
		for len(tempStack) > 0 && tempStack[len(tempStack)-1] > temp{
			j--
			numbers[j] = tempStack[len(tempStack)-1]
			tempStack = tempStack[:len(tempStack)-1]
		}
		tempStack = append(tempStack, temp)
		for j < i {
			tempStack = append(tempStack, numbers[j])
			j++
		}
	}
	return tempStack
}

func main() {
	var s TwoStatcks
	statck := s.twoStackSort([]int{5,4,3,1,5,7,102,32,3,1,1,2,3})
	fmt.Print(statck)
}
