package store

import (
	"context"
	"testing"

	"github.com/sisu-network/pairswap-be/src/model"
	"github.com/stretchr/testify/require"
)

func TestTokenStore(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	dbStore, err := NewDBStores(dbConfig)
	require.NoError(t, err)

	t.Run("CreateAndGetToken", func(t *testing.T) {
		t.Parallel()

		token := &model.Token{
			Name:     "Sisu network",
			Address:  "0x1234",
			Symbol:   "SISU",
			LogoURL:  "http://google.com",
			Decimals: 18,
		}
		require.NoError(t, dbStore.TokenStore.CreateToken(ctx, token))
		insertedToken, err := dbStore.TokenStore.GetById(ctx, token.Id)
		require.NoError(t, err)
		assertToken(t, token, insertedToken)
	})

	t.Run("GetAll", func(t *testing.T) {
		t.Parallel()

		nTokens := 10
		for i := 0; i < nTokens; i++ {
			token := &model.Token{
				Name:     "Sisu network",
				Address:  "0x1234",
				Symbol:   "SISU",
				LogoURL:  "http://google.com",
				Decimals: 18,
			}
			require.NoError(t, dbStore.TokenStore.CreateToken(ctx, token))
			insertedToken, err := dbStore.TokenStore.GetById(ctx, token.Id)
			require.NoError(t, err)
			assertToken(t, token, insertedToken)
		}

		allTokens, err := dbStore.TokenStore.GetAll(ctx)
		require.NoError(t, err)
		require.True(t, len(allTokens) >= nTokens)
	})
}

func assertToken(t *testing.T, expectedToken, actualToken *model.Token) {
	require.Equal(t, expectedToken.Name, actualToken.Name)
	require.Equal(t, expectedToken.Address, actualToken.Address)
	require.Equal(t, expectedToken.Symbol, actualToken.Symbol)
	require.Equal(t, expectedToken.LogoURL, actualToken.LogoURL)
	require.Equal(t, expectedToken.Decimals, actualToken.Decimals)
}
