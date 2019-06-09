package gormsql_test

import (
	"github.com/atymkiv/echo_frame_learning/blog/cmd/api/auth/platform/gormsql"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/mock/mockdb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindByEmail(t *testing.T) {
	db := mockdb.NewFakeDbOrFatal()

	cases := []struct {
		name    string
		email   string
		wantErr bool
		pgdb    *mockdb.DB
	}{
		{
			name:    "Fail on unknown email",
			email:   "test@l.com",
			wantErr: true,
			pgdb:    db,
		},
		{
			name:    "success",
			email:   "",
			wantErr: true,
			pgdb:    db,
		},
	}
	authDB := gormsql.NewUser(db)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := authDB.FindByEmail(tt.email)
			assert.Equal(t, tt.wantErr, err != nil)

		})
	}
}
