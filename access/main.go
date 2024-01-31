package main

import (
    "context"
    "fmt"
    "os"

    "github.com/okta/okta-sdk-golang/v2/okta"
)

func main() {
    // Set up the Okta client
    orgUrl := os.Getenv("OKTA_ORG_URL")
    token := os.Getenv("OKTA_API_TOKEN")
    ctx := context.TODO()
    client, err := okta.NewClient(ctx, okta.WithOrgUrl(orgUrl), okta.WithToken(token))
    if err != nil {
        fmt.Printf("Error creating Okta client: %v\n", err)
        return
    }

    // Get the user ID based on the email
    userEmail := "user@example.com" // Replace with the actual user email
    user, resp, err := client.User.ListUsers(ctx, &okta.ListUsersOptions{Q: userEmail})
    if err != nil {
        fmt.Printf("Error listing users: %v\n", err)
        return
    }
    if len(user) == 0 {
        fmt.Printf("No user found with email: %s\n", userEmail)
        return
    }
    userID := user[0].Id
    fmt.Printf("User ID for email %s: %s\n", userEmail, userID)

    // Get the group ID based on the group name
    groupName := "exampleGroupName" // Replace with the actual group name
    groups, _, err := client.Group.ListGroups(ctx, &okta.ListGroupsOptions{Q: groupName})
    if err != nil {
        fmt.Printf("Error listing groups: %v\n", err)
        return
    }
    if len(groups) == 0 {
        fmt.Printf("No group found with name: %s\n", groupName)
        return
    }
    groupID := groups[0].Id
    fmt.Printf("Group ID for group name %s: %s\n", groupName, groupID)

    // Add the user to the group
    _, err = client.Group.AddUserToGroup(ctx, groupID, userID)
    if err != nil {
        fmt.Printf("Error adding user to group: %v\n", err)
        return
    }
    fmt.Printf("User %s added to group %s successfully.\n", userEmail, groupName)
}
