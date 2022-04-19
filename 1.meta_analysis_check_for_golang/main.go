package main

import "fmt"


func main() {
	fmt.Println("メタ認知分析テスト！！")


	for {
		fmt.Println("\nこのまま診断を続けますか？？？（Y/n）")
		var q0 string
		fmt.Scanf("%s", &q0)

		if q0 == "Y" || q0 == "y" {
			fmt.Println("ありがとうございます！！それでははじめましょう！！")
			fmt.Println("これからの質問に 1〜4 で入力してください")
			fmt.Println("\n「1」…「当てはまらない」\n「2」…「少しだけ当てはまる」\n「3」…「やや当てはまる」\n「4」…「非常に当てはまる」")

			var q1, q2, q3, q4, q5, q6, q7, q8, q9, q10, q11, q12, q13, q14, q15, q16, q17, q18, q19, q20, q21, q22, q23, q24, q25, q26, q27, q28, q29, q30 int
			fmt.Println("\n1. 自分の記憶を信用していない")
			fmt.Scanf("%d", &q1)
			fmt.Println("\n2. 決心する前に深く考えることができる")
			fmt.Scanf("%d", &q2)
			fmt.Println("\n3. 常に自分の考えていることを意識している")
			fmt.Scanf("%d", &q3)
			fmt.Println("\n4. 一度気になり始めると頭から離れなくなる")
			fmt.Scanf("%d", &q4)
			fmt.Println("\n5. 思考を制御できなければ作業ができない")
			fmt.Scanf("%d", &q5)
			fmt.Println("\n6. 記憶力がよくない")
			fmt.Scanf("%d", &q6)
			fmt.Println("\n7. 深く考えることでうまく取り組むことができる")
			fmt.Scanf("%d", &q7)
			fmt.Println("\n8. 自分の心の動きに深く注意を払っている")
			fmt.Scanf("%d", &q8)
			fmt.Println("\n9. 心配し始めると止まらなくなる")
			fmt.Scanf("%d", &q9)
			fmt.Println("\n10. 自分の思考を制御できないのは弱さの現れだ")
			fmt.Scanf("%d", &q10)
			fmt.Println("\n11. 自分がしたことの記憶に自信がない")
			fmt.Scanf("%d", &q11)
			fmt.Println("\n12. 深く考えることで仕事を円滑に行うことができる")
			fmt.Scanf("%d", &q12)
			fmt.Println("\n13. 自分の思考についてよく考えている")
			fmt.Scanf("%d", &q13)
			fmt.Println("\n14. 心配事で病気になるかもしれない")
			fmt.Scanf("%d", &q14)
			fmt.Println("\n15. いかなるときも思考を制御しておくべきだ")
			fmt.Scanf("%d", &q15)
			fmt.Println("\n16. 場所の記憶に自信がない")
			fmt.Scanf("%d", &q16)
			fmt.Println("\n17. 深く考えることで問題解決に役立つ")
			fmt.Scanf("%d", &q17)
			fmt.Println("\n18. 常に自分の思考について審査する")
			fmt.Scanf("%d", &q18)
			fmt.Println("\n19. 心配事があると無視できない")
			fmt.Scanf("%d", &q19)
			fmt.Println("\n20. 思考の中には考えてはいけないものがある")
			fmt.Scanf("%d", &q20)
			fmt.Println("\n21. 言葉や名前についての記憶に自信がない")
			fmt.Scanf("%d", &q21)
			fmt.Println("\n22. 頭の中をきちんとしておくには深く考える必要がある")
			fmt.Scanf("%d", &q22)
			fmt.Println("\n23. 自分の思考を監視している")
			fmt.Scanf("%d", &q23)
			fmt.Println("\n24. 心配事で気が狂うかもしれない")
			fmt.Scanf("%d", &q24)
			fmt.Println("\n25. 心配事を制御できないのは自分の責任である")
			fmt.Scanf("%d", &q25)
			fmt.Println("\n26. 時々記憶違いで失敗することがある")
			fmt.Scanf("%d", &q26)
			fmt.Println("\n27. 深く考えることで問題を事前に避けるのに役立つ")
			fmt.Scanf("%d", &q27)
			fmt.Println("\n28. 問題を考えているとき自分の思考傾向を意識している")
			fmt.Scanf("%d", &q28)
			fmt.Println("\n29. 心配事は脅威となる")
			fmt.Scanf("%d", &q29)
			fmt.Println("\n30. 制御しなければ罰せられる思考がある")
			fmt.Scanf("%d", &q30)

			//合計点計測
			var a1 = q1 + q6 + q11 + q16 + q21 + q26
			var a2 = q2 + q7 + q12 + q17 + q22 + q27
			var a3 = q3 + q8 + q13 + q18 + q23 + q28
			var a4 = q4 + q9 + q14 + q19 + q24 + q29
			var a5 = q5 + q10 + q15 + q20 + q25 + q30

			fmt.Println("\n結果発表！！！！")
			fmt.Println("\n=============================================================================")
			fmt.Println("\n「認知的自信の欠如」の合計点：")
			fmt.Println(a1)
			fmt.Println("\n「心配事への積極的信念」の合計点：")
			fmt.Println(a2)
			fmt.Println("\n「認知的自己意識」の合計点：")
			fmt.Println(a3)
			fmt.Println("\n「思考制御不能と危険への消極的信念」の合計点：")
			fmt.Println(a4)
			fmt.Println("\n「迷信・罰・責任など思考一般への制御欲求の消極的信念」の合計点：")
			fmt.Println(a5)
			fmt.Println("\n=============================================================================")
			break
		} else if q0 == "N" || q0 == "n" {
			fmt.Println("\nやっていただけなくて残念！！また機会があれば是非！！")
			break
		} else{
			fmt.Println("\n残念ながら入力頂いた文字は無効な文字ですOTZ")
			continue
		}
	}
}

