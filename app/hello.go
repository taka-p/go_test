/*
 go build hello.go -> .hello/
 go run hello.go -> runtime compile
*/

/*
 変数名
 1文字目が小文字の場合、パッケージ内限定
 1文字目が大文字の場合、パッケージ外でも利用出来る
*/

package main

import (
	"fmt"
	//"time"
	"net/http"
)

//func main()  {
//	//var msg string
//	// 型宣言と値定義の省略
//	msg := "Hello World"
//	fmt.Println(msg)
//
//	//var a, b int
//	// 型宣言と値定義の省略
//	a, b := 10, 15
//
//	// 異なる型をまとめて宣言
//	var (
//		c int
//		d string
//	)
//	c = 20
//	d = "hoge"
//}

/*
 基本的な型
 string: "hoge"
 int: 10
 float64: 10.1
 bool: true/false
 nil

 初期値
 var s string -> ""
 var a int -> 0
 var f bool -> false
*/

//func main()  {
//	a := 10
//	b := 12.3
//	c := "hoge"
//	var d bool
//	/*
//	書式付き
//	%d: 整数値
//	%f: float
//	%s: 文字列
//	%t: 真偽値
//	*/
//	fmt.Printf("a: %d, b:%f, c:%s, d:%t\n", a, b, c, d)
//}

/*
 定数
 const
*/

//func main()  {
//	//const name ="taguchi"
//	//name = "hoge"
//	//fmt.Println(name)
//
//	const (
//		sun = iota // 0
//		mon // 1
//		tue // 2
//	)
//	fmt.Println(sun, mon, tue)
//}


/*
 ポインタ
 アドレスを指し示す変数
 演算はできない

*/

//func main()  {
//	a:= 5
//
//	var pa *int
//	pa = &a // &a = aのアドレス
//	// paの領域にあるデータの値 = *pa
//
//	fmt.Println(pa)
//	fmt.Println(*pa)
//}


/*
 関数と返り値1
*/

//func hi(name string) string {
//	//fmt.Println("hi! " + name)
//	msg := "hi! " + name
//	return msg
//}

// 名前付き返り値
//func hi(name string) (msg string) {
//	//fmt.Println("hi! " + name)
//	msg = "hi! " + name
//	return
//}
//
//func main()  {
//	//hi("taguchi")
//	fmt.Println(hi("taguchi"))
//}


/*
 関数と返り値2
 複数の返り値を持てる
*/

//func swap(a, b int) (int, int) {
//	return b, a
//}

//func main()  {
//	//fmt.Println(swap(2, 5)) // -> 5 2
//
//	// 変数の代入（第一級関数）
//	//f := func(a, b int) (int, int) {
//	//	return b, a
//	//}
//	//fmt.Println(f(2, 3)) // -> 3 2
//
//	// 即時関数
//	func(msg string) {
//		fmt.Println(msg)
//	}("taguchi")
//}


/*
 配列
*/

//func main() {
//	//var a [5]int // -> a[0] - a[4]
//	//a[2] = 2
//	//fmt.Println(a) // -> [0 0 2 0 0]
//
//	//b := [3]int{1, 3, 5}
//	b := [...]int{1, 3, 5}
//	fmt.Println(b) // -> [1 3 5]
//
//	// length
//	fmt.Println(len(b)) // -> 3
//}


/*
 スライス
*/

//func main()  {
//	a := [5]int{2, 10, 8, 15, 4}
//	s := a[2:4] // 8 15
//	s[1] = 12 // 配列への参照のため参照もとも変更
//
//	fmt.Println(a) // -> [2 10 8 12 4]
//	fmt.Println(s) // -> [8 12]
//	fmt.Println(len(s)) // -> 2
//	fmt.Println(cap(s)) // -> 3 (切り出しうる最大個数)
//}

//func main()  {
//	//s := make([]int, 3) // [0 0 0]
//	s := []int{1, 3, 5} // [1 3 5]
//
//	// append
//	s = append(s, 8, 2, 10) // [1 3 5 8 2 10]
//
//	// copy
//	t := make([]int, len(s)) // [1 3 5 8 2 10]
//	n := copy(t, s) // 6 (要素数を返す)
//
//	fmt.Println(s)
//	fmt.Println(t)
//	fmt.Println(n)
//}


/*
 map
 */

//func main()  {
//	//m := make(map[string]int)
//	//m["taguchi"] = 200
//
//	// 宣言と同時に値の定義
//	m := map[string]int{"taguchi":200} // map[taguchi:200]
//
//	fmt.Println(m)
//	fmt.Println(len(m))
//
//	// キーを削除
//	//delete(m, "taguchi")
//
//	// 存在チェック
//	v, ok := m["taguchi"]
//	fmt.Println(v, ok) // -> 200 true
//}
//


/*
 if
*/

