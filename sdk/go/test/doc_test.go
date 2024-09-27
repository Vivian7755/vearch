package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	client "github.com/vearch/vearch/v3/sdk/go"
	"github.com/vearch/vearch/v3/sdk/go/data"
	"github.com/vearch/vearch/v3/sdk/go/entities/models"
)

func upsertDocs(c *client.Client, dbName, spaceName string) (result *data.DocWrapper, err error) {
	ctx := context.Background()

	documents := []interface{}{
		map[string]interface{}{
			"_id":          "1",
			"field_int":    777,
			"field_float":  123.4,
			"field_vector": []float32{0.019698096, 0.041366003, 0.037426382, 0.088641435, 0.078792386, 0.07091314, 0.06303391, 0.011818858, 0.2028904, 0.11227915, 0.049245242, 0.03545657, 0.023637716, 0.04530562, 0.13000743, 0.10439991, 0.098490484, 0.04333581, 0.023637716, 0.037426382, 0.074852765, 0.059094287, 0.078792386, 0.033486765, 0.031516954, 0.07288296, 0.10833953, 0.13000743, 0.09455086, 0.039396193, 0.04727543, 0.03545657, 0.12606782, 0.061064098, 0.029547144, 0.011818858, 0.0137886675, 0.021667905, 0.065003715, 0.12606782, 0.03545657, 0.049245242, 0.09455086, 0.19698097, 0.12015839, 0.06303391, 0.11030934, 0.05712448, 0.19501115, 0.2028904, 0.1437961, 0.05318486, 0.06894334, 0.04530562, 0.027577335, 0.049245242, 0.05515467, 0.023637716, 0.0137886675, 0.03545657, 0.098490484, 0.13788667, 0.2028904, 0.04727543, 0.06697353, 0.15561496, 0.17728287, 0.08076219, 0.06697353, 0.037426382, 0.033486765, 0.04727543, 0.11621877, 0.10833953, 0.033486765, 0.041366003, 0.09061124, 0.16940363, 0.15167534, 0.06303391, 0.16940363, 0.04727543, 0.061064098, 0.06303391, 0.023637716, 0.05515467, 0.16152439, 0.17728287, 0.05712448, 0.08667162, 0.07288296, 0.049245242, 0.10636972, 0.039396193, 0.033486765, 0.0137886675, 0.06303391, 0.027577335, 0.019698096, 0.031516954, 0.04727543, 0.10636972, 0.15561496, 0.07682258, 0.13000743, 0.09061124, 0.09455086, 0.041366003, 0.025607524, 0.007879239, 0.039396193, 0.088641435, 0.023637716, 0.04530562, 0.074852765, 0.12803763, 0.13788667, 0.088641435, 0.05712448, 0.023637716, 0.13985649, 0.084701814, 0.059094287, 0.015758477, 0.0137886675, 0.027577335, 0.037426382, 0.07682258},
		},
		map[string]interface{}{
			"_id":          "2",
			"field_int":    888,
			"field_float":  345.6,
			"field_vector": []float32{28.0, 14.0, 29.0, 55.0, 3.0, 27.0, 2.0, 10.0, 68.0, 7.0, 3.0, 19.0, 1.0, 0.0, 1.0, 126.0, 23.0, 2.0, 7.0, 15.0, 5.0, 0.0, 1.0, 74.0, 5.0, 4.0, 9.0, 33.0, 49.0, 9.0, 4.0, 2.0, 27.0, 13.0, 1.0, 5.0, 56.0, 101.0, 12.0, 10.0, 140.0, 49.0, 4.0, 7.0, 8.0, 9.0, 7.0, 46.0, 34.0, 8.0, 12.0, 68.0, 34.0, 0.0, 1.0, 10.0, 34.0, 14.0, 8.0, 82.0, 115.0, 0.0, 0.0, 1.0, 18.0, 1.0, 2.0, 25.0, 51.0, 17.0, 16.0, 39.0, 57.0, 6.0, 23.0, 120.0, 5.0, 4.0, 19.0, 83.0, 66.0, 18.0, 43.0, 140.0, 13.0, 0.0, 2.0, 18.0, 13.0, 4.0, 8.0, 140.0, 56.0, 0.0, 2.0, 5.0, 14.0, 8.0, 13.0, 72.0, 7.0, 10.0, 2.0, 7.0, 35.0, 4.0, 25.0, 140.0, 2.0, 0.0, 0.0, 43.0, 50.0, 4.0, 9.0, 140.0, 28.0, 0.0, 0.0, 36.0, 9.0, 0.0, 2.0, 140.0, 19.0, 0.0, 9.0, 19.0},
		},
	}

	result, err = c.Data().Creator().WithDBName(dbName).WithSpaceName(spaceName).WithDocs(documents).Do(ctx)
	return result, err
}

