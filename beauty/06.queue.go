package beauty

/*
queue
	先进者先出
	操作受限的线性表数据结构，只允许队头入队 enqueue() 和队尾出队 dequeue()
场景：在很多偏底层系统、框架、中间件的开发中，起着关键性的作用
	高性能队列 Disruptor、Linux 环形缓存，都用到了循环并发队列
	Java concurrent 并发包利用 ArrayBlockingQueue 来实现公平锁

实现
	顺序队列：用数组实现的队列，可以实现有界队列
	链式队列：用链表实现的队列，可以实现无界队列
顺序队列
	数据搬移
		在出队时可以不用搬移数据
		如果没有空闲空间了，只需要在入队时，再集中触发一次数据的搬移操作
	解决数据搬移
		循环队列
链式队列
循环队列
	队空：head == tail
	队满：(tail+1)%n == head
	循环队列会浪费一个数组的存储空间
阻塞队列和并发队列
	阻塞队列 vs “生产者 - 消费者模型”
		在队列为空的时候，从队头取数据会被阻塞。因为此时还没有数据可取，直到队列中有了数据才能返回
		如果队列已经满了，那么插入数据的操作就会被阻塞，直到队列中有空闲位置后再插入数据，然后再返回
	并发队列
		最简单直接的实现方式是直接在 enqueue()、dequeue() 方法上加锁
		但是锁粒度大并发度会比较低，同一时刻仅允许一个存或者取操作
	Disruptor
		基于数组的循环队列，利用 CAS 原子操作，可以实现非常高效的并发队列
		详见 Disruptor	// TODO

线程池没有空闲线程时，新的任务请求线程资源时，线程池该如何处理？
	非阻塞处理：直接拒绝任务请求
	阻塞处理：将请求排队，等到有空闲线程时，取出排队的请求继续处理。那如何存储排队的请求呢？
		链式队列
			实现一个支持无限排队的无界队列（unbounded queue），但是可能会导致过多的请求排队等待，请求处理的响应时间过长
			所以，针对响应时间比较敏感的系统，基于链表实现的无限排队的线程池是不合适的
		顺序队列
			实现有界队列（bounded queue），队列的大小有限，所以线程池中排队的请求超过队列大小时，接下来的请求就会被拒绝
			这种方式对响应时间敏感的系统来说，就相对更加合理
			合理的队列大小：队列太大导致等待的请求太多，队列太小会导致无法充分利用系统资源、发挥最大性能
其他场景
	应用在任何有限资源池中，用于排队请求，比如数据库连接池等
	对于大部分资源有限的场景，当没有空闲资源时，基本上都可以通过“队列”这种数据结构来实现请求排队

思考
	有哪些类似的(线程)池结构或者场景中会用到队列的排队请求呢？
	如何实现无锁并发队列，网上有非常多的讨论。对这个问题，你怎么看呢？
*/

// CircularQueue 循环队列
type CircularQueue struct {
	arr        []int
	head, tail int
	n          int
}

func New(n int) CircularQueue {
	return CircularQueue{arr: make([]int, n+1), n: n + 1}
}
func (cq *CircularQueue) Enqueue(v int) bool {
	next := (cq.tail + 1) % cq.n
	if next == cq.head {
		return false
	}
	cq.arr[cq.tail], cq.tail = v, next
	return true
}
func (cq *CircularQueue) Dequeue() (int, bool) {
	if cq.head == cq.tail {
		return 0, false
	}
	v := cq.arr[cq.head]
	cq.head = (cq.head + 1) % cq.n
	return v, true
}
