package main

import (
	"github.com/nsan1129/unframed"
)

func prepareStatements() {
	DB.AddStatement("ListPeople",
		unframed.Dbd.Pg,
		`SELECT 
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
	)

	DB.AddStatement("ListPeople",
		unframed.Dbd.Mysql,
		`SELECT
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
	)

	DB.AddStatement("FindPerson",
		unframed.Dbd.Pg,
		`SELECT 
			"Id",
			COALESCE("Name", ''), 
			COALESCE("Password", ''),
			COALESCE("RememberToken", ''), 
			"Admin",
			COALESCE("Code", '0'),
			COALESCE("Contestant", false)
		FROM "People" 
		WHERE LOWER("Name") = LOWER($1) AND "Code" = $2`,
	)

	DB.AddStatement("ShowPerson",
		unframed.Dbd.Pg,
		`SELECT 
			"Id",
			COALESCE("Name", ''), 
			COALESCE("Password", ''),
			COALESCE("RememberToken", ''), 
			"Admin",
			COALESCE("Code", '0'),
			COALESCE("Contestant", false)
		FROM "People" 
		WHERE "Id" = $1`,
	)

	DB.AddStatement("FindPerson",
		unframed.Dbd.Pg,
		`SELECT 
			"Id",
			COALESCE("Name", ''), 
			COALESCE("Password", ''),
			COALESCE("RememberToken", ''), 
			"Admin",
			COALESCE("Code", '0'),
			COALESCE("Contestant", false)
		FROM "People" 
		WHERE LOWER("Name") = LOWER($1) AND "Code" = $2`,
	)

	DB.AddStatement("ListVotes",
		unframed.Dbd.Pg,
		`SELECT 
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
	)

	DB.AddStatement("UpdateVote",
		unframed.Dbd.Pg,
		`UPDATE "Votes"
		SET
			"Votee" = $2
		WHERE
			"Voter" = $1
		AND
			"VoteType" = $3`,
	)

	DB.AddStatement("CreateVote",
		unframed.Dbd.Pg,
		`INSERT INTO "Votes" (
			"Voter",
			"Votee",
			"VoteType"
		) VALUES (
			$1,
			$2,
			$3
		)`,
	)

	DB.AddStatement("DeleteVote",
		unframed.Dbd.Pg,
		`DELETE FROM "Votes"
		WHERE "Voter" = $1
		AND "VoteType" = $2`,
	)

	DB.PrepareStatements()
}
