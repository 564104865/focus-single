package controller

import (
	"context"

	"focus-single/api/v1"
	"focus-single/internal/consts"
	"focus-single/internal/model"
	"focus-single/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

// 用户管理
var User = cUser{}

type cUser struct{}

func (a *cUser) Index(ctx context.Context, req *v1.UserIndexReq) (res *v1.UserIndexRes, err error) {
	err = a.getContentList(ctx, req.UserId, req.ContentGetListCommonReq)
	return
}

func (a *cUser) Article(ctx context.Context, req *v1.UserArticleReq) (res *v1.UserArticleRes, err error) {
	req.Type = consts.ContentTypeArticle
	err = a.getContentList(ctx, req.UserId, req.ContentGetListCommonReq)
	return
}

func (a *cUser) Topic(ctx context.Context, req *v1.UserTopicReq) (res *v1.UserTopicRes, err error) {
	req.Type = consts.ContentTypeTopic
	err = a.getContentList(ctx, req.UserId, req.ContentGetListCommonReq)
	return
}

func (a *cUser) Ask(ctx context.Context, req *v1.UserAskReq) (res *v1.UserAskRes, err error) {
	req.Type = consts.ContentTypeAsk
	err = a.getContentList(ctx, req.UserId, req.ContentGetListCommonReq)
	return
}

func (a *cUser) getContentList(ctx context.Context, userId uint, req v1.ContentGetListCommonReq) error {
	if out, err := service.User().GetList(ctx, model.UserGetContentListInput{
		ContentGetListInput: model.ContentGetListInput{
			Type:       req.Type,
			CategoryId: req.CategoryId,
			Page:       req.Page,
			Size:       req.Size,
			Sort:       req.Sort,
			UserId:     userId,
		},
	}); err != nil {
		return err
	} else {
		title := service.View().GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: req.Type,
			CategoryId:  req.CategoryId,
		})
		service.View().Render(ctx, model.View{
			ContentType: req.Type,
			Data:        out,
			Title:       title,
		})
		return nil
	}
}

func (a *cUser) Logout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error) {
	if err = service.User().Logout(ctx); err != nil {
		return
	}
	g.RequestFromCtx(ctx).Response.RedirectTo(service.Middleware().LoginUrl)
	return
}
