package test

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
	"point-of-sale.go/v1/internal/projectpath"
	"point-of-sale.go/v1/internal/web"
)

func SetupTest() *web.App {
	app := web.NewApp()
	ctx := context.Background()
	err := app.DB().BeginFunc(ctx, func(tx pgx.Tx) error {
		// https://stackoverflow.com/questions/3327312/how-can-i-drop-all-the-tables-in-a-postgresql-database
		if _, err := tx.Exec(ctx, "DROP SCHEMA public CASCADE"); err != nil {
			return err
		}
		if _, err := tx.Exec(ctx, "CREATE SCHEMA public"); err != nil {
			return err
		}
		if _, err := tx.Exec(ctx, "GRANT ALL ON SCHEMA public TO postgres"); err != nil {
			return err
		}
		if _, err := tx.Exec(ctx, "GRANT ALL ON SCHEMA public TO public"); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	m, err := migrate.New(fmt.Sprintf("file://%s/internal/db/migrations", projectpath.Root), app.Config().DATABASE_URL)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil {
		panic(err)
	}

	return app

}
