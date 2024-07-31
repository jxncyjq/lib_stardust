package http

import (
	"fmt"
	"strings"
)

// http客户端请求分页结构
type Pageable struct {
	page  int
	size  int
	index int
	sort  []Sort
}

// http客户端请求排序结构
type Sort struct {
	Key       string
	Direction string
}

func NewPageable(page, size int) *Pageable {
	if size == 0 {
		size = 30
	}
	return &Pageable{page: page, size: size}
}

func (p *Pageable) Skip() int {
	s := p.size * p.page
	if s < 0 {
		s = 0
	}
	return s
}

func (p *Pageable) Page() int {
	return p.page
}
func (p *Pageable) Size() int {
	return p.size
}
func (p *Pageable) Sort() string {
	if nil == p.sort || 0 == len(p.sort) {
		return ""
	}

	sorts := make([]string, len(p.sort))
	for idx := range p.sort {
		sorts[idx] = fmt.Sprintf("%s %s", p.sort[idx].Key, p.sort[idx].Direction)
	}
	return strings.Join(sorts, ",")
}
