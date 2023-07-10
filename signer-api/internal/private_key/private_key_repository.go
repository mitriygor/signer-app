package private_key

import (
	"context"
	"database/sql"
	"log"
	"signer-api/config"
	ah "signer-api/pkg/args_helper"
	"strconv"
	"strings"
)

type Repository interface {
	GetAll(args Args) ([]*PrivateKey, error)
}

type PrivateKeyRepository struct {
	DB *sql.DB
}

func NewPrivateKeyRepository(db *sql.DB) Repository {
	return &PrivateKeyRepository{
		DB: db,
	}
}

func (a *PrivateKeyRepository) GetAll(args Args) ([]*PrivateKey, error) {
	sqlQueryBuilder := strings.Builder{}
	ctx, cancel := context.WithTimeout(context.Background(), config.DBTimeout)
	defer cancel()

	sqlQueryBuilder.WriteString(`SELECT id, title, secret FROM private_key WHERE 1 = 1`)

	if args.Title != "" {
		ah.AddArgToQuery(&sqlQueryBuilder, "AND first_name LIKE", args.Title)
	}

	if args.Limit != 0 {
		ah.AddArgToQuery(&sqlQueryBuilder, "LIMIT", strconv.Itoa(args.Limit))
	}

	rows, err := a.DB.QueryContext(ctx, sqlQueryBuilder.String())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	privateKeys := make([]*PrivateKey, 0)

	for rows.Next() {
		privateKey := PrivateKey{}
		err := rows.Scan(&privateKey.ID, &privateKey.Title, &privateKey.Secret)
		if err != nil {
			log.Fatal(err)
		}
		privateKeys = append(privateKeys, &privateKey)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return privateKeys, nil
}
