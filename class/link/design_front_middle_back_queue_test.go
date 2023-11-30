package link

import (
	"container/list"
)

// List的API了解一下还是特别有用的。
// 链表头部操作：
// 链表尾部操作：
// 插入到某个节点前面：
// 插入到某个节点后面：
// 获取头结点：
// 获取尾部节点：
// 移除某个节点：

type FrontMiddleBackQueue struct {
	list   *list.List
	middle *list.Element
}

// 从上面来看，只要你能想到的，就能在上面找到相应的处理方式；
// 对于上面的每个操作，我们只需要添加节点，删除节点，然后处理中间节点即可，这道题没啥意思；
