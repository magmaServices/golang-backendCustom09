package ofb

import (
	"net/http"

	"localhost/go-heroes/fesl-backend/magma/tpl"
)

func (c *Controller) ofbProducts(w http.ResponseWriter, r *http.Request) {
	c.rdr.RenderXML(w, r, tpl.XmlProducts, nil)
}
