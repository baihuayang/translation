defer:
一个基础但十分实用的实现代码执行追踪的方案就是在进入和离开某个函数打印相关的消息，即可以提炼为下面两个函数：
func trace(s string) { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }


defer 语句和匿名函数
关键字 defer （详见第 6.4 节）经常配合匿名函数使用，它可以用于改变函数的命名返回值。


闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
例子：https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/06.9.md


一个返回值为另一个函数的函数可以被称之为工厂函数，这在您需要创建一系列相似的函数的时候非常有用：书写一个工厂函数而不是针对每种情况都书写一个函数