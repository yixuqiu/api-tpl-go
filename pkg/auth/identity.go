package auth

import (
	"api/ent"
	"api/ent/user"
	"api/lib/db"
	"api/lib/log"

	"context"
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/shenghui0779/yiigo/crypto"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type CtxKeyAuth int

const AuthIdentityKey CtxKeyAuth = 0

// Identity 授权身份
type Identity interface {
	// ID 授权ID
	ID() int64
	// AuthToken 生成auth_token
	AuthToken() (string, error)
	// Check 校验
	Check(ctx context.Context) error
	// String 用于日志记录
	String() string
}

type identity struct {
	I int64  `json:"i,omitempty"`
	T string `json:"t,omitempty"`
}

func (i *identity) ID() int64 {
	return i.I
}

func (i *identity) AuthToken() (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", errors.Wrap(err, "marshal identity")
	}

	key := []byte(viper.GetString("app.secret"))

	ct, err := crypto.AESEncryptCBC(key, key[:aes.BlockSize], b)
	if err != nil {
		return "", errors.Wrap(err, "encrypt identity")
	}

	return ct.String(), nil
}

func (i *identity) Check(ctx context.Context) error {
	if i.I == 0 {
		return errors.New("未授权，请先登录")
	}

	record, err := db.Client().User.Query().Unique(false).Select(
		user.FieldID,
		user.FieldLoginToken,
	).Where(user.ID(i.I)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.New("无效的账号")
		}

		log.Error(ctx, "Error auth check", zap.Error(err))

		return errors.New("授权校验失败")
	}

	if len(record.LoginToken) == 0 || record.LoginToken != i.T {
		return errors.New("授权已失效")
	}

	return nil
}

func (i *identity) String() string {
	if i.I == 0 {
		return "<nil>"
	}

	return fmt.Sprintf("id:%d|token:%s", i.I, i.T)
}

// NewEmptyIdentity 空授权信息
func NewEmptyIdentity() Identity {
	return new(identity)
}

// NewIdentity 用户授权信息
func NewIdentity(id int64, token string) Identity {
	return &identity{
		I: id,
		T: token,
	}
}

// GetIdentity 获取授权信息
func GetIdentity(ctx context.Context) Identity {
	if ctx == nil {
		return NewEmptyIdentity()
	}

	identity, ok := ctx.Value(AuthIdentityKey).(Identity)
	if !ok {
		return NewEmptyIdentity()
	}

	return identity
}

// AuthTokenToIdentity 解析授权Token
func AuthTokenToIdentity(ctx context.Context, token string) Identity {
	cipherText, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		log.Error(ctx, "Error invalid auth_token", zap.Error(err))
		return NewEmptyIdentity()
	}

	key := []byte(viper.GetString("app.secret"))

	plainText, err := crypto.AESDecryptCBC(key, key[:aes.BlockSize], cipherText)
	if err != nil {
		log.Error(ctx, "Error invalid auth_token", zap.Error(err))
		return NewEmptyIdentity()
	}

	identity := NewEmptyIdentity()
	if err = json.Unmarshal(plainText, identity); err != nil {
		log.Error(ctx, "Error invalid auth_token", zap.Error(err))
		// 此处应返回空Identify，因为若仅部分字段解析失败，Identity可能依然有效
		return NewEmptyIdentity()
	}

	return identity
}
