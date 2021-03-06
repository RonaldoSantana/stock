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

func testTaxRates(t *testing.T) {
	t.Parallel()

	query := TaxRates(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testTaxRatesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxRate := &TaxRate{}
	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = taxRate.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := TaxRates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTaxRatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxRate := &TaxRate{}
	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = TaxRates(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := TaxRates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTaxRatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxRate := &TaxRate{}
	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TaxRateSlice{taxRate}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := TaxRates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testTaxRatesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxRate := &TaxRate{}
	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := TaxRateExists(tx, taxRate.ID)
	if err != nil {
		t.Errorf("Unable to check if TaxRate exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TaxRateExistsG to return true, but got false.")
	}
}
func testTaxRatesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxRate := &TaxRate{}
	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Insert(tx); err != nil {
		t.Error(err)
	}

	taxRateFound, err := FindTaxRate(tx, taxRate.ID)
	if err != nil {
		t.Error(err)
	}

	if taxRateFound == nil {
		t.Error("want a record, got nil")
	}
}
func testTaxRatesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxRate := &TaxRate{}
	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = TaxRates(tx).Bind(taxRate); err != nil {
		t.Error(err)
	}
}

func testTaxRatesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxRate := &TaxRate{}
	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := TaxRates(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTaxRatesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxRateOne := &TaxRate{}
	taxRateTwo := &TaxRate{}
	if err = randomize.Struct(seed, taxRateOne, taxRateDBTypes, false, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}
	if err = randomize.Struct(seed, taxRateTwo, taxRateDBTypes, false, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRateOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = taxRateTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := TaxRates(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTaxRatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	taxRateOne := &TaxRate{}
	taxRateTwo := &TaxRate{}
	if err = randomize.Struct(seed, taxRateOne, taxRateDBTypes, false, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}
	if err = randomize.Struct(seed, taxRateTwo, taxRateDBTypes, false, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRateOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = taxRateTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TaxRates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func taxRateBeforeInsertHook(e boil.Executor, o *TaxRate) error {
	*o = TaxRate{}
	return nil
}

func taxRateAfterInsertHook(e boil.Executor, o *TaxRate) error {
	*o = TaxRate{}
	return nil
}

func taxRateAfterSelectHook(e boil.Executor, o *TaxRate) error {
	*o = TaxRate{}
	return nil
}

func taxRateBeforeUpdateHook(e boil.Executor, o *TaxRate) error {
	*o = TaxRate{}
	return nil
}

func taxRateAfterUpdateHook(e boil.Executor, o *TaxRate) error {
	*o = TaxRate{}
	return nil
}

func taxRateBeforeDeleteHook(e boil.Executor, o *TaxRate) error {
	*o = TaxRate{}
	return nil
}

func taxRateAfterDeleteHook(e boil.Executor, o *TaxRate) error {
	*o = TaxRate{}
	return nil
}

func taxRateBeforeUpsertHook(e boil.Executor, o *TaxRate) error {
	*o = TaxRate{}
	return nil
}

func taxRateAfterUpsertHook(e boil.Executor, o *TaxRate) error {
	*o = TaxRate{}
	return nil
}

func testTaxRatesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &TaxRate{}
	o := &TaxRate{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, taxRateDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TaxRate object: %s", err)
	}

	AddTaxRateHook(boil.BeforeInsertHook, taxRateBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	taxRateBeforeInsertHooks = []TaxRateHook{}

	AddTaxRateHook(boil.AfterInsertHook, taxRateAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	taxRateAfterInsertHooks = []TaxRateHook{}

	AddTaxRateHook(boil.AfterSelectHook, taxRateAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	taxRateAfterSelectHooks = []TaxRateHook{}

	AddTaxRateHook(boil.BeforeUpdateHook, taxRateBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	taxRateBeforeUpdateHooks = []TaxRateHook{}

	AddTaxRateHook(boil.AfterUpdateHook, taxRateAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	taxRateAfterUpdateHooks = []TaxRateHook{}

	AddTaxRateHook(boil.BeforeDeleteHook, taxRateBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	taxRateBeforeDeleteHooks = []TaxRateHook{}

	AddTaxRateHook(boil.AfterDeleteHook, taxRateAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	taxRateAfterDeleteHooks = []TaxRateHook{}

	AddTaxRateHook(boil.BeforeUpsertHook, taxRateBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	taxRateBeforeUpsertHooks = []TaxRateHook{}

	AddTaxRateHook(boil.AfterUpsertHook, taxRateAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	taxRateAfterUpsertHooks = []TaxRateHook{}
}
func testTaxRatesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxRate := &TaxRate{}
	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TaxRates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTaxRatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxRate := &TaxRate{}
	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Insert(tx, taxRateColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := TaxRates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTaxRateToOneGeoZoneUsingGeoZone(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local TaxRate
	var foreign GeoZone

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, geoZoneDBTypes, true, geoZoneColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GeoZone struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.GeoZoneID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.GeoZone(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TaxRateSlice{&local}
	if err = local.L.LoadGeoZone(tx, false, (*[]*TaxRate)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.GeoZone == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.GeoZone = nil
	if err = local.L.LoadGeoZone(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.GeoZone == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTaxRateToOneSetOpGeoZoneUsingGeoZone(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a TaxRate
	var b, c GeoZone

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taxRateDBTypes, false, strmangle.SetComplement(taxRatePrimaryKeyColumns, taxRateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, geoZoneDBTypes, false, strmangle.SetComplement(geoZonePrimaryKeyColumns, geoZoneColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, geoZoneDBTypes, false, strmangle.SetComplement(geoZonePrimaryKeyColumns, geoZoneColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*GeoZone{&b, &c} {
		err = a.SetGeoZone(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.GeoZone != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TaxRates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.GeoZoneID != x.ID {
			t.Error("foreign key was wrong value", a.GeoZoneID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.GeoZoneID))
		reflect.Indirect(reflect.ValueOf(&a.GeoZoneID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.GeoZoneID != x.ID {
			t.Error("foreign key was wrong value", a.GeoZoneID, x.ID)
		}
	}
}
func testTaxRatesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxRate := &TaxRate{}
	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = taxRate.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testTaxRatesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxRate := &TaxRate{}
	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TaxRateSlice{taxRate}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testTaxRatesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxRate := &TaxRate{}
	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := TaxRates(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	taxRateDBTypes = map[string]string{`CreatedAt`: `timestamp`, `GeoZoneID`: `int`, `ID`: `int`, `Name`: `varchar`, `Rate`: `double`, `Type`: `enum('value','percent')`, `UpdatedAt`: `timestamp`}
	_              = bytes.MinRead
)

func testTaxRatesUpdate(t *testing.T) {
	t.Parallel()

	if len(taxRateColumns) == len(taxRatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	taxRate := &TaxRate{}
	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TaxRates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	if err = taxRate.Update(tx); err != nil {
		t.Error(err)
	}
}

func testTaxRatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(taxRateColumns) == len(taxRatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	taxRate := &TaxRate{}
	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TaxRates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, taxRate, taxRateDBTypes, true, taxRatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(taxRateColumns, taxRatePrimaryKeyColumns) {
		fields = taxRateColumns
	} else {
		fields = strmangle.SetComplement(
			taxRateColumns,
			taxRatePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(taxRate))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := TaxRateSlice{taxRate}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testTaxRatesUpsert(t *testing.T) {
	t.Parallel()

	if len(taxRateColumns) == len(taxRatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	taxRate := TaxRate{}
	if err = randomize.Struct(seed, &taxRate, taxRateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxRate.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert TaxRate: %s", err)
	}

	count, err := TaxRates(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &taxRate, taxRateDBTypes, false, taxRatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TaxRate struct: %s", err)
	}

	if err = taxRate.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert TaxRate: %s", err)
	}

	count, err = TaxRates(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
