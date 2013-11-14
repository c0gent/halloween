package main

import (
	"github.com/nsan1129/auctionLog/log"
	"github.com/nsan1129/unframed"
)

type Vote struct {
	Id        int
	Voter     int
	Votee     int
	VoteType  int
	VoteeName string
	ImageFile string
}

type Person struct {
	Id            int
	Name          string
	Password      string
	RememberToken string
	Admin         bool
	Code          int
	Contestant    bool
	ImageFile     string
	Votes         []*Vote
}

type PeopleMdl struct {
	unframed.DataModel
	People []*Person
	Person *Person
}

func (s *PeopleMdl) List() *PeopleMdl {

	rows, err := DB.Stmts["ListPeople"].Query()
	if err != nil {
		log.Error(err)
	}

	defer rows.Close()

	for rows.Next() {
		sa := new(Person)
		err := rows.Scan(&sa.Id, &sa.Name, &sa.Password, &sa.RememberToken, &sa.Admin, &sa.Code, &sa.Contestant, &sa.ImageFile)
		if err != nil {
			log.Error(err)
		}
		s.People = append(s.People, sa)
	}

	return s
}

/*
DB. stmtText{
		dbdPg: `
		SELECT
			"Id",
			COALESCE("Name", ''),
			COALESCE("Password", ''),
			COALESCE("RememberToken", ''),
			"Admin",
			COALESCE("Code", '0'),
			COALESCE("Contestant", false)
		FROM "People"
		WHERE LOWER("Name") = LOWER($1) AND "Code" = $2`,
	}
*/

func (s *PeopleMdl) find(na string, co int) (bool, *Person) {

	row := DB.Stmts["FindPerson"].QueryRow(na, co)

	pe := new(Person)
	err := row.Scan(&pe.Id, &pe.Name, &pe.Password, &pe.RememberToken, &pe.Admin, &pe.Code, &pe.Contestant)
	if err != nil {
		log.Error(err)
		return false, new(Person)
	}
	return true, pe
}

func (s *PeopleMdl) showPerson(id int) {
	row := DB.Stmts["ShowPerson"].QueryRow(id)

	pe := new(Person)
	err := row.Scan(&pe.Id, &pe.Name, &pe.Password, &pe.RememberToken, &pe.Admin, &pe.Code, &pe.Contestant)
	if err != nil {
		log.Error(err)
	}

	rows, err := DB.Stmts["ListVotes"].Query(pe.Id)
	if err != nil {
		log.Error(err)
	}
	defer rows.Close()

	log.Message("for rows.Next()")
	for rows.Next() {
		sa := new(Vote)
		err := rows.Scan(&sa.Id, &sa.Voter, &sa.Votee, &sa.VoteType, &sa.VoteeName, &sa.ImageFile)
		if err != nil {
			log.Error(err)
		}
		pe.Votes = append(pe.Votes, sa)
	}
	for _, p := range pe.Votes {
		log.Message("Votes:", p.Id, p.Voter, p.Votee, p.VoteType, p.VoteeName, p.ImageFile, "pe.Id=", pe.Id)
	}

	s.Person = pe
}

func (s *PeopleMdl) vote(voter int, votee int, voteType int) {
	DB.Stmts["DeleteVote"].Exec(voter, voteType)
	DB.Stmts["CreateVote"].Exec(voter, votee, voteType)

}

func (s *PeopleMdl) newPerson() *Person {
	ns := new(Person)
	s.People = append(s.People, ns)
	return ns
}

type Session struct {
	Name string
	Code int
}

type SessionsMdl struct {
	*unframed.DataModel
	Sessions []*Session
}

func (s *SessionsMdl) newSession() *Session {
	ns := new(Session)
	s.Sessions = append(s.Sessions, ns)
	return ns
}

/*
func (s *SessionsMdl) commit() {
	log.Message(len(s.Sessions))
	for _, sa := range s.Sessions {
		_, err := DB.OldStatementStore.CreateSession.Exec(sa.Name, sa.Code)
		if err != nil {
			log.Error(err)
		}
	}
}
*/
