// +build medium

package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/web-platform-tests/wpt.fyi/shared"
	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

func TestCompleteSHAs(t *testing.T) {
	i, err := aetest.NewInstance(&aetest.Options{StronglyConsistentDatastore: true})
	assert.Nil(t, err)
	defer i.Close()
	r, err := i.NewRequest("GET", "/api/shas?complete", nil)
	assert.Nil(t, err)

	ctx := appengine.NewContext(r)
	browserNames, _ := shared.GetBrowserNames()

	// Nothing in datastore.
	shas, _ := getCompleteRunSHAs(ctx, nil, nil)
	assert.Equal(t, 0, len(shas))

	// Only 3 browsers.
	run := shared.TestRun{
		ProductAtRevision: shared.ProductAtRevision{
			Revision: "abcdef0000",
		},
		CreatedAt: time.Now().AddDate(0, 0, -1),
	}
	for _, browser := range browserNames[:len(browserNames)-1] {
		run.BrowserName = browser
		datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "TestRun", nil), &run)
	}
	shas, _ = getCompleteRunSHAs(ctx, nil, nil)
	assert.Equal(t, 0, len(shas))

	// All 4 browsers, but experimental.
	run.Revision = "abcdef0111"
	run.CreatedAt = time.Now().AddDate(0, 0, -2)
	for _, browser := range browserNames {
		run.BrowserName = browser + "-" + shared.ExperimentalLabel
		datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "TestRun", nil), &run)
	}
	shas, _ = getCompleteRunSHAs(ctx, nil, nil)
	assert.Equal(t, 0, len(shas))

	// 2 browsers, and other 2, but experimental.
	run.Revision = "abcdef0222"
	run.CreatedAt = time.Now().AddDate(0, 0, -2)
	for i, browser := range browserNames {
		run.BrowserName = browser
		if i > 1 {
			run.BrowserName += "-" + shared.ExperimentalLabel
		}
		datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "TestRun", nil), &run)
	}
	shas, _ = getCompleteRunSHAs(ctx, nil, nil)
	assert.Equal(t, 0, len(shas))

	// All 4 browsers.
	run.Revision = "abcdef0123"
	run.CreatedAt = time.Now()
	for _, browser := range browserNames {
		run.BrowserName = browser
		datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "TestRun", nil), &run)
	}
	shas, _ = getCompleteRunSHAs(ctx, nil, nil)
	assert.Equal(t, []string{"abcdef0123"}, shas)

	// Another (earlier) run, all 4 browsers.
	run.Revision = "abcdef9999"
	run.CreatedAt = time.Now().AddDate(0, 0, -5)
	for _, browser := range browserNames {
		run.BrowserName = browser
		datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "TestRun", nil), &run)
	}
	shas, _ = getCompleteRunSHAs(ctx, nil, nil)
	assert.Equal(t, []string{"abcdef0123", "abcdef9999"}, shas)
	// Limit 1
	one := 1
	shas, _ = getCompleteRunSHAs(ctx, nil, &one)
	assert.Equal(t, []string{"abcdef0123"}, shas)
	// From yesterday.
	from := time.Now().AddDate(0, 0, -1)
	shas, _ = getCompleteRunSHAs(ctx, &from, nil)
	assert.Equal(t, []string{"abcdef0123"}, shas)
}
