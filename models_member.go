package memberships

import (
	"context"

	"github.com/graphql-services/memberships/database"
)

type Member struct {
	ID                  string       `json:"id" gorm:"primary_key"`
	Name                string       `json:"name"`
	Email               string       `json:"email"`
	GivenName           string       `json:"given_name"`
	FamilyName          string       `json:"family_name"`
	MiddleName          string       `json:"middle_name"`
	MembershipsEntities []Membership `json:"memberships" gorm:"foreignkey:MemberEntityID"`
}

func (m *Member) Memberships(ctx context.Context) (memberships []Membership) {
	db := ctx.Value(DBContextKey).(*database.DB)
	db.Query().Model(m).Association("MembershipsEntities").Find(&memberships)
	return
}

func (m *Member) Is_Entity() {}