//func main()  {
//	//score := 83
//
//	// ブロック内で有効な変数定義
//	if score := 83; score > 80 {
//		fmt.Println("great")
//	} else if score > 60 {
//		fmt.Println("good")
//	} else {
//		fmt.Println("soso")
//	}
//}
//


/*
 switch
*/

//func main()  {
//	//signal := "red"
//	//
//	//switch signal {
//	//case "red":
//	//	fmt.Println("Stop")
//	//case "blue", "green":
//	//	fmt.Println("Go")
//	//default:
//	//	fmt.Println("wrong signal")
//	//}
//
//
//	score := 82
//
//	switch {
//	case score > 80:
//		fmt.Println("great")
//	default:
//		fmt.Println("soso..")
//	}
//}


/*
 for
*/

//func main()  {
//
//	//for i := 0; i < 10; i++ {
//	//	//if i == 3 { break }
//	//	if i == 3 { continue }
//	//	fmt.Println(i) // -> 0 ~ 9
//	//}
//
//	//i := 0
//	//for i < 10 {
//	//	fmt.Println(i)
//	//	i++
//	//}
//
//	//i := 0
//	//for {
//	//	fmt.Println(i)
//	//	i++
//	//	if i == 3 { break }
//	//}
//}


/*
 range
 イテレーター
*/

//func main()  {
//	// i:index v:value
//	//s := []int{2, 3, 8}
//	//for i, v := range s {
//	//	fmt.Println(i, v)
//	//}
//
//	// ブランク修飾子
//	//s := []int{2, 3, 8}
//	//for _, v := range s {
//	//	fmt.Println(v)
//	//}
//
//	m := map[string]int{"taguchi":200, "fkoji":100}
//	for k, v := range m {
//		fmt.Println(k, v)
//	}
//}


/*
 構造体
 カスタムキャスト作成
*/

//type user struct {
//	name string
//	score int
//}
//
//func main()  {
//	//u := new(user)
//	//u.name = "taguchi"
//	//u.score = 20
//
//	u := user{name:"taguchi", score:20}
//	fmt.Println(u)
//}


/*
 メソッド
 データ型に紐づいた関数
*/

//type user struct {
//	name string
//	score int
//}
//
//// レシーバーは値渡し
//func (u user) show()  {
//	fmt.Println("name: %s, score: %d\n", u.name, u.score)
//}
//
//// レシーバーは参照渡し
//func (u *user) hit()  {
//	u.score++
//}
//
//func main()  {
//	u := user{name:"taguchi", score:20}
//
//	u.hit()
//	u.show()
//}


/*
 インターフェース
 メソッドの一覧を定義したデータ型
 異なる構造体でも同じインターフェースを満たす型として処理出来る
*/

//type greeter interface {
//	greet()
//}
//
//type japanese struct{}
//type american struct{}
//
//func (j japanese) greet()  {
//	fmt.Println("こんにちは")
//}
//
//func (a american) greet()  {
//	fmt.Println("Hello")
//}
//
//func main() {
//	greeters := []greeter{japanese{}, american{}}
//
//	for _, greeter := range greeters {
//		greeter.greet()
//	}
//}


/*
 空のインターフェース
*/

//type greeter interface {
//	greet()
//}
//
//type japanese struct{}
//type american struct{}
//
//func (j japanese) greet()  {
//	fmt.Println("こんにちは")
//}
//
//func (a american) greet()  {
//	fmt.Println("Hello")
//}
//
//func show(t interface{})  {
//	// 型アサーション
//	//_, ok := t.(japanese)
//	//if ok {
//	//	fmt.Println("iam japanese")
//	//} else {
//	//	fmt.Println("iam not japanese")
//	//}
//
//	// 型スイッチ
//	switch t.(type) {
//	case japanese:
//		fmt.Println("i am japanese")
//	default:
//		fmt.Println("i am not jappanese")
//	}
//}
//
//func main() {
//	greeters := []greeter{japanese{}, american{}}
//
//	for _, greeter := range greeters {
//		greeter.greet()
//		show(greeter)
//	}
//}
//


/*
 goroutine: 並行処理
*/

//func task1()  {
//	time.Sleep(time.Second * 2)
//	fmt.Println("taks1 finished")
//}
//
//func task2()  {
//	fmt.Println("taks2 finished")
//}
//
//func main()  {
//	go task1()
//	go task2()
//
//	time.Sleep(time.Second * 3)
//}
//


/*
 チャネル: データの受け渡しをするパイプ
*/

//func task1(result chan string)  {
//	time.Sleep(time.Second * 2)
//	//fmt.Println("taks1 finished")
//	result<- "task1 result"
//}
//
//func task2()  {
//	fmt.Println("taks2 finished")
//}
//
//func main()  {
//	result := make(chan string)
//
//	// バックグラウンド
//	go task1(result)
//	go task2()
//
//	// メインルーチン
//	// <-resultを待機
//	fmt.Println(<-result)
//
//	time.Sleep(time.Second * 3)
//}
//


/*
 package
 web server
*/

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
}

func main()  {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

