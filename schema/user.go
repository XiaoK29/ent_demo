package schema

import (
	gen "ent_demo/gen/ent"
	"ent_demo/gen/ent/hook"
	"ent_demo/schema/mixin"

	"context"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"

	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"golang.org/x/crypto/bcrypt"
)

type TestJSON struct {
	TestField string `json:"test_field"`
}

type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.BaseMixin{},
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("phone").Unique().NotEmpty().Comment("手机号"),
		field.String("password").NotEmpty().Comment("密码"),
		field.String("nikename").Optional().Comment("昵称"),
		field.String("email").Optional().Comment("邮箱"),
		field.String("avatar").Optional().Comment("头像"),
		field.Enum("gender").Optional().Values("male", "female").Comment("性别"),
		field.JSON("test_json", &TestJSON{}).Optional().Comment("测试json"),
	}
}

func (u User) Hooks() []ent.Hook {
	return []ent.Hook{
		u.EncryptPasswordOnMutate(),
	}
}

// EncryptPasswordOnMutate 密码加密
func (u User) EncryptPasswordOnMutate() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
			if password, ok := m.Password(); ok {
				bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				if err != nil {
					return nil, err
				}

				m.SetPassword(string(bytes))
			}

			return next.Mutate(ctx, m)
		})
	}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}

func (User) Indexes() []ent.Index {
	return []ent.Index{}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}
