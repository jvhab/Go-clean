package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PaginationQuery struct {
	Size    int    `json:"size,omitempty"`
	Page    int    `json:"page,omitempty"`
	OrderBy string `json:"orderBy,omitempty"`
}

func (p *PaginationQuery) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		p.Size = 10
		return nil
	}
	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}
	p.Size = n
	return nil
}

func (p *PaginationQuery) SetPage(pageQuery string) error {
	if pageQuery == "" {
		p.Page = 0
		return nil
	}
	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return err
	}
	p.Page = n
	return nil
}

func (p *PaginationQuery) SetOrderBy(orderByQuery string) {
	p.OrderBy = orderByQuery
}

func (p *PaginationQuery) GetOffset() int {
	if p.Page == 0 {
		return 0
	}
	return (p.Page - 1) * p.Size
}

func (p *PaginationQuery) GetSize() int {
	return p.Size
}

func (p *PaginationQuery) GetPage() int {
	return p.Page
}

func (p *PaginationQuery) GetOrderBy() string {
	return p.OrderBy
}

func (p *PaginationQuery) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v&orderBy=%s", p.GetPage(), p.GetSize(), p.GetOrderBy())
}

func GetPaginationFromCtx(c *gin.Context) (*PaginationQuery, error) {
	p := &PaginationQuery{}
	if err := p.SetPage(c.Params.ByName("page")); err != nil {
		return nil, err
	}
	if err := p.SetSize(c.Params.ByName("size")); err != nil {
		return nil, err
	}
	p.SetOrderBy(c.Params.ByName("orderBy"))
	return p, nil
}