func TestUpsertDoc(t *testing.T) {
	ctx := context.Background()
	dbName := "ts_db"
	spaceName := "ts_space"

	c := setupClient(t)

	err := createDB(c, dbName)
	require.Nil(t, err)

	err = createSpace(c, dbName, spaceName)
	require.Nil(t, err)

	result, err := upsertDocs(c, dbName, spaceName)
	require.Nil(t, err)
	fmt.Printf("result %v\n", result.Docs.Data.Total)

	err = c.Schema().SpaceDeleter().WithDBName(dbName).WithSpaceName(spaceName).Do(ctx)
	require.Nil(t, err)

	err = c.Schema().DBDeleter().WithDBName(dbName).Do(ctx)
	require.Nil(t, err)
}

func TestSearchtDoc(t *testing.T) {
	ctx := context.Background()
	dbName := "ts_db"
	spaceName := "ts_space"

	c := setupClient(t)
	err := createDB(c, dbName)
	require.Nil(t, err)

	err = createSpace(c, dbName, spaceName)
	require.Nil(t, err)

	_, err = upsertDocs(c, dbName, spaceName)
	require.Nil(t, err)

	vector := []models.Vector{
		{
			Field:   "field_vector",
			Feature: []float32{0.019698096, 0.041366003, 0.037426382, 0.088641435, 0.078792386, 0.07091314, 0.06303391, 0.011818858, 0.2028904, 0.11227915, 0.049245242, 0.03545657, 0.023637716, 0.04530562, 0.13000743, 0.10439991, 0.098490484, 0.04333581, 0.023637716, 0.037426382, 0.074852765, 0.059094287, 0.078792386, 0.033486765, 0.031516954, 0.07288296, 0.10833953, 0.13000743, 0.09455086, 0.039396193, 0.04727543, 0.03545657, 0.12606782, 0.061064098, 0.029547144, 0.011818858, 0.0137886675, 0.021667905, 0.065003715, 0.12606782, 0.03545657, 0.049245242, 0.09455086, 0.19698097, 0.12015839, 0.06303391, 0.11030934, 0.05712448, 0.19501115, 0.2028904, 0.1437961, 0.05318486, 0.06894334, 0.04530562, 0.027577335, 0.049245242, 0.05515467, 0.023637716, 0.0137886675, 0.03545657, 0.098490484, 0.13788667, 0.2028904, 0.04727543, 0.06697353, 0.15561496, 0.17728287, 0.08076219, 0.06697353, 0.037426382, 0.033486765, 0.04727543, 0.11621877, 0.10833953, 0.033486765, 0.041366003, 0.09061124, 0.16940363, 0.15167534, 0.06303391, 0.16940363, 0.04727543, 0.061064098, 0.06303391, 0.023637716, 0.05515467, 0.16152439, 0.17728287, 0.05712448, 0.08667162, 0.07288296, 0.049245242, 0.10636972, 0.039396193, 0.033486765, 0.0137886675, 0.06303391, 0.027577335, 0.019698096, 0.031516954, 0.04727543, 0.10636972, 0.15561496, 0.07682258, 0.13000743, 0.09061124, 0.09455086, 0.041366003, 0.025607524, 0.007879239, 0.039396193, 0.088641435, 0.023637716, 0.04530562, 0.074852765, 0.12803763, 0.13788667, 0.088641435, 0.05712448, 0.023637716, 0.13985649, 0.084701814, 0.059094287, 0.015758477, 0.0137886675, 0.027577335, 0.037426382, 0.07682258},
		},
	}

	result, err := c.Data().Searcher().WithDBName(dbName).WithSpaceName(spaceName).WithLimit(2).WithVectors(vector).Do(ctx)
	require.Nil(t, err)
	fmt.Printf("result %v\n", result.Docs.Data.Documents...)

	err = c.Schema().SpaceDeleter().WithDBName(dbName).WithSpaceName(spaceName).Do(ctx)
	require.Nil(t, err)

	err = c.Schema().DBDeleter().WithDBName(dbName).Do(ctx)
	require.Nil(t, err)
}

