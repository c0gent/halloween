package unframed

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
	//"github.com/gorilla/pat"
	//"github.com/gorilla/sessions"
	//"github.com/nsan1129/auctionLog/log"
)

type StmtText map[dbDialogue]string

type OldStatementStore struct {
	ListPeople    *sql.Stmt
	CreatePerson  *sql.Stmt
	CreateSession *sql.Stmt
	FindPerson    *sql.Stmt
	ShowPerson    *sql.Stmt
	ListVotes     *sql.Stmt
	UpdateVote    *sql.Stmt
	CreateVote    *sql.Stmt
	DeleteVote    *sql.Stmt
}

func (s *OldStatementStore) init(db *DB) *OldStatementStore {
	s.makeStmts(db)
	return s
}

func (s *OldStatementStore) makeStmts(db *DB) {

	//  Done
	ListPeopleT := StmtText{
		dbdPg: `
		SELECT 
			"Id",
			COALESCE("Name", ''), 
			COALESCE("Password", ''),
			COALESCE("RememberToken", ''), 
			"Admin",
			COALESCE("Code", '0'),
			COALESCE("Contestant", false),
			COALESCE("ImageFile", '')
		FROM "People"
		WHERE "Contestant" = true
		ORDER BY "Name" ASC`,
	}
	s.storeStmt(ListPeopleT, &s.ListPeople, db)

	/*
		loginT := StmtTexts{
			dbdPg: `SELECT`,
		}
		s.storeStmt(loginT, &s.login)
	*/

	/*
		CreatePersonT := StmtTexts{
			dbdPg: `insert into "contestants"
			("ItemName", "SoldPrice", "Quality", "Qty", "ItemId", "Comment")
			values ($1, $2, $3, $4, $5, $6)`,
		}
		s.storeStmt(CreatePersonT, &s.CreatePerson)
	*/

	FindPersonT := StmtText{
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

	s.storeStmt(FindPersonT, &s.FindPerson, db)

	ShowPersonT := StmtText{
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
		WHERE "Id" = $1`,
	}
	s.storeStmt(ShowPersonT, &s.ShowPerson, db)

	ListVotesT := StmtText{
		dbdPg: `
		SELECT 
			"Votes"."Id" as "Id",
			"Votes"."Voter" as "Voter",
			"Votes"."Votee" as "Votee",
			"Votes"."VoteType" as "VoteType",
			COALESCE("People"."Name", '') as "VoteeName",
			COALESCE("People"."ImageFile", '') as "ImageFile"
		FROM "Votes" 
		INNER JOIN "People"
		ON "Votes"."Votee" = "People"."Id"
		WHERE "Voter" = $1`,
	}
	s.storeStmt(ListVotesT, &s.ListVotes, db)

	UpdateVoteT := StmtText{
		dbdPg: `
		UPDATE "Votes"
		SET
			"Votee" = $2
		WHERE
			"Voter" = $1
		AND
			"VoteType" = $3
		`,
	}
	s.storeStmt(UpdateVoteT, &s.UpdateVote, db)

	CreateVoteT := StmtText{
		dbdPg: `
		INSERT INTO "Votes" (
			"Voter",
			"Votee",
			"VoteType"
		) VALUES (
			$1,
			$2,
			$3
		)
		`,
	}
	s.storeStmt(CreateVoteT, &s.CreateVote, db)

	DeleteVoteT := StmtText{
		dbdPg: `
		DELETE FROM "Votes"
		WHERE "Voter" = $1
		AND "VoteType" = $2
		`,
	}
	s.storeStmt(DeleteVoteT, &s.DeleteVote, db)
}

func (s *OldStatementStore) storeStmt(stt StmtText, storeField **sql.Stmt, db *DB) {
	var err error
	*storeField, err = db.Prepare(stt[db.cdd])

	if err != nil {
		panic(err)
	}

}
