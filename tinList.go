package tinList

// TinList 双向链表 表头
type TinList struct {
	rootEl TinElement
	len    int
}

// TinElement 双向链表 每个元素
type TinElement struct {
	nextEl, froEl *TinElement
	list          *TinList
	Value         interface{}
}

// NewList 初始化list
func NewList() *TinList {
	return new(TinList).initList()
}

// initList 初始化操作
func (l *TinList) initList() *TinList {
	l.rootEl.nextEl = &l.rootEl
	l.rootEl.froEl = &l.rootEl
	l.len = 0

	return l
}

// PushGetFront 把数据放入表头 并返回元素指针
func (l *TinList) PushGetFront(val interface{}) *TinElement {
	insertElement := &TinElement{Value: val}
	l.Insert(insertElement, &l.rootEl)
	return insertElement
}

// PushFront 把数据放入表头 并返回list指针 用于链式调用
func (l *TinList) PushFront(val interface{}) *TinList {
	l.PushGetFront(val)
	return l
}

// PushGetBack 把数据放入表尾 并返回元素指针
func (l *TinList) PushGetBack(val interface{}) *TinElement {
	insertElement := &TinElement{Value: val}
	l.Insert(insertElement, l.rootEl.froEl)
	return insertElement
}

// PushBack 把数据放入表尾 并返回list指针 用于链式调用
func (l *TinList) PushBack(val interface{}) *TinList {
	l.PushGetBack(val)
	return l
}

// Len 返回list长度
func (l *TinList) Len() int {
	return l.len
}

// Insert 在after元素后面加入一个新元素e
func (l *TinList) Insert(e *TinElement, after *TinElement) {
	e.froEl = after
	e.nextEl = after.nextEl
	e.nextEl.froEl = e
	e.froEl.nextEl = e
	e.list = l
	l.len++
}

// GetFirst 获取第一个元素
func (l *TinList) GetFirst() *TinElement {
	return l.rootEl.nextEl
}

// GetNext 获取下一个元素 用于迭代
func (e *TinElement) GetNext() *TinElement {
	if ne := e.nextEl; ne != nil && ne != &e.list.rootEl {
		return ne
	}

	return nil
}

// GetFront 获取上一个元素 用于迭代
func (e *TinElement) GetFront() *TinElement {
	if ne := e.froEl; ne != nil && ne != &e.list.rootEl {
		return ne
	}

	return nil
}
