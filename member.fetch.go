package memberships

func getMember(id string) (member Member, err error) {
	member = Member{
		ID:   id,
		Name: "john.doe",
	}
	return
}
