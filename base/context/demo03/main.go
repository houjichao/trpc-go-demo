package main

import (
	"context"
	"fmt"
)

/*
type Context interface{

        Deadline() (deadline time.Time, ok bool)

        Done() <-chan struct{}

        Err() error

        Value(key any) any

}
重要接口：
Deadline() (deadline time.Time, ok bool) // 获取 deadline 时间，如果没有的话 ok 会返回 false
Done() <-chan struct{} // 返回的是一个 channel ，用来应用监听任务是否已经完成
Err() error // 返回取消原因 例如：Canceled\DeadlineExceeded
Value(key any) any // 根据指定的 key 获取是否存在其 value 有则返回

可以看到这个接口非常清晰简单明了，并且没有过多的 method，这也是 go 设计理念，接口尽量简单、小巧，通过组合来实现丰富的功能，后面会看到如何组合的。

再来看另一个接口 canceler ，这是一个取消接口，其中一个非导出 method cancel，接收一个 bool 和一个 error ，
bool 用来决定是否将其从父 Context 中移除， err 用来标明被取消的原因。
还有个 Done() 和 Context 接口一样，这个接口为何这么设计，后面再揭晓。
type canceler interface{
	cancel(removeFromParent bool, err error)
	Done() <-chan struct{}
}
接下来看这两个接口的实现者都有谁，首先 Context 直接实现者有 *valueCtx（比较简单放最后讲） 和 *emptyCtx

而 canceler 直接实现者有 *cancelCtx 和 *timerCtx ，并且这两个同时也实现了 Context 接口
（记住我前面说得另外两个是直接实现，这俩是嵌套接口实现松耦合，后面再说具体好处），下面逐一讲解每个实现。

canceler是不可导出的，外部不能直接操作canceler类型对象，只能通过func()操作。
*/

//-------------
//在使用 WithValue 时，包应该将键定义为未导出的类型以避免发生碰撞，这里贴个官网的例子：
// package user 这里为了演示直接在 main 包定义
// User 是存储在 Context 值
type User struct {
	Name string
	Age  int
}

// key 是非导出的，可以防止碰撞
type key int

// userKey 是存储 User 类型的键值，也是非导出的。
var userKey key

// NewContext 创建一个新的 Context，携带 *User
func NewContext(ctx context.Context, u *User) context.Context {
	return context.WithValue(ctx, userKey, u)
}

// FromContext 返回存储在 ctx 中的 *User
func FromContext(ctx context.Context) (*User, bool) {
	u, ok := ctx.Value(userKey).(*User)
	return u, ok
}


/**
那怎么能够防止碰撞呢？可以做个示例,看最后输出。

我们在第一行就用 userKey  的值 0，存储了一个值 “a"

然后再利用 NewContext 存储了 &User，底层实际用的是 context.WithValue(ctx, userKey, u)

读取时用的是 FromContext ，两次存储即使底层的 key 值都为 0， 但是互不影响，这是为什么呢？

还记得 WithValue 怎么实现的么？你每调用一次都会包一层，并且一层一层解析，
而且它会比较 c.key == key ，这里记住 go 的 == 比较是比较值和类型的，二者都相等才为 true
，而我们使用 type key int 所以 userKey 与 0 底层值虽然一样，但是类型已经不一样了
（这里就是 main.userKey 与 0），所以外部无论定义何种类型都无法影响包内的类型。这也是容易令人迷惑的地方
 */
func main() {
	ctx := context.WithValue(context.Background(), 0, "a")
	ctx = NewContext(ctx, &User{})
	v, _ := FromContext(ctx)
	fmt.Println(ctx.Value(0), v) // Output: a, &{ 0}
}
