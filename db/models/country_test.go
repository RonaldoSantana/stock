// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testCountries(t *testing.T) {
	t.Parallel()

	query := Countries(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testCountriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	country := &Country{}
	if err = randomize.Struct(seed, country, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = country.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Countries(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCountriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	country := &Country{}
	if err = randomize.Struct(seed, country, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Countries(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Countries(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCountriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	country := &Country{}
	if err = randomize.Struct(seed, country, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CountrySlice{country}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Countries(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testCountriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	country := &Country{}
	if err = randomize.Struct(seed, country, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := CountryExists(tx, country.ID)
	if err != nil {
		t.Errorf("Unable to check if Country exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CountryExistsG to return true, but got false.")
	}
}
func testCountriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	country := &Country{}
	if err = randomize.Struct(seed, country, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Insert(tx); err != nil {
		t.Error(err)
	}

	countryFound, err := FindCountry(tx, country.ID)
	if err != nil {
		t.Error(err)
	}

	if countryFound == nil {
		t.Error("want a record, got nil")
	}
}
func testCountriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	country := &Country{}
	if err = randomize.Struct(seed, country, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Countries(tx).Bind(country); err != nil {
		t.Error(err)
	}
}

func testCountriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	country := &Country{}
	if err = randomize.Struct(seed, country, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Countries(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCountriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	countryOne := &Country{}
	countryTwo := &Country{}
	if err = randomize.Struct(seed, countryOne, countryDBTypes, false, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}
	if err = randomize.Struct(seed, countryTwo, countryDBTypes, false, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = countryOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = countryTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Countries(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCountriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	countryOne := &Country{}
	countryTwo := &Country{}
	if err = randomize.Struct(seed, countryOne, countryDBTypes, false, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}
	if err = randomize.Struct(seed, countryTwo, countryDBTypes, false, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = countryOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = countryTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Countries(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func countryBeforeInsertHook(e boil.Executor, o *Country) error {
	*o = Country{}
	return nil
}

func countryAfterInsertHook(e boil.Executor, o *Country) error {
	*o = Country{}
	return nil
}

func countryAfterSelectHook(e boil.Executor, o *Country) error {
	*o = Country{}
	return nil
}

func countryBeforeUpdateHook(e boil.Executor, o *Country) error {
	*o = Country{}
	return nil
}

func countryAfterUpdateHook(e boil.Executor, o *Country) error {
	*o = Country{}
	return nil
}

func countryBeforeDeleteHook(e boil.Executor, o *Country) error {
	*o = Country{}
	return nil
}

func countryAfterDeleteHook(e boil.Executor, o *Country) error {
	*o = Country{}
	return nil
}

func countryBeforeUpsertHook(e boil.Executor, o *Country) error {
	*o = Country{}
	return nil
}

func countryAfterUpsertHook(e boil.Executor, o *Country) error {
	*o = Country{}
	return nil
}

func testCountriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Country{}
	o := &Country{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, countryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Country object: %s", err)
	}

	AddCountryHook(boil.BeforeInsertHook, countryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	countryBeforeInsertHooks = []CountryHook{}

	AddCountryHook(boil.AfterInsertHook, countryAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	countryAfterInsertHooks = []CountryHook{}

	AddCountryHook(boil.AfterSelectHook, countryAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	countryAfterSelectHooks = []CountryHook{}

	AddCountryHook(boil.BeforeUpdateHook, countryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	countryBeforeUpdateHooks = []CountryHook{}

	AddCountryHook(boil.AfterUpdateHook, countryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	countryAfterUpdateHooks = []CountryHook{}

	AddCountryHook(boil.BeforeDeleteHook, countryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	countryBeforeDeleteHooks = []CountryHook{}

	AddCountryHook(boil.AfterDeleteHook, countryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	countryAfterDeleteHooks = []CountryHook{}

	AddCountryHook(boil.BeforeUpsertHook, countryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	countryBeforeUpsertHooks = []CountryHook{}

	AddCountryHook(boil.AfterUpsertHook, countryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	countryAfterUpsertHooks = []CountryHook{}
}
func testCountriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	country := &Country{}
	if err = randomize.Struct(seed, country, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Countries(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCountriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	country := &Country{}
	if err = randomize.Struct(seed, country, countryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Insert(tx, countryColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Countries(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCountryToManyAddresses(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Country
	var b, c Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, addressDBTypes, false, addressColumnsWithDefault...)
	randomize.Struct(seed, &c, addressDBTypes, false, addressColumnsWithDefault...)

	b.CountryID.Valid = true
	c.CountryID.Valid = true
	b.CountryID.Int = a.ID
	c.CountryID.Int = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	address, err := a.Addresses(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range address {
		if v.CountryID.Int == b.CountryID.Int {
			bFound = true
		}
		if v.CountryID.Int == c.CountryID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CountrySlice{&a}
	if err = a.L.LoadAddresses(tx, false, (*[]*Country)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Addresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Addresses = nil
	if err = a.L.LoadAddresses(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Addresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", address)
	}
}

func testCountryToManyGeoZoneMatrices(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Country
	var b, c GeoZoneMatrix

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, geoZoneMatrixDBTypes, false, geoZoneMatrixColumnsWithDefault...)
	randomize.Struct(seed, &c, geoZoneMatrixDBTypes, false, geoZoneMatrixColumnsWithDefault...)

	b.CountryID = a.ID
	c.CountryID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	geoZoneMatrix, err := a.GeoZoneMatrices(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range geoZoneMatrix {
		if v.CountryID == b.CountryID {
			bFound = true
		}
		if v.CountryID == c.CountryID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CountrySlice{&a}
	if err = a.L.LoadGeoZoneMatrices(tx, false, (*[]*Country)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GeoZoneMatrices); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.GeoZoneMatrices = nil
	if err = a.L.LoadGeoZoneMatrices(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GeoZoneMatrices); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", geoZoneMatrix)
	}
}

func testCountryToManyRegions(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Country
	var b, c Region

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, regionDBTypes, false, regionColumnsWithDefault...)
	randomize.Struct(seed, &c, regionDBTypes, false, regionColumnsWithDefault...)

	b.CountryID = a.ID
	c.CountryID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	region, err := a.Regions(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range region {
		if v.CountryID == b.CountryID {
			bFound = true
		}
		if v.CountryID == c.CountryID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CountrySlice{&a}
	if err = a.L.LoadRegions(tx, false, (*[]*Country)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Regions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Regions = nil
	if err = a.L.LoadRegions(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Regions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", region)
	}
}

func testCountryToManyAddOpAddresses(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Country
	var b, c, d, e Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, countryDBTypes, false, strmangle.SetComplement(countryPrimaryKeyColumns, countryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Address{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Address{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAddresses(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.CountryID.Int {
			t.Error("foreign key was wrong value", a.ID, first.CountryID.Int)
		}
		if a.ID != second.CountryID.Int {
			t.Error("foreign key was wrong value", a.ID, second.CountryID.Int)
		}

		if first.R.Country != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Country != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Addresses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Addresses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Addresses(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCountryToManySetOpAddresses(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Country
	var b, c, d, e Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, countryDBTypes, false, strmangle.SetComplement(countryPrimaryKeyColumns, countryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Address{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetAddresses(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Addresses(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetAddresses(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Addresses(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.CountryID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.CountryID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.ID != d.CountryID.Int {
		t.Error("foreign key was wrong value", a.ID, d.CountryID.Int)
	}
	if a.ID != e.CountryID.Int {
		t.Error("foreign key was wrong value", a.ID, e.CountryID.Int)
	}

	if b.R.Country != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Country != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Country != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Country != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Addresses[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Addresses[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testCountryToManyRemoveOpAddresses(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Country
	var b, c, d, e Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, countryDBTypes, false, strmangle.SetComplement(countryPrimaryKeyColumns, countryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Address{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddAddresses(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Addresses(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveAddresses(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Addresses(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.CountryID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.CountryID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Country != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Country != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Country != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Country != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Addresses) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Addresses[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Addresses[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testCountryToManyAddOpGeoZoneMatrices(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Country
	var b, c, d, e GeoZoneMatrix

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, countryDBTypes, false, strmangle.SetComplement(countryPrimaryKeyColumns, countryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*GeoZoneMatrix{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, geoZoneMatrixDBTypes, false, strmangle.SetComplement(geoZoneMatrixPrimaryKeyColumns, geoZoneMatrixColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*GeoZoneMatrix{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddGeoZoneMatrices(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.CountryID {
			t.Error("foreign key was wrong value", a.ID, first.CountryID)
		}
		if a.ID != second.CountryID {
			t.Error("foreign key was wrong value", a.ID, second.CountryID)
		}

		if first.R.Country != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Country != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.GeoZoneMatrices[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.GeoZoneMatrices[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.GeoZoneMatrices(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testCountryToManyAddOpRegions(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Country
	var b, c, d, e Region

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, countryDBTypes, false, strmangle.SetComplement(countryPrimaryKeyColumns, countryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Region{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, regionDBTypes, false, strmangle.SetComplement(regionPrimaryKeyColumns, regionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Region{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRegions(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.CountryID {
			t.Error("foreign key was wrong value", a.ID, first.CountryID)
		}
		if a.ID != second.CountryID {
			t.Error("foreign key was wrong value", a.ID, second.CountryID)
		}

		if first.R.Country != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Country != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Regions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Regions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Regions(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCountriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	country := &Country{}
	if err = randomize.Struct(seed, country, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = country.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testCountriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	country := &Country{}
	if err = randomize.Struct(seed, country, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := CountrySlice{country}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testCountriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	country := &Country{}
	if err = randomize.Struct(seed, country, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Countries(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	countryDBTypes = map[string]string{`ID`: `int`, `IsoCode2`: `varchar`, `IsoCode3`: `varchar`, `Name`: `varchar`, `Status`: `tinyint`}
	_              = bytes.MinRead
)

func testCountriesUpdate(t *testing.T) {
	t.Parallel()

	if len(countryColumns) == len(countryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	country := &Country{}
	if err = randomize.Struct(seed, country, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Countries(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, country, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	if err = country.Update(tx); err != nil {
		t.Error(err)
	}
}

func testCountriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(countryColumns) == len(countryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	country := &Country{}
	if err = randomize.Struct(seed, country, countryDBTypes, true, countryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Countries(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, country, countryDBTypes, true, countryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(countryColumns, countryPrimaryKeyColumns) {
		fields = countryColumns
	} else {
		fields = strmangle.SetComplement(
			countryColumns,
			countryPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(country))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := CountrySlice{country}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testCountriesUpsert(t *testing.T) {
	t.Parallel()

	if len(countryColumns) == len(countryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	country := Country{}
	if err = randomize.Struct(seed, &country, countryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = country.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Country: %s", err)
	}

	count, err := Countries(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &country, countryDBTypes, false, countryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Country struct: %s", err)
	}

	if err = country.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Country: %s", err)
	}

	count, err = Countries(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
