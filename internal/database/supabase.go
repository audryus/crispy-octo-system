package database

import (
	"log"

	appConf "github.com/audryus/crispy-octo-system/configs"
	"github.com/nedpals/supabase-go"
)

var Supa *supabase.Client

func InitSupabase(conf appConf.Supabase) {
	Supa = supabase.CreateClient(conf.Url, conf.AnonKey)
	log.Print("Supabase connected")
}
