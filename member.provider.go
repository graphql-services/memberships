package memberships

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/machinebox/graphql"
)

const (
	fetchMemberQuery = `
query($id: ID!) {
	result: user(id: $id) {
		id
		email
		given_name
		family_name
		middle_name
	}
}  
`
	inviteMemberQuery = `
mutation($email: String!) {
	result: inviteUser(email: $email) {
		id
		email
		given_name
		family_name
		middle_name
	}
}  
`
)

type MemberProviderFetchResponse struct {
	Result *Member
}
type MemberProviderInviteResponse struct {
	Result Member
}

func fetchMember(ctx context.Context, id string) (member *Member, err error) {
	var res MemberProviderFetchResponse

	req := graphql.NewRequest(fetchMemberQuery)
	req.Var("id", id)
	err = sendRequest(ctx, req, &res)

	member = res.Result

	return
}

func inviteMember(ctx context.Context, email string) (member Member, err error) {
	var res MemberProviderInviteResponse

	req := graphql.NewRequest(inviteMemberQuery)
	req.Var("email", email)
	err = sendRequest(ctx, req, &res)

	member = res.Result

	return
}

func sendRequest(ctx context.Context, req *graphql.Request, data interface{}) error {
	URL := os.Getenv("MEMBER_PROVIDER_URL")

	if URL == "" {
		return fmt.Errorf("Missing required environment variable MEMBER_PROVIDER_URL")
	}

	client := graphql.NewClient(URL)
	client.Log = func(s string) {
		log.Println(s)
	}

	return client.Run(ctx, req, data)
}
