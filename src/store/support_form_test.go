package store

import (
	"context"
	"testing"

	"github.com/sisu-network/pairswap-be/src/model"
	"github.com/stretchr/testify/require"
)

func TestSupportFormStore(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	t.Run("create_support_form", func(t *testing.T) {
		t.Parallel()

		db, err := NewDBStores(dbConfig)
		require.NoError(t, err)
		supportForm := &model.SupportForm{
			Name:    "David",
			Email:   "david@gmail.com",
			TxURL:   "https://etherscan.io/tx/0x3612544d76174af7bbdd4a222cba8d73e0e89fe741ad0aa77d8a0a774e00d48f",
			Comment: "Waiting too long",
		}

		require.NoError(t, db.SupportFormStore.CreateSupportForm(ctx, supportForm))

		insertedSupportForm, err := db.SupportFormStore.GetById(ctx, supportForm.Id)
		require.NoError(t, err)
		assertSupportForm(t, supportForm, insertedSupportForm)
	})
}

func assertSupportForm(t *testing.T, expected, actual *model.SupportForm) {
	require.Equal(t, expected.Name, actual.Name)
	require.Equal(t, expected.Email, actual.Email)
	require.Equal(t, expected.TxURL, actual.TxURL)
	require.Equal(t, expected.Comment, actual.Comment)
}
