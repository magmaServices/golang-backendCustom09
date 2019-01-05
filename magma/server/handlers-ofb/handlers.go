package ofb

import (
	"net/http"

	"gitlab.com/oiacow/fesl3/magma/tpl"
)

func (c *Controller) ofbProducts(w http.ResponseWriter, r *http.Request) {
	c.rdr.RenderXML(w, r, tpl.XmlProducts, nil)
}
