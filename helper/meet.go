package helper

import (
	"math/rand"
	"time"
)

type MeetInterface interface {
	GetMeetLink() []string
}

type Meet struct{}

func NewMeet() MeetInterface {
	return &Meet{}
}

func (m *Meet) GetMeetLink() []string {
	allMeetLinks := []string{
		"https://meet.google.com/uoc-ztpf-tqe",
		"https://meet.google.com/hgn-qqrr-gud",
		"https://meet.google.com/vxi-ytwh-iok",
		"https://meet.google.com/aba-bszh-epg",
		"https://meet.google.com/srg-bbas-quf",
		"https://meet.google.com/cnh-hevt-pfg",
		"https://meet.google.com/ijs-briv-zvj",
		"https://meet.google.com/jco-xazh-ptd",
		"https://meet.google.com/oej-toid-ipd",
		"https://meet.google.com/dju-doak-tdj",
		"https://meet.google.com/zep-heyr-hkw",
		"https://meet.google.com/mpq-smjx-ags",
		"https://meet.google.com/tno-gigx-mbb",
		"https://meet.google.com/evs-ufwy-xvv",
		"https://meet.google.com/mbu-xbcv-ubr",
		"https://meet.google.com/xds-yubt-uzm",
		"https://meet.google.com/raq-nrzg-kky",
		"https://meet.google.com/kha-kugz-byu",
		"https://meet.google.com/nhr-fnbi-aam",
		"https://meet.google.com/eet-urvm-buz",
		"https://meet.google.com/ncg-tnbr-njf",
		"https://meet.google.com/hnp-quhs-osh",
		"https://meet.google.com/niu-pegr-ydp",
		"https://meet.google.com/nhz-fnpy-yga",
		"https://meet.google.com/jrn-johw-qvq",
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	random.Shuffle(len(allMeetLinks), func(i, j int) {
		allMeetLinks[i], allMeetLinks[j] = allMeetLinks[j], allMeetLinks[i]
	})

	return allMeetLinks
}
