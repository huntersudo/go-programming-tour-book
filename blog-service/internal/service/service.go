package service

import (
	"context"

	otgorm "github.com/eddycjy/opentracing-gorm"

	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

// New
// 这玩意的作用，可以理解为就是一个类名,承载各个方法，
// NEW是用必须持有的一些资源来初始化,用于传递ctx等
func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	// global.DBEngine init at begin
	// todo tracing
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	// svc.dao = dao.New(global.DBEngine)
	return svc
}
