Feature: Invitations

    Scenario: Invite new member and check response
        When I send query:
            """
            mutation { inviteMember(input:{email:"john.doe@example.com",entityID:"default",entity:"project",role:"admin"}) { entityID entity role member { email } } }
            """
        Then the response should be:
            """
            {
                "inviteMember": {
                    "entityID": "default",
                    "entity": "project",
                    "role": "admin",
                    "member": {
                        "email": "john.doe@example.com"
                    }
                }
            }
            """

    Scenario: Invite new member and fetch it in separate request
    Scenario: Invite already existing member