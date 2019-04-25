package memberships

import (
	"context"

	"github.com/graphql-services/memberships/database"
)

type Membership struct {
	ID             string  `json:"id" gorm:"primary_key"`
	EntityID       string  `json:"entityID"`
	Entity         *string `json:"entity"`
	Role           *string `json:"role"`
	Accepted       bool    `json:"accepted"`
	MemberEntity   Member  `json:"member"`
	MemberEntityID string
}

func (m *Membership) Member(ctx context.Context) (member Member) {
	db := ctx.Value(DBContextKey).(*database.DB)
	db.Query().Model(m).Related(&member, "MemberEntity")
	return
}
