package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mharner33/bookings/internal/config"
	"github.com/mharner33/bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

// Create methods to satisfy ResponseWriter interface
type myWriter struct{}

func (mw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (mw *myWriter) WriteHeader(i int) {}

func (mw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})
	// Set this to true when running in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	app = &testApp

	os.Exit(m.Run())
}
