// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/lumalabs/luma-agents-cli/internal/mocktest"
)

func TestFilesCreate(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--auth-token", "string",
			"files", "create",
			"--mime-type", "x",
			"--size-bytes", "1",
			"--expires-at", "'2019-12-27T18:11:19.117Z'",
			"--filename", "filename",
			"--purpose", "input",
			"--user-id", "user_id",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"mime_type: x\n" +
			"size_bytes: 1\n" +
			"expires_at: '2019-12-27T18:11:19.117Z'\n" +
			"filename: filename\n" +
			"purpose: input\n" +
			"user_id: user_id\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--auth-token", "string",
			"files", "create",
		)
	})
}

func TestFilesList(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--auth-token", "string",
			"files", "list",
			"--cursor", "cursor",
			"--limit", "1",
			"--purpose", "input",
			"--state", "pending",
		)
	})
}

func TestFilesDelete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--auth-token", "string",
			"files", "delete",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestFilesComplete(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--auth-token", "string",
			"files", "complete",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestFilesGet(t *testing.T) {
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--auth-token", "string",
			"files", "get",
			"--file-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
