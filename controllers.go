package main

import (
	"github.com/nsan1129/auctionLog/log"
	"github.com/nsan1129/unframed"
	"net/http"
)

func listPeople(w http.ResponseWriter, r *http.Request) {

	pm := new(PeopleMdl).List()
	pm.SetSession(r)

	err := TS.listPeople.ExecuteTemplate(w, "base", pm)
	if err != nil {
		log.Error(err)
	}

}

func composeSession(w http.ResponseWriter, r *http.Request) {
	err := TS.composeSession.ExecuteTemplate(w, "base", new(unframed.DataModel).SetLoginFailure(r, false))
	if err != nil {
		log.Error(err)
	}
	log.Message("---- composeSession() ---- ")
}

func failCreateSession(w http.ResponseWriter, r *http.Request) {

	err := TS.failCreateSession.ExecuteTemplate(w, "base", new(unframed.DataModel).SetLoginFailure(r, true))
	if err != nil {
		log.Error(err)
	}
	log.Message("---- failCreateSession() complete ----")
}

func createSession(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error(err)
	}
	sm := new(PeopleMdl)
	newses := sm.newPerson()
	err = DB.Decoder.Decode(newses, r.PostForm)
	if err != nil {
		log.Error(err)
	}

	valid, pe := sm.find(newses.Name, newses.Code)

	if valid {
		sm.SetSession(r)

		sm.SetSessionValues(w, r, pe)

		http.Redirect(w, r, "/Person/show", http.StatusFound)
	} else {
		http.Redirect(w, r, "/Session/create/fail", http.StatusFound)
	}

	log.Message("---- createSession() complete ----  Name:", newses.Name, "Code:", newses.Code, "Found?:", valid)

}

func deleteSession(w http.ResponseWriter, r *http.Request) {
	dm := new(unframed.DataModel).InitSession(r)
	dm.DeleteSession(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func showPerson(w http.ResponseWriter, r *http.Request) {
	sm := new(PeopleMdl)
	sm.SetSession(r)
	if sm.IsLoggedIn() {
		sm.showPerson(sm.GetSessionValues().(*Person).Id)
		err := TS.showPerson.ExecuteTemplate(w, "base", sm)
		if err != nil {
			log.Error(err)
		}
	} else {
		http.Redirect(w, r, "/Session/create", http.StatusFound)
	}

	log.Message("---- showPerson() complete ----")
}

func voteUpdatePerson(w http.ResponseWriter, r *http.Request) {

	sm := new(PeopleMdl)
	sm.SetSession(r)

	err := r.ParseForm()
	if err != nil {
		log.Error(err)
	}

	vc := new(Vote)
	err = DB.Decoder.Decode(vc, r.PostForm)
	log.Message("vc.VoteType:", vc.VoteType)
	log.Message("vc.Votee:", vc.Votee)
	log.Message("sm.GetSessionValues(); Voter:", sm.GetSessionValues().(*Person).Id)
	sm.vote(sm.GetSessionValues().(*Person).Id, vc.Votee, vc.VoteType)
	if err != nil {
		log.Error(err)
	}
	http.Redirect(w, r, "/Person/show", http.StatusFound)

	//getvotes
	//check for existing vote of that type
	//forward change to either update or insert

}

/*

func composePerson(w http.ResponseWriter, r *http.Request) {
	err := TS.composePerson.ExecuteTemplate(w, "base", "None Yet")
	if err != nil {
		log.Error(err)
	}
	log.Message("---- composePerson() complete  ----")
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error(err)
	}
	sm := new(PeopleMdl)

	err = DB.decoder.Decode(sm.newPerson(), r.PostForm)
	if err != nil {
		log.Error(err)
	}
	sm.commit()

	http.Redirect(w, r, "/People/list", http.StatusFound)

	log.Message("---- createPerson() complete ----")
}
*/
