package main

import "fmt"

type patient struct {
	name string
	height int
	weight int
	chest int
	upb int
	hbp int
	lbp int
}

// patientのインスタンスを戻り値として返す
func newpatient(name string) patient {
	return patient{name, 0, 0, 0, 0, 0 ,0}
}

// 身長, 体重
// patientインスタンスのbodysメソッドとして使用可能
// こう書いてあるとき実際にはどこにアクセスしているのか？ -> インスタンス
// => patient{"HMさん", 0, 0, 0, 0, 0 ,0}
func (pat patient) bodys(height int, weight int) patient {
	pat.height = height
	pat.weight = weight
	return pat
	// patient{"HMさん", 170, 80, 0, 0, 0, 0 ,0}

}

// ふるまいを追加する
func (pat patient) bloodp(hbp int, lbp int) patient {
	// patientインスタンスの値を更新する
	pat.hbp = hbp
	pat.lbp = lbp
	fmt.Println(pat)
	return pat
}

func (pat patient) bmi() string {
	if pat.weight * pat.height <= 0 {
		return "測定エラー"
	}
	bmi := (pat.weight * 10000) / (pat.height * pat.height)
	if bmi < 19 {
		return "もっと食べましょう"
	}
	return "バランスの取れた体格です"
}

func (pat patient) bprange() string {
	if pat.hbp * pat.lbp <= 0 {
		return "測定エラー"
	}
	if pat.hbp < 100 {
		return "血圧が低いかもしれません"
	}
	if pat.hbp < 140 && pat.lbp < 190 {
		return "血圧は正常です"
	}
	return "血圧に注意が必要です"
}

func (pat patient) report() string {
	report := pat.name + "さん\n"
	report += pat.bmi() + "\n"
	report += pat.bprange()
	return report
	
}

func main() {
	// patientインスタンスの作成を行う, スライスで作成
	patients := []patient{
		newpatient("Sho").bodys(169, 58).bloodp(120, 80),
		newpatient("Yua").bodys(166, 45).bloodp(115, 75),
	}

	fmt.Println("***健康診断の結果です***")

	for i:=0; i < len(patients); i ++ {
		fmt.Println(patients[i].report())
		fmt.Println()
	}
}