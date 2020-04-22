package viewertest

import (
	"net/http"

	"github.com/facebookincubator/symphony/graph/viewer"
)

const (
	TenantHeader  = viewer.TenantHeader
	DefaultTenant = "test"
	UserHeader    = viewer.UserHeader
	DefaultUser   = "tester@example.com"
	RoleHeader    = viewer.RoleHeader
	DefaultRole   = "superuser"
)

func SetDefaultViewerHeaders(req *http.Request) {
	req.Header.Set(TenantHeader, DefaultTenant)
	req.Header.Set(UserHeader, DefaultUser)
	req.Header.Set(RoleHeader, RoleHeader)
}
