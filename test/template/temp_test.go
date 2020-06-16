// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package template

import (
	"github.com/xfali/gobatis"
	tmp2 "github.com/xfali/gobatis/parsing/template"
	"os"
	"testing"
	"text/template"
	"time"
)

type TestTable struct {
	Id       int
	UserName string
	Password string
	Status   int
	Time     time.Time
}

var driverName = "mysql"

func TestTemplate(t *testing.T) {
	tpl, err := template.ParseFiles("./sql.tpl")
	if err != nil {
		t.Fatal(err)
	}

	s := tpl.Templates()
	for _, v := range s {
		t.Log(v.Name())
	}

	var param = TestTable{Id: 1, UserName: "user", Password: "pw"}
	t.Run("select", func(t *testing.T) {
		tpl = tpl.Lookup("selectTestTable")
		if tpl == nil {
			t.Fatal("not found")
		}

		err = tpl.Execute(os.Stdout, param)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("insert", func(t *testing.T) {
		tpl = tpl.Lookup("insertTestTable")
		if tpl == nil {
			t.Fatal("not found")
		}

		err = tpl.Execute(os.Stdout, param)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("update", func(t *testing.T) {
		tpl = tpl.Lookup("updateTestTable")
		if tpl == nil {
			t.Fatal("not found")
		}

		err = tpl.Execute(os.Stdout, param)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("delete", func(t *testing.T) {
		tpl = tpl.Lookup("deleteTestTable")
		if tpl == nil {
			t.Fatal("not found")
		}

		err = tpl.Execute(os.Stdout, param)
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestParser(t *testing.T) {
	mgr := tmp2.NewManager()
	mgr.RegisterFile("./sql.tpl")
	var param = TestTable{Id: 1, UserName: "user", Password: "pw", Time: time.Now()}

	t.Run("select", func(t *testing.T) {
		tmp, _ := mgr.FindSqlParser("selectTestTable")
		md, err := tmp.ParseMetadata(driverName, param)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(md)
	})

	t.Run("insert", func(t *testing.T) {
		tmp, _ := mgr.FindSqlParser("insertTestTable")
		md, err := tmp.ParseMetadata(driverName, param)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(md)
	})

	t.Run("insertBatch", func(t *testing.T) {
		tmp, _ := mgr.FindSqlParser("insertBatchTestTable")
		md, err := tmp.ParseMetadata(driverName, []TestTable{
			{Id: 1, UserName: "user1", Password: "pw1"},
			{Id: 2, UserName: "user2", Password: "pw2"},
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Log(md)
	})

	t.Run("update", func(t *testing.T) {
		tmp, _ := mgr.FindSqlParser("updateTestTable")
		md, err := tmp.ParseMetadata(driverName, param)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(md)
	})

	t.Run("delete", func(t *testing.T) {
		tmp, _ := mgr.FindSqlParser("deleteTestTable")
		md, err := tmp.ParseMetadata(driverName, param)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(md)
	})
}

func TestParser2(t *testing.T) {
	gobatis.RegisterTemplateFile("./sql.tpl")
	t.Run("select", func(t *testing.T) {
		p, ok := gobatis.FindTemplateSqlParser("selectTestTable")
		if !ok {
			t.Fatal(ok)
		}
		md, err := p.ParseMetadata("mysql", TestTable{Id: 1, UserName: "user", Password: "pw", Status: 10})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(md)
	})

	t.Run("insert", func(t *testing.T) {
		p, ok := gobatis.FindTemplateSqlParser("insertTestTable")
		if !ok {
			t.Fatal(ok)
		}
		md, err := p.ParseMetadata("mysql", TestTable{Id: 1, UserName: "user", Password: "pw", Status: 10})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(md)
	})

	t.Run("insertBatch", func(t *testing.T) {
		p, ok := gobatis.FindTemplateSqlParser("insertBatchTestTable")
		if !ok {
			t.Fatal(ok)
		}
		md, err := p.ParseMetadata("mysql", []TestTable{
			{Id: 11, UserName: "user11", Password: "pw11"},
			{Id: 12, UserName: "user12", Password: "pw12"},
		})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(md)
	})

	t.Run("update", func(t *testing.T) {
		p, ok := gobatis.FindTemplateSqlParser("updateTestTable")
		if !ok {
			t.Fatal(ok)
		}
		md, err := p.ParseMetadata("mysql", TestTable{Id: 1, UserName: "user", Password: "pw", Status: 10})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(md)
	})

	t.Run("delete", func(t *testing.T) {
		p, ok := gobatis.FindTemplateSqlParser("deleteTestTable")
		if !ok {
			t.Fatal(ok)
		}
		md, err := p.ParseMetadata("mysql", TestTable{Id: 1, UserName: "user", Password: "pw", Status: 10})
		if err != nil {
			t.Fatal(err)
		}

		t.Log(md)
	})
}
