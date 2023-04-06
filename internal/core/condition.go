/*
 * Copyright 2023 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import "github.com/bytedance/arishem/internal/typedef"

type arishemCondition struct {
	passed    bool
	left      interface{}
	leftExpr  string
	right     interface{}
	rightExpr string
	operator  string
	error     error
}

func newArishemCondition(
	passed bool,
	left interface{},
	leftExpr string,
	right interface{},
	rightExpr string,
	operator string,
	error error,
) typedef.JudgeNode {
	return &arishemCondition{
		passed:    passed,
		left:      left,
		leftExpr:  leftExpr,
		right:     right,
		rightExpr: rightExpr,
		operator:  operator,
		error:     error,
	}
}

func (a *arishemCondition) Passed() bool {
	return a.passed
}

func (a *arishemCondition) Left() interface{} {
	return a.left
}

func (a *arishemCondition) LeftExpr() string {
	return a.leftExpr
}

func (a *arishemCondition) Right() interface{} {
	return a.right
}

func (a *arishemCondition) RightExpr() string {
	return a.rightExpr
}

func (a *arishemCondition) Operator() string {
	return a.operator
}

func (a *arishemCondition) Error() error {
	return a.error
}
