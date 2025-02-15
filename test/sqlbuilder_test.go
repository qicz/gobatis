/*
 * Copyright (c) 2022, AcmeStack
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package test

import (
	"github.com/acmestack/gobatis/builder"
	"strings"
	"testing"
)

func TestSqlBuilderSelect(t *testing.T) {
	hook := func(f *builder.SQLFragment) *builder.SQLFragment {
		t.Log(f)
		return f
	}
	t.Run("once call", func(t *testing.T) {
		str := builder.Select("test1", "test2").
			Hook(hook).
			From("test_a").
			Hook(hook).
			Where("id = 1").
			And().
			Hook(hook).
			Where("name=2").
			Hook(hook).
			GroupBy("name").
			Hook(hook).
			OrderBy("name").
			Hook(hook).
			Desc().
			Hook(hook).
			Limit(5, 10).
			Hook(hook).
			String()
		t.Log(str)

		if strings.TrimSpace(str) != `SELECT test1, test2 FROM test_a WHERE id = 1 AND name=2 GROUP BY name ORDER BY name DESC LIMIT 5, 10` {
			t.FailNow()
		}
	})

	t.Run("multi call", func(t *testing.T) {
		str := builder.Select("A.test1", "B.test2").
			Select("B.test3").
			From("test_a AS A").
			From("test_b AS B").
			Where("id = 1").
			And().
			Where("name=2").
			GroupBy("name").
			OrderBy("name").
			Desc().
			Limit(5, 10).
			String()
		t.Log(str)

		if strings.TrimSpace(str) != `SELECT A.test1, B.test2, B.test3 FROM test_a AS A, test_b AS B WHERE id = 1 AND name=2 GROUP BY name ORDER BY name DESC LIMIT 5, 10` {
			t.FailNow()
		}
	})
}

func TestSqlBuilderInsert(t *testing.T) {
	str := builder.InsertInto("test_table").
		IntoColumns("a", "b").
		IntoColumns("c").
		IntoValues("#{0}, #{1}").
		IntoValues("#{3}").
		String()
	t.Log(str)

	if strings.TrimSpace(str) != `INSERT INTO test_table (a, b, c) VALUES(#{0}, #{1}, #{3})` {
		t.FailNow()
	}
}

func TestSqlBuilderUpdate(t *testing.T) {
	str := builder.Update("test_table").
		Set("a", "#{0}").
		Set("b", "#{1}").
		Where("id = #{3}").
		Or().
		Where("name = #{4}").
		String()
	t.Log(str)
	if strings.TrimSpace(str) != `UPDATE test_table SET a = #{0} , b = #{1} WHERE id = #{3} OR name = #{4}` {
		t.FailNow()
	}
}

func TestSqlBuilderDelete(t *testing.T) {
	str := builder.DeleteFrom("test_table").
		Where("id = #{3}").
		Or().
		Where("name = #{4}").
		String()
	t.Log(str)
	if strings.TrimSpace(str) != `DELETE FROM test_table WHERE id = #{3} OR name = #{4}` {
		t.FailNow()
	}
}

func TestSqlBuilderWhere(t *testing.T) {
	str := builder.DeleteFrom("test_table").
		Or().
		Where("name = #{4}").
		String()
	t.Log(str)

	if strings.TrimSpace(str) != `DELETE FROM test_table WHERE name = #{4}` {
		t.FailNow()
	}
}

func TestSqlBuilderError(t *testing.T) {
	f := builder.InsertInto("test_table")
	f.IntoColumns("a").IntoValues("#{0}").IntoValues("a").IntoColumns("a")

	t.Log(f.String())
}
