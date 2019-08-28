协程讲解：https://www.cnblogs.com/liang1101/p/7285955.html

https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.2.md
精彩的质数打印算法

基本思路 
大的for循环，第一次的两个管道负责过滤2的倍数(非质数)
第二次3的倍数，第三次5的倍数

其中 3 5 。。。 都是之前算出来的