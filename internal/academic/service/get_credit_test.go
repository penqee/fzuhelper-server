/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/bytedance/mockey"
	"github.com/stretchr/testify/assert"

	"github.com/west2-online/fzuhelper-server/kitex_gen/model"
	meta "github.com/west2-online/fzuhelper-server/pkg/base/context"
	"github.com/west2-online/jwch"
)

func TestAcademicService_GetCredit(t *testing.T) {
	type testCase struct {
		name           string
		mockReturn     []*jwch.CreditStatistics
		mockError      error
		expectedResult []*jwch.CreditStatistics
		expectingError bool
	}

	expectedResult := []*jwch.CreditStatistics{
		{
			Type:  "Compulsory",
			Gain:  "4.0",
			Total: "8.0",
		},
	}

	testCases := []testCase{
		{
			name:           "GetCreditSuccess",
			mockReturn:     expectedResult,
			mockError:      nil,
			expectedResult: expectedResult,
		},
		{
			name:           "GetCreditFailure",
			mockReturn:     nil,
			mockError:      fmt.Errorf("get credit info fail"),
			expectedResult: nil,
			expectingError: true,
		},
	}

	defer mockey.UnPatchAll()
	for _, tc := range testCases {
		mockey.PatchConvey(tc.name, t, func() {
			mockey.Mock((*jwch.Student).GetCredit).Return(tc.mockReturn, tc.mockError).Build()
			mockey.Mock(meta.GetLoginData).To(func(ctx context.Context) (*model.LoginData, error) {
				return &model.LoginData{
					Id:      "1111111111111111111111111111111111",
					Cookies: "",
				}, nil
			}).Build()
			academicService := AcademicService{}
			result, err := academicService.GetCredit()
			if tc.expectingError {
				assert.Nil(t, result)
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "Get credit info fail")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}
		})
	}
}
