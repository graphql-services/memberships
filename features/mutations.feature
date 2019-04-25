Feature: Managing memberships

    Scenario: Add membership and check response
        When I send query:
            """
            mutation { createMembership(input:{memberID:"john.doe",entityID:"default",entity:"project",role:"admin"}) { entityID entity role } }
            """
        Then the response should be:
            """
            {
            "createMembership" : {"entityID":"default","entity":"project","role":"admin" }
            }
            """

    Scenario: Add membership and fetch by entityID
        Given I send query:
            """
            mutation { createMembership(input:{memberID:"john.doe",entityID:"blah123",entity:"foo"}) { entityID entity } }
            """
        When I send query:
            """
            query { memberships(entityID:"blah123") { entity entityID member { id } } }
            """
        Then the response should be:
            """
            {
            "memberships" : [{ "entity":"foo","entityID":"blah123","member":{"id":"john.doe"} }]
            }
            """
