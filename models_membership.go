package memberships

import (
	"context"

	"github.com/graphql-services/memberships/database"
)

type Membership struct {
	ID             string  `json:"id" gorm:"primary_key"`
	EntityID       string  `json:"entityID" gorm:"unique_index:memberentity"`
	Entity         *string `json:"entity"`
	Role           *string `json:"role"`
	MemberEntity   Member  `json:"member"`
	MemberEntityID string  `gorm:"unique_index:memberentity"`
}

func (m *Membership) Member(ctx context.Context) (member Member) {
	db := ctx.Value(DBContextKey).(*database.DB)
	db.Query().Model(m).Related(&member, "MemberEntity")
	return
}

func (m *Membership) Is_Entity() {}
