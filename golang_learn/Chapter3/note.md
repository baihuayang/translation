没有很好的调试器

gofmt     构建之前执行，重构代码风格
go build
go install


Go 语言与 C 语言的性能差距大概在 10%~20% 之间 (由于出版时间限制，该数据应为 2013 年 3 月 28 日之前产生)
保守估计在相同的环境和执行目标的情况下，Go 程序比 Java 或 Scala 应用程序要快上 2 倍，并比这两门语言占用的内存降低了 70%

但是
其实比较多门语言之间的性能是一种非常猥琐的行为，因为任何一种语言都有其所擅长和薄弱的方面

可以调用C库