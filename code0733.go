/*
有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在 0 到 65535 之间。

给你一个坐标 (sr, sc) 表示图像渲染开始的像素值（行 ，列）和一个新的颜色值 newColor，让你重新上色这幅图像。

为了完成上色工作，从初始坐标开始，记录初始坐标的上下左右四个方向上像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对应四个方向上像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为新的颜色值。

最后返回经过上色渲染后的图像。

示例 1:

输入:
image = [[1,1,1],[1,1,0],[1,0,1]]
sr = 1, sc = 1, newColor = 2
输出: [[2,2,2],[2,2,0],[2,0,1]]
解析:
在图像的正中间，(坐标(sr,sc)=(1,1)),
在路径上所有符合条件的像素点的颜色都被更改成2。
注意，右下角的像素没有更改为2，
因为它不是在上下左右四个方向上与初始点相连的像素点。
注意:

image 和 image[0] 的长度在范围 [1, 50] 内。
给出的初始点将满足 0 <= sr < image.length 和 0 <= sc < image[0].length。
image[i][j] 和 newColor 表示的颜色值在范围 [0, 65535]内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flood-fill
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	m := len(image)
	if m == 0 || sr >= m {
		return image
	}
	n := len(image[0])
	if sc >= n {
		return image
	}
	if image[sr][sc] == newColor {
		return image
	}
	oldColor := image[sr][sc]

	queue := NewLinkedQueue()
	queue.Add([2]int{sr, sc})
	for queue.Size() != 0 {
		ele, _ := queue.Peek()
		queue.Remove()
		i, j := ele.([2]int)[0], ele.([2]int)[1]

		if i-1 >= 0 && image[i-1][j] == oldColor {
			queue.Add([2]int{i - 1, j})
		}
		if i+1 < m && image[i+1][j] == oldColor {
			queue.Add([2]int{i + 1, j})
		}
		if j-1 >= 0 && image[i][j-1] == oldColor {
			queue.Add([2]int{i, j - 1})
		}
		if j+1 < n && image[i][j+1] == oldColor {
			queue.Add([2]int{i, j + 1})
		}
		image[i][j] = newColor
	}

	return image
}
