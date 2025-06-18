package lib

import (
	"context"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type DBArea struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type DBResearchArea struct {
	Id       int      `json:"id" db:"id"`
	Name     string   `json:"name" db:"name"`
	SubAreas []DBArea `json:"subareas" db:"-"`
}

func getParents(conn *pgx.Conn) ([]DBResearchArea, error) {

	// connection will be closed by the caller

	// get distinct parents
	qry := "SELECT researchareas.id, name " +
		"FROM researchareas " +
		"LEFT OUTER JOIN researchareataxonomy " +
		"ON researchareas.id = researchareataxonomy.child_id " +
		"WHERE researchareataxonomy.child_id IS NULL " +
		"ORDER BY id ASC"

	log.Printf("Query Parents: %s", qry)

	rows, err := conn.Query(context.Background(), qry)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	parents, err := pgx.CollectRows(rows, pgx.RowToStructByName[DBResearchArea])

	if err != nil {
		return nil, err
	}

	return parents, nil
}

func getChildren(conn *pgx.Conn, parentid int) ([]DBArea, error) {

	// connection will be closed by caller

	// get the parent's children
	qry := "SELECT researchareataxonomy.child_id as id, name " +
		"FROM researchareataxonomy " +
		"INNER JOIN researchareas " +
		"ON researchareataxonomy.child_id = researchareas.id " +
		"WHERE researchareataxonomy.parent_id = "

	qry += strconv.Itoa(parentid)

	log.Printf("Query Children: %s", qry)

	rows, err := conn.Query(context.Background(), qry)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	children, err := pgx.CollectRows(rows, pgx.RowToStructByName[DBArea])

	if err != nil {
		return nil, err
	}

	return children, nil
}

func GetResearchAreas() ([]DBResearchArea, error) {

	conn, err := GetPostgresConn("ie2")

	if err != nil {
		return nil, err
	}

	defer CloseConn(conn)

	log.Print("Getting ResearchAreas...")
	areas, err := getParents(conn)

	if err != nil {
		return nil, err
	}

	log.Printf("Retrieved %d ResearchAreas", len(areas))

	if len(areas) > 0 {

		// note that range returns a copy to the value element
		// so in order to return the allocated subareas, we need
		// to work with each area directly using the index
		for i := range areas {

			log.Print("Getting ResearchArea children...")
			children, err := getChildren(conn, areas[i].Id)

			if err != nil {
				return areas, err
			}

			log.Printf("Retrieved %d ResearchArea children", len(children))

			// allocate size before copying!
			areas[i].SubAreas = make([]DBArea, len(children))
			copy(areas[i].SubAreas, children)

			log.Printf("SubAreas now contains %d children", len(areas[i].SubAreas))
		}

		cnt := 0
		log.Printf("Returning %d researchareas.", len(areas))

		for _, val := range areas {
			log.Printf("	Returning %d subareas for %s", len(val.SubAreas), val.Name)
			cnt += len(val.SubAreas)
		}

		log.Printf("with %d total subareas", cnt)
	}

	log.Printf("-- AREAS --")
	log.Print(areas)

	return areas, nil
}
