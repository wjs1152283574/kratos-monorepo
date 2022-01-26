package data

import (
	"casso/app/user/service/internal/biz"
	"casso/app/user/service/internal/pkg/utill/passmd5"
	"casso/pkg/errors/normal"
	"casso/pkg/util/pagination"
	"casso/pkg/util/token"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (r *UserRepo) Create(ctx context.Context, b *biz.User) (*biz.User, error) {
	user := &biz.User{Name: b.Name, Age: b.Age, Mobile: b.Mobile, Pass: passmd5.Base64Md5(b.Pass)}
	err := r.data.db.WithContext(ctx).Create(user).First(user).Error
	if err != nil {
		r.log.Errorf("[data.Create] err : %#v", err)
		return &biz.User{}, normal.UnknownError
	}
	return user, nil
}

func (r *UserRepo) Get(ctx context.Context, id int64) (*biz.User, error) {
	user := biz.User{}
	err := r.data.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return &biz.User{}, normal.RecordNotFound
	}
	return &user, nil
}

func (r *UserRepo) Update(ctx context.Context, b *biz.User) (*biz.User, error) {
	user := biz.User{}
	if err := r.data.db.Updates(b).First(&user).Error; err != nil {
		return &biz.User{}, normal.UnknownError
	}
	return &user, nil
}

func (r *UserRepo) Delete(ctx context.Context, id int64) (*biz.User, error) {
	user := biz.User{}
	user.ID = uint(id)
	result := r.data.db.WithContext(ctx).First(&user).Delete(&user, id).Error
	if result != nil {
		return &biz.User{}, normal.UnknownError
	}
	return &user, nil
}

func (r *UserRepo) ListUser(ctx context.Context, pageNum, pageSize int64) ([]*biz.User, error) {
	var userList []biz.User
	result := r.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Find(&userList)
	if result.Error != nil {
		return nil, result.Error
	}
	rv := make([]*biz.User, 0)
	for _, user := range userList {
		rv = append(rv, &biz.User{
			Name: user.Name,
		})
	}
	return rv, nil
}

// GetToken check user exit and return token
func (r *UserRepo) GetToken(ctx context.Context, u *biz.UserForToken) (string, error) {
	var user biz.User
	result := r.data.db.WithContext(ctx).Where("mobile = ?", u.Mobile).First(&user)
	if result.Error != nil {
		return "", normal.RecordNotFound
	}
	if user.Pass != passmd5.Base64Md5(u.Pass) {
		return "", normal.InvalidParams
	}
	t, err := token.NewJWT().CreateToken(token.CustomClaims{
		ID: int(user.ID),
	})
	if err != nil {
		return "", normal.MakeTokenFaild
	}
	return t, nil
}

func (r *UserRepo) GetUserByName(ctx context.Context, name string) (*biz.User, error) {
	return &biz.User{}, nil
}
