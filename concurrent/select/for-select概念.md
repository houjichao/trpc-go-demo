1. 定义
先理解三个概念：

没有条件的for是死循环：

for { 
    //code here 
}
select-case 是专门用来轮询通道channel是否有值传递过来，有多个的时候会随机选择一个
符合条件的case，如果一个也没有且没有 default 选项，它将成为一个阻塞操作，
直到有新的值符合case为止, 有点像WaitGroup

select { 
    case <-channelA :
    case b := <-channelB :
       //oper b
}
通道
channel <- 10 //发送数据10
num := <- channel  // 接收数据