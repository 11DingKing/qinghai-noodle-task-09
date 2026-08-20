package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask09(t *testing.T) {
	now := time.Now()
	s := NewService(NewRegistry(), func() time.Time { return now })
	chili := IngredientLot{ID: "c", OriginRegion: "循化", ProducedAt: now.Add(-time.Hour), ExpiresAt: now.Add(time.Hour)}
	oil := IngredientLot{ID: "o"}
	b := ChiliOilBatch{ChiliLotID: "c", OilLotID: "o", Cultivar: "循化线辣椒", ProducedAt: now}
	require.NoError(t, s.CheckChiliOil(context.Background(), b, chili, oil))
}
