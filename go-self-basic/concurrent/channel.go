package main
/**

使用内置的make函数，我们可以创建一个channel：
ch := make(chan int) // ch has type 'chan int'


发送和接收两个操作都是用<-运算符。在发送语句中，<-运算符分割channel和要发送的值。在接收语句中，<-运算符写在channel对象之前。一个不使用接收结果的接收操作也是合法的。

ch <- x  // a send statement
x = <-ch // a receive expression in an assignment statement
<-ch     // a receive statement; result is discarded

 */
