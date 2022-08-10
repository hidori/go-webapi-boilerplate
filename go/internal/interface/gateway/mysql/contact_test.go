package mysql

import (
	"testing"

	"github.com/ahmetb/go-linq/v3"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hidori/go-webapi-boilerplate/go/internal/domain/model"
	"github.com/hidori/go-webapi-boilerplate/go/test/fixture"
	"github.com/hidori/go-webapi-boilerplate/go/test/xtest"
	"github.com/hidori/go-webapi-boilerplate/go/test/xtestgorm"
	"github.com/stretchr/testify/assert"
)

func TestContactRepository_GetList(t *testing.T) {
	type testCase = xtest.Case[testContext, any, any, []model.Contact]
	tests := []testCase{
		{
			Name:    "正常系:すべてのレコードを取得する",
			Context: newTestContext(),
			Up: func(tt testCase) error {
				return xtestgorm.InsertAll(tt.Context.db, []interface{}{
					fixture.NewContact(1001),
					fixture.NewContact(1002),
					fixture.NewContact(1003),
				})
			},
			Want: []model.Contact{
				*fixture.NewContact(1001),
				*fixture.NewContact(1002),
				*fixture.NewContact(1003),
			},
		},
		{
			Name:    "正常系:レコードが存在しない時は空の配列を返す。",
			Context: newTestContext(),
			Up: func(tt testCase) error {
				return nil
			},
			Want: []model.Contact{},
		},
	}
	for _, tt := range tests {
		err := deleteAll(tt.Context.db)
		if !assert.NoError(t, err, "fail to deleteAll()") {
			continue
		}

		xtest.Run(t, tt, func(t *testing.T, tt testCase) {
			rp := NewContactRepository()
			got, err := rp.GetList(tt.Context.db)
			if err != nil {
				if tt.WantError {
					assert.Contains(t, err.Error(), tt.Error)
				} else {
					t.Errorf("fail to GetList(): err=%v, tt.args=%v", err, tt.Args)
				}
			} else {
				xtest.Equal(t, tt.Want, got, cmpopts.IgnoreFields(model.Contact{}, "CreatedAt", "UpdatedAt"))
			}
		})
	}
}

func TestContactRepository_GetByID(t *testing.T) {
	type testArgs struct {
		contactID int
	}
	type testCase = xtest.Case[testContext, any, testArgs, *model.Contact]
	tests := []testCase{
		{
			Name:    "正常系:指定された条件にマッチする１件のレコードを取得する。",
			Context: newTestContext(),
			Args: testArgs{
				contactID: 1002,
			},
			Up: func(tt testCase) error {
				return xtestgorm.InsertAll(tt.Context.db, []interface{}{
					fixture.NewContact(1001),
					fixture.NewContact(1002),
					fixture.NewContact(1003),
				})
			},
			Want: fixture.NewContact(1002),
		},
		{
			Name:    "正常系:レコードが存在しない時は nil を返す。",
			Context: newTestContext(),
			Args: testArgs{
				contactID: 1,
			},
			Up: func(tt testCase) error {
				return xtestgorm.InsertAll(tt.Context.db, []interface{}{
					fixture.NewContact(1001),
					fixture.NewContact(1002),
					fixture.NewContact(1003),
				})
			},
			Want: nil,
		},
	}
	for _, tt := range tests {
		err := deleteAll(tt.Context.db)
		if !assert.NoError(t, err, "fail to deleteAll()") {
			continue
		}

		xtest.Run(t, tt, func(t *testing.T, tt testCase) {
			rp := NewContactRepository()
			got, err := rp.GetByID(tt.Context.db, tt.Args.contactID)
			if err != nil {
				if tt.WantError {
					assert.Contains(t, err.Error(), tt.Error)
				} else {
					t.Errorf("fail to GetByID(): err=%v, tt.args=%v", err, tt.Args)
				}
			} else {
				xtest.Equal(t, tt.Want, got, cmpopts.IgnoreFields(model.Contact{}, "CreatedAt", "UpdatedAt"))
			}
		})
	}
}

