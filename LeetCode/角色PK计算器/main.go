package main

func PK(Ax int, Ay int, Az int, Aw int, Bx int, By int, Bz int, Bw int) int {
	count1 := 0
	count2 := 0
	Blood1 := Ax
	Blood2 := Bx
	for {
		// 回复血量
		if Blood1+Aw > Ax {
			Blood1 = Ax
		} else {
			Blood1 += Aw
		}
		if Blood2+Bw > Bx {
			Blood2 = Bx
		} else {
			Blood2 += Bw
		}
		// 开始进行攻击
		if count1 == 0 {
			Blood2 -= Ay
		} else {
			count1--
		}
		if count2 == 0 {
			Blood1 -= By
		} else {
			count2--
		}
		// 攻击完成之后判断是否死亡
		if Blood1 > 0 && Blood2 <= 0 {
			return 1
		} else if Blood2 > 0 && Blood1 <= 0 {
			return 2
		} else if Blood1 <= 0 && Blood2 <= 0 {
			return 3
		}
		if Ay <= Bw && By <= Aw {
			return 4
		}
	}
}

func main() {

}
