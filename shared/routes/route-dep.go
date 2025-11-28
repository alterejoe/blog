package routes

import (
	"zotes/shared/interfaces"

	"github.com/casbin/casbin/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDependencies[T any](
	db *pgxpool.Pool,
	queries func(pgx.Tx) T,
	logger interfaces.Logger,
	session interfaces.SessionManager,
	sanitizer interfaces.Sanitizer,

	oidc interfaces.OIDCAuthenticator,
	jwtissuer interfaces.JWTIssuer,
	jwtvalidator interfaces.JWTValidator,
	// authManage interfaces.AuthAccess,
	// JWT interfaces.Auth0Target,
	rbac *casbin.Enforcer,
	s3manager interfaces.S3,
) *Dependencies[T] {
	return &Dependencies[T]{
		db:                db,
		queries:           queries,
		logger:            logger,
		session:           session,
		sanitizer:         sanitizer,
		oidcauthenticator: oidc,
		jwtissuer:         jwtissuer,
		jwtvalidator:      jwtvalidator,
		// authAccess: authManage,
		rbac: rbac,
		// JWT:        JWT,
		s3manager: s3manager,
	}
}

type Dependencies[T any] struct {
	db                *pgxpool.Pool
	queries           func(pgx.Tx) T
	logger            interfaces.Logger
	session           interfaces.SessionManager
	sanitizer         interfaces.Sanitizer
	oidcauthenticator interfaces.OIDCAuthenticator
	jwtissuer         interfaces.JWTIssuer
	jwtvalidator      interfaces.JWTValidator
	// authAccess interfaces.AuthAccess
	rbac *casbin.Enforcer
	// JWT        interfaces.Auth0Target
	s3manager interfaces.S3
}

func (d *Dependencies[T]) Logger() interfaces.Logger                 { return d.logger }
func (d *Dependencies[T]) SessionManager() interfaces.SessionManager { return d.session }
func (d *Dependencies[T]) Sanitizer() interfaces.Sanitizer           { return d.sanitizer }

// func (d *Dependencies[T]) DB() interfaces.DBQueryContext[T]          { return d.db }
func (d *Dependencies[T]) Session() interfaces.SessionManager { return d.session }
func (d *Dependencies[T]) OidcAuthenticator() interfaces.OIDCAuthenticator {
	return d.oidcauthenticator
}
func (d *Dependencies[T]) JwtIssuer() interfaces.JWTIssuer {
	return d.jwtissuer
}

func (d *Dependencies[T]) JwtValidator() interfaces.JWTValidator {
	return d.jwtvalidator
}

// func (d *Dependencies[T]) AuthAccess() interfaces.AuthAccess       { return d.authAccess }
func (d *Dependencies[T]) NewQueries(tx pgx.Tx) T {
	return d.queries(tx)
}
func (d *Dependencies[T]) S3() interfaces.S3 {
	return d.s3manager
}
func (d *Dependencies[T]) RBAC() *casbin.Enforcer {
	return d.rbac
}
