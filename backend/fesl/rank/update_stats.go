package rank

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"

	"gitlab.com/oiacow/fesl3/backend/network"
)

type ansUpdateStats struct {
	Txn   string      `fesl:"TXN"`
	Users []userStats `fesl:"u"`
}

type userStats struct {
	OwnerID   int          `fesl:"o"`  // 3
	OwnerType int          `fesl:"ot"` // 1
	Stats     []updateStat `fesl:"s"`
}

type updateStat struct {
	Key        string `fesl:"k"`  // c_ltp
	PointType  int    `fesl:"pt"` // 0
	Text       string `fesl:"t"`  // ""
	UpdateType int    `fesl:"ut"` // 0
	Value      string `fesl:"v"`  // 9025.0000
}

type stat struct {
	text  string
	value float64
}

// UpdateStats - updates stats about a soldier
func (r *Ranking) UpdateStats(event network.EventClientCommand) {
	switch event.Client.GetClientType() {
	case "server":
		r.serverUpdateStats(&event)
	default:
		r.clientUpdateStats(&event)
	}
}

func (r *Ranking) clientUpdateStats(event *network.EventClientCommand) {
	r.updateStats(event)
}

func (r *Ranking) serverUpdateStats(event *network.EventClientCommand) {
	r.updateStats(event)
}

func (r *Ranking) updateStats(event *network.EventClientCommand) {
	reply := event.Command.Message
	users, _ := strconv.Atoi(event.Command.Message["u.[]"])
	sess := r.DB.NewSession()

	for i := 0; i < users; i++ {
		heroID, _ := reply.IntVal(fmt.Sprintf("u.%d.o", i))
		p, err := r.DB.FindHeroStats(sess, heroID)
		if err != nil {
			logrus.
				WithError(err).
				WithField("heroID", reply[fmt.Sprintf("u.%d.o", i)]).
				Warn("Cant find heroStats when updating stats")
			return
		}

		numKeys, _ := reply.IntVal(fmt.Sprintf("u.%d.s.[]", i))
		for j := 0; j < numKeys; j++ {
			prefix := fmt.Sprintf("u.%d.s.%d.", i, j)

			key :=  reply.Get(prefix + "k")
			ut :=   reply.Get(prefix + "ut")
			pt :=   reply.Get(prefix + "pt")
			val :=  reply.Get(prefix + "v")
			txt :=  reply.Get(prefix + "t")
			if txt != "" {
				// c_items, c_eqp..
				val = txt
			}

			err := r.changeStats(&p, key, val, ut, pt)
			if err != nil {
				// TODO: adders for stats
				logrus.
					WithError(err).
					WithFields(logrus.Fields{
						"key":        key,
						"updatetype": ut,
						"pointtype":  pt,
						"value":      val,
						"txt":       txt,
					}).
					Warn("rank.UpdateStats, query update ignored!")
			}
		}

		if err = r.commitStats(sess, &p, heroID); err != nil {
			logrus.Error(err)
		}
	}

	r.answer(event.Client, event.Command.PayloadID, ansUpdateStats{
		Txn: rankUpdateStats,
	})
}
