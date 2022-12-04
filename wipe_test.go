package wipe

import (
	"errors"
	"testing"
	"time"
)

type user struct {
	Name     string
	Birthday *time.Time
	Nickname string
	Role     string
	Age      int32
	FakeAge  *int32
	Notes    []string
	flags    []byte
	Tags     map[string]bool
}

type employee struct {
	AUser      *user
	User       user
	Name       string
	Birthday   *time.Time
	Nickname   *string
	Age        int64
	FakeAge    int
	EmployeeID int64
	DoubleAge  int32
	SuperRule  string
	Notes      []*string
	flags      []byte
	Tags       []int
	Cur        time.Time
	Err        error
}

func Test(t *testing.T) {
	birthday := time.Now()
	nickName := "nickName"
	e1 := employee{
		AUser: &user{
			Name: "e1._User.name",
		},
		User: user{
			Name: "2dsajiodasojid",
			Tags: map[string]bool{
				"e1.User.Tags": true,
			},
		},
		Name:      "123",
		Birthday:  &birthday,
		Nickname:  &nickName,
		Age:       23,
		FakeAge:   1231,
		DoubleAge: -123,
		Notes:     []*string{&nickName},
		flags:     []byte{12, 213},
		Tags:      []int{123, 213, 4, 34},
		Cur:       time.Now(),
		Err:       errors.New("asdasdasmkpmpo"),
	}
	t.Logf("%+v", e1)
	t.Logf("AUser %+v", e1.AUser)
	t.Log(*e1.Nickname)
	t.Log(e1.User.Tags)
	t.Log(e1.Err.Error())
	t.Log(Wipe(&e1))
	t.Log("wipe ", e1)
	t.Log("wipe AUser", e1.AUser)
	t.Log("wipe *e1.Nickname", *e1.Nickname)
	t.Log("wipe User.Tags", e1.User.Tags)
	t.Log("wipe User.Err", e1.Err.Error())

	//.....
	e1.Cur = time.Now()
	t.Log(e1.Cur)
	t.Log(Wipe(&e1.Cur))
	t.Log("wipe ", e1.Cur)

}

func Test_Ptr(t *testing.T) {
	type Person struct {
		Name *string
	}
	name := "123"
	p := Person{
		Name: &name,
	}

	t.Log("*p.Name", *p.Name)
	t.Log(Wipe(&p))
	t.Log("wipe *p.Name", *p.Name)

}

func Test_Ptr_Struct(t *testing.T) {
	type Person struct {
		Name *string
	}
	type School struct {
		P *Person
	}
	name := "12345"
	p := Person{
		Name: &name,
	}
	s := School{
		P: &p,
	}

	t.Log("*s.P.Name", *s.P.Name)
	t.Log(Wipe(&s))
	t.Log("wipe *s.P", *s.P)
	//t.Log("wipe *s.P.Name", *s.P.Name)

}
