package rank

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"

	"gitlab.com/oiacow/fesl3/backend/network"
)

const (
	k = "k"
	v = "v"
	t = "t"
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

// UpdateStats - UpdateStats about a selected hero
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
	//pointers 
	ans := event.Command.Message
	users, _ := strconv.Atoi(ans["u.[]"])
	sess := r.DB.NewSession()
	////////

	for i := 0; i < users; i++ {
		heroID, _ := ans.IntVal(fmt.Sprintf("u.%d.o", i))
		p, err := r.DB.FindHeroStats(sess, heroID)
		if err != nil {
			logrus.
				WithError(err).
				WithField("heroID", ans[fmt.Sprintf("u.%d.o", i)]).
				Warn("Cant resolve heroStats when updatingStats")
			return
		}

		numKeys, _ := ans.IntVal(fmt.Sprintf("u.%d.s.[]", i))
		for j := 0; j < numKeys; j++ {

			pre := fmt.Sprintf("u.%d.s.%d.", i, j)
			//suggestion create a const dictionary -> c_wallet_hero = HP -> "k" = k -> event.Command.Message = reply
			key := ans.Get(pre + k)
			ut := ans.Get(pre + "ut")
			pt := ans.Get(pre + "pt")
			val := ans.Get(pre + v)
			text := ans.Get(pre + t)

			//ChangeStats in both cases
			if text != "" {
				// c_items, c_eqp..
				val = text
				logrus.Println("--UpdateStat replace ut 0--"+ key, val, ut)
				r.changeStats(&p, key, val, ut, pt)
				//GOTO LN102
			} else{
				logrus.Println("--UpdateStat sum ut 3"+ key, val, ut)
				r.changeStats(&p, key, val, ut, pt)
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
