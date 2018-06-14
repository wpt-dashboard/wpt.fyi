package shared

import (
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

// LoadTestRuns loads the TestRun entities for the given parameters.
// It is encapsulated because we cannot run single queries with multiple inequality
// filters, so must load the keys and merge the results.
func LoadTestRuns(
	ctx context.Context,
	products []Product,
	labels mapset.Set,
	shas []string,
	from *time.Time,
	limit *int) (result []TestRun, err error) {
	var testRuns []TestRun
	baseQuery := datastore.NewQuery("TestRun")
	// NOTE(lukebjerring): While we can't filter on multiple SHAs, it's still much more efficient
	// to (pre-)filter for a single SHA during the query.
	if len(shas) == 1 && !IsLatest(shas[0]) {
		baseQuery = baseQuery.Filter("Revision =", shas[0])
	}
	experimentalOnly := false
	if labels != nil {
		for i := range labels.Iter() {
			label := i.(string)
			if IsStableBrowserName(label) {
				// Browser name labels are already handled in GetProductsForRequest (which produces `products`).
				continue
			}
			if label == ExperimentalLabel {
				// The "experimental" label is handled specially at the end of the function.
				// TODO(Hexcles): Remove this once we convert all history runs.
				experimentalOnly = true
				continue
			}
			baseQuery = baseQuery.Filter("Labels =", label)
		}
	}
	for _, product := range products {
		var prefiltered *mapset.Set
		query := baseQuery.Filter("BrowserName =", product.BrowserName)
		if product.BrowserVersion != "" {
			if prefiltered, err = loadKeysForBrowserVersion(ctx, query, product.BrowserVersion); err != nil {
				return nil, err
			}
		}
		// TODO(lukebjerring): Indexes + filtering for OS + version.
		query = query.Order("-CreatedAt")

		if from != nil {
			query = query.Filter("CreatedAt >", *from)
		}

		fetched, err := query.KeysOnly().GetAll(ctx, nil)
		if err != nil {
			return nil, err
		}
		var keys []*datastore.Key
		for _, key := range fetched {
			if len(shas) > 1 || limit == nil || *limit > len(keys) {
				if prefiltered == nil || (*prefiltered).Contains(key.String()) {
					keys = append(keys, key)
				}
			}
		}
		testRunResults := make(TestRuns, len(keys))
		if err = datastore.GetMulti(ctx, keys, testRunResults); err != nil {
			return nil, err
		}
		// Append the keys as ID
		for i, key := range keys {
			testRunResults[i].ID = key.IntID()
		}
		appended := 0
		for _, testRun := range testRunResults {
			// Handle the "experimental" label specially.
			// Some history experimental runs don't have the experimental
			// label; instead, their browser names have the suffix. We'd
			// like to support both the suffix and the label.
			// TODO(Hexcles): Remove this once we convert history runs.
			if experimentalOnly && !testRun.IsExperimental() {
				continue
			}
			if len(shas) > 1 && !contains(shas, testRun.Revision) {
				continue
			}
			testRuns = append(testRuns, testRun)
			appended++
			if limit != nil && appended >= *limit {
				break
			}
		}
	}
	return testRuns, nil
}

func contains(s []string, x string) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

// Loads any keys for a full string match or a version prefix (Between [version].* and [version].9*)
func loadKeysForBrowserVersion(ctx context.Context, query *datastore.Query, version string) (result *mapset.Set, err error) {
	versionQuery := VersionPrefix(query, "BrowserVersion", version, true)
	var keys []*datastore.Key
	keyset := mapset.NewSet()
	if keys, err = versionQuery.KeysOnly().GetAll(ctx, nil); err != nil {
		return nil, err
	}
	for _, key := range keys {
		keyset.Add(key.String())
	}
	if keys, err = query.Filter("BrowserVersion =", version).KeysOnly().GetAll(ctx, nil); err != nil {
		return nil, err
	}
	for _, key := range keys {
		keyset.Add(key.String())
	}
	return &keyset, nil
}

// VersionPrefix returns the given query with a prefix filter on the given
// field name, using the >= and < filters.
func VersionPrefix(query *datastore.Query, fieldName, versionPrefix string, desc bool) *datastore.Query {
	order := fieldName
	if desc {
		order = "-" + order
	}
	return query.
		Order(order).
		Filter(fieldName+" >=", fmt.Sprintf("%s.", versionPrefix)).
		Filter(fieldName+" <=", fmt.Sprintf("%s.%c", versionPrefix, '9'+1))
}
