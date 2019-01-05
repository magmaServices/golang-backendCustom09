package nucleus

import (
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/oiacow/fesl3/magma/server/auth"
	"gitlab.com/oiacow/fesl3/magma/tpl"
)

type dtSession struct {
	Token string
}

// nucleusAuthToken authorizes the client by assigning a `magma` cookie.
func (s *Controller) nucleusAuthToken(w http.ResponseWriter, r *http.Request) {
	if serverKey := auth.GetServerHeader(r); serverKey != "" {
		s.rdr.RenderXML(w, r, tpl.XmlSession, nil)
		return
	}

	userKey, err := r.Cookie("magma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	s.rdr.RenderXML(w, r, tpl.XmlSessionNew, dtSession{userKey.Value})
}

// nucleusCheckUser is requested by both: game-server and game-client.
//
// See also TestServer_nucleusCheckUser as an example request
// made by game-client.
func (s *Controller) nucleusCheckUser(w http.ResponseWriter, r *http.Request) {
	// userID := chi.URLParam("userID")
	w.WriteHeader(http.StatusOK)
}

type dtHero struct {
	HeroID string
}

func (s *Controller) nucleusEntitlements(w http.ResponseWriter, r *http.Request) {
	s.rdr.RenderXML(w, r, tpl.XmlEntitlements, &dtHero{chi.URLParam(r, "heroID")})
}

func (s *Controller) walletsHandler(w http.ResponseWriter, r *http.Request) {
	// DV = VP,
	// DH = HP,
	// DB = BF,
	s.rdr.RenderXML(w, r, tpl.XmlWallets, nil)
}