func TestContactRepository_AddOrUpdate(t *testing.T) {
	type testArgs struct {
		model *model.Contact
	}
	type testCase = xtest.Case[testContext, any, testArgs, *model.Contact]
	tests := []testCase{
		{
			Name:    "正常系:１件のレコードを追加して、そのレコードを返す。",
			Context: newTestContext(),
			Args: testArgs{
				model: fixture.NewContact(0),
			},
			Up: func(tt testCase) error {
				return xtestgorm.InsertAll(tt.Context.db, []interface{}{
					fixture.NewContact(1001),
					fixture.NewContact(1002),
					fixture.NewContact(1003),
				})
			},
			Want: fixture.NewContact(0),
			More: func(t *testing.T, tt testCase, got interface{}) {
				contacts, err := xtestgorm.SelectAll[model.Contact](tt.Context.db)
				if err != nil {
					t.Errorf("fail to xtestgorm.SelectAll[): err=%v", err)
					return
				}

				assert.Equal(t, 4, len(contacts), "len(contacts)")

				var ids []int
				linq.From(contacts).
					Select(func(i interface{}) interface{} { return i.(model.Contact).ContactID }).
					Except(linq.From([]int{1001, 1002, 1003})).
					ToSlice(&ids)

				assert.Equal(t, 1, len(ids))
				assert.NotEqual(t, ids[0], 0)
			},
		},
		{
			Name:    "正常系:１件のレコードを更新して、そのレコードを返す。",
			Context: newTestContext(),
			Args: testArgs{
				model: func() *model.Contact {
					fixture := fixture.NewContact(1002)
					return &model.Contact{
						ContactID:      fixture.ContactID,
						FamilyName:     e(fixture.FamilyName),
						FirstName:      e(fixture.FirstName),
						FamilyNameKana: e(fixture.FamilyNameKana),
						FirstNameKana:  e(fixture.FirstNameKana),
						PhoneNumber:    e(fixture.PhoneNumber),
						PostalCode:     e(fixture.PostalCode),
						PrefectureCode: e(fixture.PrefectureCode),
						CityCode:       e(fixture.CityCode),
						AddressLine1:   e(fixture.AddressLine1),
						AddressLine2:   e(fixture.AddressLine2),
					}
				}(),
			},
			Up: func(tt testCase) error {
				return xtestgorm.InsertAll(tt.Context.db, []interface{}{
					fixture.NewContact(1001),
					fixture.NewContact(1002),
					fixture.NewContact(1003),
				})
			},
			Want: func() *model.Contact {
				fixture := fixture.NewContact(1002)
				return &model.Contact{
					ContactID:      fixture.ContactID,
					FamilyName:     e(fixture.FamilyName),
					FirstName:      e(fixture.FirstName),
					FamilyNameKana: e(fixture.FamilyNameKana),
					FirstNameKana:  e(fixture.FirstNameKana),
					PhoneNumber:    e(fixture.PhoneNumber),
					PostalCode:     e(fixture.PostalCode),
					PrefectureCode: e(fixture.PrefectureCode),
					CityCode:       e(fixture.CityCode),
					AddressLine1:   e(fixture.AddressLine1),
					AddressLine2:   e(fixture.AddressLine2),
				}
			}(),
			More: func(t *testing.T, tt testCase, got interface{}) {
				contacts, err := xtestgorm.SelectAll[model.Contact](tt.Context.db)
				if err != nil {
					t.Errorf("fail to xtestgorm.SelectAll[): err=%v", err)
					return
				}

				assert.Equal(t, 3, len(contacts), "len(contacts)")

				var ids []int
				linq.From(contacts).
					Select(func(i interface{}) interface{} { return i.(model.Contact).ContactID }).
					Except(linq.From([]int{1001, 1003})).
					ToSlice(&ids)

				assert.Equal(t, 1, len(ids))
				xtest.Equal(t, tt.Want.ContactID, ids[0])
			}},
	}
	for _, tt := range tests {
		err := deleteAll(tt.Context.db)
		if !assert.NoError(t, err, "fail to deleteAll()") {
			continue
		}

		xtest.Run(t, tt, func(t *testing.T, tt testCase) {
			rp := NewContactRepository()
			got, err := rp.AddOrUpdate(tt.Context.db, tt.Args.model)
			if err != nil {
				if tt.WantError {
					assert.Contains(t, err.Error(), tt.Error)
				} else {
					t.Errorf("fail to AddOrUpdate(): err=%v, tt.args=%v", err, tt.Args)
				}
			} else {
				if !xtest.Equal(t, tt.Want, got, cmpopts.IgnoreFields(model.Contact{}, "ContactID", "CreatedAt", "UpdatedAt")) {
					return
				}

				if tt.More != nil {
					tt.More(t, tt, got)
				}
			}
		})
	}
}

