package jiaohu

import (
	"fmt"
	"testing"
)

func TestMsgGoodb(t *testing.T) {
	var busdBidsZu [10][2]float64

	var usdtBidsZu [20][2]float64

	var BidsZu [20][3]float64 //绿下 从高到低
	busdBidsZu[0][0] = 2.40
	busdBidsZu[0][1] = 12
	busdBidsZu[1][0] = 2.39
	busdBidsZu[1][1] = 13
	busdBidsZu[2][0] = 2.38
	busdBidsZu[2][1] = 14
	busdBidsZu[3][0] = 2.35
	busdBidsZu[3][1] = 15
	busdBidsZu[4][0] = 2.33
	busdBidsZu[4][1] = 16
	busdBidsZu[5][0] = 2.32
	busdBidsZu[5][1] = 17
	busdBidsZu[6][0] = 2.30
	busdBidsZu[6][1] = 18
	busdBidsZu[7][0] = 2.24
	busdBidsZu[7][1] = 19
	busdBidsZu[8][0] = 2.22
	busdBidsZu[8][1] = 20
	busdBidsZu[9][0] = 2.10
	busdBidsZu[9][1] = 20

	usdtBidsZu[0][0] = 2.38
	usdtBidsZu[0][1] = 12
	usdtBidsZu[1][0] = 2.37
	usdtBidsZu[1][1] = 13
	usdtBidsZu[2][0] = 2.36
	usdtBidsZu[2][1] = 14
	usdtBidsZu[3][0] = 2.32
	usdtBidsZu[3][1] = 15
	usdtBidsZu[4][0] = 2.29
	usdtBidsZu[4][1] = 16
	usdtBidsZu[5][0] = 2.28
	usdtBidsZu[5][1] = 17
	usdtBidsZu[6][0] = 2.27
	usdtBidsZu[6][1] = 18
	usdtBidsZu[7][0] = 2.24
	usdtBidsZu[7][1] = 19
	usdtBidsZu[8][0] = 2.23
	usdtBidsZu[8][1] = 20
	usdtBidsZu[9][0] = 2.09
	usdtBidsZu[9][1] = 20

	startB := 0
	startU := 0
	tolAmount := 0.0
	for iii := 0; iii < 10; iii++ {
		if busdBidsZu[startB][0] > usdtBidsZu[startU][0] {
			//busd是最大的
			BidsZu[iii][0] = busdBidsZu[startB][0] //价格 定为busd
			BidsZu[iii][1] = busdBidsZu[startB][1] //数量 定为busd
			//总量
			tolAmount += BidsZu[iii][1]
			//busd递增
			startB++
		} else if busdBidsZu[startB][0] < usdtBidsZu[startU][0] {
			//usdt是最大的
			BidsZu[iii][0] = usdtBidsZu[startU][0] //价格 定为usdt
			BidsZu[iii][1] = usdtBidsZu[startU][1] //数量 定为usdt
			//总量
			tolAmount += BidsZu[iii][1]
			//busd递增
			startU++
		} else {
			//两个最大的相等 总单价等于or 数量总B+U
			BidsZu[iii][0] = busdBidsZu[startB][0]
			BidsZu[iii][1] = busdBidsZu[startB][1] + usdtBidsZu[startU][1]
			//总量
			tolAmount += BidsZu[iii][1]
			//两个都查询下个
			startB++
			startU++
		}
		BidsZu[iii][2] = tolAmount
		fmt.Println(iii, BidsZu[iii][0], BidsZu[iii][1], BidsZu[iii][2])
	}
}
