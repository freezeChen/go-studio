package config

import "testing"

func TestN(t *testing.T) {
	var s struct {
		Server struct {
			Http struct {
				Port int
			}
		}
	}

	config := New()
	config.LoadFileSource("../etc/conf.yml")
	if err := config.Scan(&s); err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("%+v",s)
}