func TestContactRepository_DeleteByID(t *testing.T) {
	type testArgs struct {
		contactID int
	}
	type testCase = xtest.Case[testContext, any, testArgs, *model.Contact]
	tests := []testCase{
		{
			Name:    "正常系:指定された条件にマッチする１件のレコードを削除する。",
			Context: newTestContext(),
			Args: testArgs{
				contactID: 1002,
			},
			Up: func(tt testCase) error {
				return xtestgorm.InsertAll(tt.Context.db, []interface{}{
					fixture.NewContact(1001),
					fixture.NewContact(1002),
					fixture.NewContact(1003),
				})
			},
			More: func(t *testing.T, tt testCase, got interface{}) {
				var contacts []model.Contact
				contacts, err := xtestgorm.SelectAll[model.Contact](tt.Context.db)
				if err != nil {
					t.Errorf("fail to xtestgorm.SelectAll[): err=%v", err)
					return
				}

				assert.Equal(t, 2, len(contacts), "len(contacts)")

				var ids []int
				linq.From(contacts).
					Select(func(i interface{}) interface{} { return i.(model.Contact).ContactID }).
					Except(linq.From([]int{1001, 1003})).
					ToSlice(&ids)

				assert.Equal(t, 0, len(ids))
			},
		},
		{
			Name:    "正常系:条件にマッチする１件のレコードが存在しない時はエラーを返す。",
			Context: newTestContext(),
			Args: testArgs{
				contactID: 1,
			},
			Up: func(tt testCase) error {
				return xtestgorm.InsertAll(tt.Context.db, []interface{}{
					fixture.NewContact(1001),
					fixture.NewContact(1002),
					fixture.NewContact(1003),
				})
			},
			WantError: true,
			Error:     "rowsAffected is not 1: rowsAffected=0",
			More: func(t *testing.T, tt testCase, got interface{}) {
				contacts, err := xtestgorm.SelectAll[model.Contact](tt.Context.db)
				if err != nil {
					t.Errorf("fail to xtestgorm.SelectAll[): err=%v", err)
					return
				}

				assert.Equal(t, 3, len(contacts), "len(contacts)")

				var ids []int
				linq.From(contacts).
					Select(func(i interface{}) interface{} { return i.(model.Contact).ContactID }).
					Except(linq.From([]int{1001, 1002, 1003})).
					ToSlice(&ids)

				assert.Equal(t, 0, len(ids))
			}},
	}
	for _, tt := range tests {
		err := deleteAll(tt.Context.db)
		if !assert.NoError(t, err, "fail to deleteAll()") {
			continue
		}

		xtest.Run(t, tt, func(t *testing.T, tt testCase) {
			rp := NewContactRepository()
			err = rp.DeleteByID(tt.Context.db, tt.Args.contactID)
			if err != nil {
				if tt.WantError {
					assert.Contains(t, err.Error(), tt.Error)
				} else {
					t.Errorf("fail to DeleteByID(): err=%v, tt.args=%v", err, tt.Args)
				}
			}
		})
	}
}
