package rbac

import (
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var (
	Once        sync.Once
	NewEnforcer *casbin.Enforcer
)

type CasbinSubRule struct {
	Path   string `json:"path" form:"path"`
	Method string `json:"method" form:"method"`
}

type CasbinRules struct {
	Role        string          `json:"role" form:"role"`
	CasbinInfos []CasbinSubRule `json:"casbininfos" form:"casbininfos"`
}

type CasbinRole struct {
	UserID string `json:"user_id" form:"user_id"`
	Role   string `json:"role" form:"role"`
}

type casbinImpl struct {
	Adaptor *gorm.DB
}

func (c *casbinImpl) NewCasbin() *casbin.Enforcer {
	Once.Do(func() {
		modelText := `
			[request_definition]
			r = sub, obj, act

			[policy_definition]
			p = sub, obj, act

			[role_definition]
			g = _, _

			[policy_effect]
			e = some(where (p.eft == allow))

			[matchers]
			m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act || r.sub == "admin"
		`
		model, _ := model.NewModelFromString(modelText)
		adapter, _ := gormadapter.NewAdapterByDB(c.Adaptor)
		NewEnforcer, _ = casbin.NewEnforcer(model, adapter)
		NewEnforcer.LoadPolicy()
	})
	return NewEnforcer
}

func (c *casbinImpl) AddPolicy(role, path, method string) {
	e := c.NewCasbin()
	e.AddPolicy(role, path, method)
}

func (c *casbinImpl) AddPolicies(role string, locals []CasbinSubRule) {
	e := c.NewCasbin()
	rules := [][]string{}
	for _, v := range locals {
		rules = append(rules, []string{role, v.Path, v.Method})
	}
	e.AddPolicies(rules)
}

func (c *casbinImpl) AddRole(user string, roles []string) (bool, error) {
	e := c.NewCasbin()
	rules := [][]string{}
	for _, v := range roles {
		rules = append(rules, []string{user, v})
	}
	return e.AddPolicies(rules)
}

func CasbinImpl(db *gorm.DB) *casbinImpl {
	return &casbinImpl{Adaptor: db}
}
