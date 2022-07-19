package main

import "fmt"

func main() {
	fmt.Println("あなたは男？女？コミュニケーションタイプチェックシート")

	for {
		fmt.Println("\nこのまま診断を続けますか？？？（Y/n）")
		var q0 string
		fmt.Scanf("%s", &q0)

		if q0 == "Y" || q0 == "y" {
			fmt.Println("ありがとうございます！！それでははじめましょう！！")
			fmt.Println("これからの質問に A か B で入力してください")

			var countA int = 0
			var countB int = 0

			fmt.Println("Q1.仕事で褒められたときには？\nA：「すごい」と褒められたい\nB：「頑張ったね！」と言われたい")
			fmt.Scanf("%s", &q0)
			for {
				if q0 == "A" || q0 == "a" {
					countA = +1
					break
				} else if q0 == "B" || q0 == "b" {
					countB = +1
					break
				} else {
					fmt.Println("\n残念ながら入力頂いた文字は無効な文字ですOTZ")
					continue
				}
			}

			fmt.Println(countA)
			fmt.Println(countB)

			break
		} else if q0 == "N" || q0 == "n" {
			fmt.Println("\nやっていただけなくて残念！！また機会があれば是非！！")
			break
		} else {
			fmt.Println("\n残念ながら入力頂いた文字は無効な文字ですOTZ")
			continue
		}
	}
}
