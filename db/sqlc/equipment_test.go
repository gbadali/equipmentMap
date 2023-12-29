package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateEquipment(t *testing.T) {
	arg := CreateEquipmentParams{
		ModelNumber:    "test",
		Description:    "test",
		Location:       1,
		Parent:         sql.NullInt64{1, true},
		PoweredFrom:    sql.NullInt64{0, false},
		ControlledFrom: sql.NullInt64{0, false},
	}

	equipment, err := testQueries.CreateEquipment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, equipment)

	require.Equal(t, arg.ModelNumber, equipment.ModelNumber)
	require.Equal(t, arg.Description, equipment.Description)
	require.Equal(t, arg.Location, equipment.Location)
	require.Equal(t, arg.Parent, equipment.Parent)
	require.Equal(t, arg.PoweredFrom, equipment.PoweredFrom)
	require.Equal(t, arg.ControlledFrom, equipment.ControlledFrom)

}
