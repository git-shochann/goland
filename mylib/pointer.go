package mylib

// 値渡し
func CallByValue(a, b int) (int, int){
	a += 1
	b += 2
	return a , b
}

// 参照渡し
// ポインタ型変数を受け取り, 戻り値でポインタ型変数を返す
func CallByReference(c *int) {
	*c += 5000 // 中身を変える
}