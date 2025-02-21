package stack

import (
	"backend/ent/gen"
	"backend/ent/gen/liveshow"
	"backend/pkg/datastore"
	"context"
	"log"
	"os"
	"time"
)

var CurrentMedia *gen.LiveShow

func StartStack(db *datastore.EntDB) {
	for {
		_liveShow, err := db.LiveShow.Query().Where(liveshow.Viewed(false), liveshow.CurrentPlayed(false)).WithUpload().First(context.TODO())
		if err != nil || _liveShow == nil {
			CurrentMedia = nil
			time.Sleep(2 * time.Second)
			continue
		}

		now := time.Now()
		err = _liveShow.Update().SetStartedTime(now).SetEndedTime(now.Add(time.Second * time.Duration(_liveShow.Duration))).SetCurrentPlayed(true).Exec(context.TODO())
		CurrentMedia = _liveShow
		if err != nil {
			log.Println(err)
			continue
		}
		time.Sleep(time.Duration(_liveShow.Duration) * time.Second)
		err = _liveShow.Update().SetViewed(true).SetCurrentPlayed(false).Exec(context.TODO())
		if err != nil {
			log.Println(err)
			continue
		}

		err = os.Remove(os.Getenv("UPLOAD_DIRECTORY") + "/" + _liveShow.Edges.Upload.FilePath)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
