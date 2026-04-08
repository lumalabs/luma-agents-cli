// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/luma-agents-cli/internal/mocktest"
	"github.com/stainless-sdks/luma-agents-cli/internal/requestflag"
)

func TestUsersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users", "create",
			"--id", "10",
			"--email", "john@email.com",
			"--first-name", "John",
			"--last-name", "James",
			"--password", "12345",
			"--phone", "12345",
			"--username", "theUser",
			"--user-status", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: 10\n" +
			"email: john@email.com\n" +
			"firstName: John\n" +
			"lastName: James\n" +
			"password: '12345'\n" +
			"phone: '12345'\n" +
			"username: theUser\n" +
			"userStatus: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"users", "create",
		)
	})
}

func TestUsersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users", "retrieve",
			"--username", "username",
		)
	})
}

func TestUsersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users", "update",
			"--existing-username", "username",
			"--id", "10",
			"--email", "john@email.com",
			"--first-name", "John",
			"--last-name", "James",
			"--password", "12345",
			"--phone", "12345",
			"--username", "theUser",
			"--user-status", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"id: 10\n" +
			"email: john@email.com\n" +
			"firstName: John\n" +
			"lastName: James\n" +
			"password: '12345'\n" +
			"phone: '12345'\n" +
			"username: theUser\n" +
			"userStatus: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"users", "update",
			"--existing-username", "username",
		)
	})
}

func TestUsersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users", "delete",
			"--username", "username",
		)
	})
}

func TestUsersCreateWithList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users", "create-with-list",
			"--item", "{id: 10, email: john@email.com, firstName: John, lastName: James, password: '12345', phone: '12345', username: theUser, userStatus: 1}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(usersCreateWithList)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users", "create-with-list",
			"--item.id", "10",
			"--item.email", "john@email.com",
			"--item.first-name", "John",
			"--item.last-name", "James",
			"--item.password", "12345",
			"--item.phone", "12345",
			"--item.username", "theUser",
			"--item.user-status", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"- id: 10\n" +
			"  email: john@email.com\n" +
			"  firstName: John\n" +
			"  lastName: James\n" +
			"  password: '12345'\n" +
			"  phone: '12345'\n" +
			"  username: theUser\n" +
			"  userStatus: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"users", "create-with-list",
		)
	})
}

func TestUsersLogin(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users", "login",
			"--password", "password",
			"--username", "username",
		)
	})
}

func TestUsersLogout(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users", "logout",
		)
	})
}
