package service

import (
	"time"

	"github.com/cloudandbuild/blog-article-resrouces/content/blog/article-5/interfaces"
)

type AService struct {
	ARepository interfaces.IRepository
}

func (r *AService) DoSomething() error {
	d := &interfaces.Data{
		ID:        1,
		UpdatedAt: time.Now(),
	}

	return r.ARepository.Update(d)
}
