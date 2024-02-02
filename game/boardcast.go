package game

type BroadcastService[T any] struct {
	// 输入通道
	in chan T
	// 转发给这些channel
	listeners []chan T
	// 添加消费者
	addList chan (chan T)
	// 移除消费者
	removeList chan (chan T)
}

// public

// NewBroadcastService 创建一个广播服务
func NewBroadcastService[T any]() *BroadcastService[T] {
	return &BroadcastService[T]{
		in:         make(chan T),
		listeners:  make([]chan T, 3),
		addList:    make(chan (chan T)),
		removeList: make(chan (chan T)),
	}
}

// Listener 这会创建一个新消费者并返回一个监听通道
func (bs *BroadcastService[T]) Listener() chan T {
	ch := make(chan T)
	bs.addList <- ch
	return ch
}

// UnListener 移除一个消费者
func (bs *BroadcastService[T]) UnListener(ch chan T) {
	bs.removeList <- ch
}
func (bs *BroadcastService[T]) Run() chan T {
	go func() {
		for {
			// 处理新建消费者或者移除消费者
			select {
			case newListener := <-bs.addList:
				bs.addListener(newListener)
			case removeTarget := <-bs.removeList:
				bs.removeListener(removeTarget)
			case v, ok := <-bs.in:
				// 如果广播通道关闭，则关闭掉所有的消费者通道
				if !ok {
					goto terminate
				}
				//转发给所有的消费者
				for _, listener := range bs.listeners {
					if listener == nil {
						continue
					}
					listener <- v
				}
			}
		}
	terminate:
		//关闭所有的消费通道
		for _, ch := range bs.listeners {
			if ch == nil {
				continue
			}
			close(ch)

		}
	}()
	return bs.in
}

// private
func (bs *BroadcastService[T]) addListener(ch chan T) {
	for i, v := range bs.listeners {
		if v == nil {
			bs.listeners[i] = ch
			return
		}
	}
	bs.listeners = append(bs.listeners, ch)
}

func (bs *BroadcastService[T]) removeListener(ch chan T) {
	for i, v := range bs.listeners {
		if v == ch {
			bs.listeners[i] = nil
			// 一定要关闭! 否则监听它的goroutine将会一直block
			close(ch)
			return
		}
	}
}