func TestSearchtDocWithFilter(t *testing.T) {
	ctx := context.Background()
	dbName := "ts_db"
	spaceName := "ts_space"

	c := setupClient(t)
	err := createDB(c, dbName)
	require.Nil(t, err)

	err = createSpace(c, dbName, spaceName)
	require.Nil(t, err)

	_, err = upsertDocs(c, dbName, spaceName)
	require.Nil(t, err)

	vector := []models.Vector{
		{
			Field:   "field_vector",
			Feature: []float32{0.019698096, 0.041366003, 0.037426382, 0.088641435, 0.078792386, 0.07091314, 0.06303391, 0.011818858, 0.2028904, 0.11227915, 0.049245242, 0.03545657, 0.023637716, 0.04530562, 0.13000743, 0.10439991, 0.098490484, 0.04333581, 0.023637716, 0.037426382, 0.074852765, 0.059094287, 0.078792386, 0.033486765, 0.031516954, 0.07288296, 0.10833953, 0.13000743, 0.09455086, 0.039396193, 0.04727543, 0.03545657, 0.12606782, 0.061064098, 0.029547144, 0.011818858, 0.0137886675, 0.021667905, 0.065003715, 0.12606782, 0.03545657, 0.049245242, 0.09455086, 0.19698097, 0.12015839, 0.06303391, 0.11030934, 0.05712448, 0.19501115, 0.2028904, 0.1437961, 0.05318486, 0.06894334, 0.04530562, 0.027577335, 0.049245242, 0.05515467, 0.023637716, 0.0137886675, 0.03545657, 0.098490484, 0.13788667, 0.2028904, 0.04727543, 0.06697353, 0.15561496, 0.17728287, 0.08076219, 0.06697353, 0.037426382, 0.033486765, 0.04727543, 0.11621877, 0.10833953, 0.033486765, 0.041366003, 0.09061124, 0.16940363, 0.15167534, 0.06303391, 0.16940363, 0.04727543, 0.061064098, 0.06303391, 0.023637716, 0.05515467, 0.16152439, 0.17728287, 0.05712448, 0.08667162, 0.07288296, 0.049245242, 0.10636972, 0.039396193, 0.033486765, 0.0137886675, 0.06303391, 0.027577335, 0.019698096, 0.031516954, 0.04727543, 0.10636972, 0.15561496, 0.07682258, 0.13000743, 0.09061124, 0.09455086, 0.041366003, 0.025607524, 0.007879239, 0.039396193, 0.088641435, 0.023637716, 0.04530562, 0.074852765, 0.12803763, 0.13788667, 0.088641435, 0.05712448, 0.023637716, 0.13985649, 0.084701814, 0.059094287, 0.015758477, 0.0137886675, 0.027577335, 0.037426382, 0.07682258},
		},
	}

	filter := &models.Filters{
		Operator: "AND",
		Conditions: []models.Condition{
			{
				Operator: "<",
				Field:    "field_float",
				Value:    200,
			},
		},
	}

	result, err := c.Data().Searcher().WithDBName(dbName).WithSpaceName(spaceName).WithLimit(2).WithVectors(vector).WithFilters(filter).Do(ctx)
	require.Nil(t, err)
	for i, doc := range result.Docs.Data.Documents {
		fmt.Printf("result[%d] %v\n", i, doc)
	}

	err = c.Schema().SpaceDeleter().WithDBName(dbName).WithSpaceName(spaceName).Do(ctx)
	require.Nil(t, err)

	err = c.Schema().DBDeleter().WithDBName(dbName).Do(ctx)
	require.Nil(t, err)
}

func TestQuerytDoc(t *testing.T) {
	ctx := context.Background()
	dbName := "ts_db"
	spaceName := "ts_space"

	c := setupClient(t)
	err := createDB(c, dbName)
	require.Nil(t, err)

	err = createSpace(c, dbName, spaceName)
	require.Nil(t, err)

	_, err = upsertDocs(c, dbName, spaceName)
	require.Nil(t, err)

	ids := []string{
		"1",
		"2",
	}

	result, err := c.Data().Query().WithDBName(dbName).WithSpaceName(spaceName).WithIDs(ids).Do(ctx)
	require.Nil(t, err)
	for i, doc := range result.Docs.Data.Documents {
		fmt.Printf("result[%d] %v\n", i, doc)
	}

	err = c.Schema().SpaceDeleter().WithDBName(dbName).WithSpaceName(spaceName).Do(ctx)
	require.Nil(t, err)

	err = c.Schema().DBDeleter().WithDBName(dbName).Do(ctx)
	require.Nil(t, err)
}

func TestDeletetDoc(t *testing.T) {
	ctx := context.Background()
	dbName := "ts_db"
	spaceName := "ts_space"

	c := setupClient(t)
	err := createDB(c, dbName)
	require.Nil(t, err)

	err = createSpace(c, dbName, spaceName)
	require.Nil(t, err)

	_, err = upsertDocs(c, dbName, spaceName)
	require.Nil(t, err)

	ids := []string{
		"1",
		"2",
	}

	result, err := c.Data().Deleter().WithDBName(dbName).WithSpaceName(spaceName).WithIDs(ids).Do(ctx)
	require.Nil(t, err)
	fmt.Printf("delete result %v\n", result.Docs.Data.DocumentsIDs)
	for i, doc := range result.Docs.Data.DocumentsIDs {
		fmt.Printf("result[%d] %v\n", i, doc)
	}

	err = c.Schema().SpaceDeleter().WithDBName(dbName).WithSpaceName(spaceName).Do(ctx)
	require.Nil(t, err)

	err = c.Schema().DBDeleter().WithDBName(dbName).Do(ctx)
	require.Nil(t, err)
}
