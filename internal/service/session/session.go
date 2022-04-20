package session

import (
	"context"

	"focus-single/internal/model/entity"
	"focus-single/internal/service/bizctx"
)

// Notice 存放在Session中的提示信息，往往使用后则删除
type Notice struct {
	Type    string // 消息类型
	Content string // 消息内容
}

const (
	sessionKeyUser         = "SessionKeyUser"    // 用户信息存放在Session中的Key
	sessionKeyLoginReferer = "SessionKeyReferer" // Referer存储，当已存在该session时不会更新。用于用户未登录时引导用户登录，并在登录后跳转到登录前页面。
	sessionKeyNotice       = "SessionKeyNotice"  // 存放在Session中的提示信息，往往使用后则删除
)

// 设置用户Session.
func SetUser(ctx context.Context, user *entity.User) error {
	return bizctx.Get(ctx).Session.Set(sessionKeyUser, user)
}

// 获取当前登录的用户信息对象，如果用户未登录返回nil。
func GetUser(ctx context.Context) *entity.User {
	customCtx := bizctx.Get(ctx)
	if customCtx != nil {
		v, _ := customCtx.Session.Get(sessionKeyUser)
		if !v.IsNil() {
			var user *entity.User
			_ = v.Struct(&user)
			return user
		}
	}
	return &entity.User{}
}

// 删除用户Session。
func RemoveUser(ctx context.Context) error {
	customCtx := bizctx.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyUser)
	}
	return nil
}

// 设置LoginReferer.
func SetLoginReferer(ctx context.Context, referer string) error {
	if GetLoginReferer(ctx) == "" {
		customCtx := bizctx.Get(ctx)
		if customCtx != nil {
			return customCtx.Session.Set(sessionKeyLoginReferer, referer)
		}
	}
	return nil
}

// 获取LoginReferer.
func GetLoginReferer(ctx context.Context) string {
	customCtx := bizctx.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.MustGet(sessionKeyLoginReferer).String()
	}
	return ""
}

// 删除LoginReferer.
func RemoveLoginReferer(ctx context.Context) error {
	customCtx := bizctx.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyLoginReferer)
	}
	return nil
}

// 设置Notice
func SetNotice(ctx context.Context, message *Notice) error {
	customCtx := bizctx.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Set(sessionKeyNotice, message)
	}
	return nil
}

// 获取Notice
func GetNotice(ctx context.Context) (*Notice, error) {
	customCtx := bizctx.Get(ctx)
	if customCtx != nil {
		var message *Notice
		v, err := customCtx.Session.Get(sessionKeyNotice)
		if err != nil {
			return nil, err
		}
		if err = v.Scan(&message); err != nil {
			return nil, err
		}
		return message, nil
	}
	return nil, nil
}

// 删除Notice
func RemoveNotice(ctx context.Context) error {
	customCtx := bizctx.Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyNotice)
	}
	return nil
}