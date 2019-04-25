package memberships

import (
	"context"

	"github.com/graphql-services/memberships/database"
	uuid "github.com/satori/go.uuid"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	DB *database.DB
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) getOrCreateMember(ctx context.Context, id string) (member Member, err error) {
	res := r.DB.Query().First(&member, "id = ?", id)
	err = res.Error

	if err != nil && !res.RecordNotFound() {
		return
	}

	if res.RecordNotFound() {
		member, err = getMember(id)
		if err != nil {
			return
		}
		err = r.DB.Query().Save(member).Error
	}

	return
}

func (r *mutationResolver) InviteMember(ctx context.Context, input *MembershipInvitationInput) (membership *Membership, err error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateMembership(ctx context.Context, input MembershipInput) (membership *Membership, err error) {
	member, err := r.getOrCreateMember(ctx, input.MemberID)

	membership = &Membership{
		ID:           uuid.Must(uuid.NewV4()).String(),
		Entity:       input.Entity,
		EntityID:     input.EntityID,
		Role:         input.Role,
		MemberEntity: member,
	}

	err = r.DB.Query().Save(membership).Error

	return
}
func (r *mutationResolver) DeleteMembership(ctx context.Context, id string) (membership *Membership, err error) {
	err = r.DB.Query().Delete(&Membership{ID: id}).Error
	return
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Member(ctx context.Context, id string) (member *Member, err error) {
	member = &Member{}
	err = r.DB.Query().First(member, "id = ?", id).Error
	return
}
func (r *queryResolver) Members(ctx context.Context, q *string) (members []Member, err error) {
	query := r.DB.Query().Model(&Member{})

	if q != nil {
		query = query.Where("name LIKE ?", "%"+*q+"%")
	}

	err = query.Find(members).Error
	return
}
func (r *queryResolver) Membership(ctx context.Context, id string) (membership *Membership, err error) {
	err = r.DB.Query().Model(&Membership{}).First(membership, "id = ?", id).Error
	return
}
func (r *queryResolver) Memberships(ctx context.Context, memberID *string, entityID *string, entity *string, role *string) (memberships []Membership, err error) {
	memberships = []Membership{}
	query := r.DB.Query()

	// if q != nil {
	// 	query = query.Where("name LIKE ?", "%"+*q+"%")
	// }

	err = query.Find(&memberships).Error
	return
}
