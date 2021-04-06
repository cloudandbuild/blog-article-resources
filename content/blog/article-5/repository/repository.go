package repository

import "github.com/cloudandbuild/blog-article-resrouces/content/blog/article-5/interfaces"

type ARespository struct {
	interfaces.IRepository
}

func NewRepository() *ARespository {
	return &ARespository{}
}
func (r *ARespository) Update(data *interfaces.Data) error {
	// not implemented yet
	return nil
}
