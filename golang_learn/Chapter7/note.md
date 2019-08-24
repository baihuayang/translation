如果你有一个函数需要对数组做操作，你可能总是需要把参数声明为切片。当你调用该函数时，把数组分片，创建为一个 切片引用并传递给该函数

new 函数分配内存，make 函数初始化

类型 []byte 的切片十分常见，Go 语言有一个 bytes 包专门用来解决这种类型的操作方法。

bytes包 通过 buffer 串联字符串 类似于 Java 的 StringBuilder 类。

append方法

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}


假设 s 是一个字符串（本质上是一个字节数组），那么就可以直接通过 c := []byte(s) 来获取一个字节的切片 c。另外，您还可以通过 copy 函数来达到相同的目的：copy(dst []byte, src string)。