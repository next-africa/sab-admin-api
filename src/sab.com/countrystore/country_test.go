package countrystore

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
	"testing"
)

const COUNTRY_CODE = "CA"
const COUNTRY_NAME = "CANADA"
const NON_EXISTING_CODE = "MOCHO"

func TestSaveCountry(t *testing.T) {
	ctx, done := getNewContext(t)

	defer done()

	aCountry := Country{Code: COUNTRY_CODE, Name: COUNTRY_NAME}

	if err := SaveCountry(aCountry, ctx); err != nil {
		t.Fatal(err)
	}

	savedCountry := Country{}

	key := datastore.NewKey(ctx, "Country", "CA", 0, nil)
	datastore.Get(ctx, key, &savedCountry)

	assert.Equal(t, COUNTRY_CODE, aCountry.Code)
	assert.Equal(t, COUNTRY_NAME, aCountry.Name)
}

func TestGetCountry(t *testing.T) {
	ctx, done := getNewContext(t)

	defer done()

	testExistingCountryCase(t, ctx)

	testNonExistingCountryCase(t, ctx)

}
func testNonExistingCountryCase(t *testing.T, ctx context.Context) {

	_, err := GetCountryByCode(NON_EXISTING_CODE, ctx)

	if assert.Error(t, err) {
		assert.Equal(t, err, &CountryNotFoundError{countryCode: NON_EXISTING_CODE})
	}
}

func testExistingCountryCase(t *testing.T, ctx context.Context) {
	givenAnExistingCountry(ctx)

	country, err := GetCountryByCode(COUNTRY_CODE, ctx)

	if err != nil {
		t.Fatal(err)
	}

	if assert.NotNil(t, country) {
		assert.Equal(t, COUNTRY_CODE, country.Code)
	}
}

func givenAnExistingCountry(ctx context.Context) {
	existingCountry := Country{Code: COUNTRY_CODE, Name: COUNTRY_NAME}

	SaveCountry(existingCountry, ctx)
}

func getNewContext(t *testing.T) (ctx context.Context, done func()) {
	ctx, done, err := aetest.NewContext()

	if err != nil {
		t.Fatal(err)
	}
	return
}
