package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/domain"
	"github.com/masrayfa/internals/models/web"
)

type NodeRepositoryImpl struct {
}

func NewNodeRepository() NodeRepository {
	return &NodeRepositoryImpl{}
}

func (n *NodeRepositoryImpl) FindAll(ctx context.Context, pool *pgxpool.Pool, currentUser *web.UserRead) ([]domain.Node, error) {
	tx, err := pool.Begin(ctx)
	helper.PanicIfError(err)

	var nodes []domain.Node

	var script string
	var rows pgx.Rows

	// if user is admin, show all nodes
	if currentUser.IsAdmin {
		script = "SELECT * FROM node"
		rows, err := tx.Query(ctx, script)
		helper.PanicIfError(err)
		defer rows.Close()
	} else { // if user is not admin, show only nodes that belong to user
		script = "SELECT * FROM node WHERE id_user = $1"

		rows, err := tx.Query(ctx, script, currentUser.IdUser)
		helper.PanicIfError(err)
		defer rows.Close()
	}


	for rows.Next() {
		var node domain.Node
		err := rows.Scan(&node.IdNode, &node.Name, &node.Location, &node.IdUser, &node.IdHardwareNode)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	// if err := rows.Err(); err != nil {
	// 	return nil, err
	// }

	return nodes, nil
}

func (n *NodeRepositoryImpl) FindById(ctx context.Context, pool *pgxpool.Pool, id int64) (domain.Node, error) {
	tx, err := pool.Begin(ctx)
	helper.PanicIfError(err)

	var node domain.Node

	script := "SELECT * FROM node WHERE id = $1"

	err = tx.QueryRow(ctx, script, id).Scan(&node.IdNode, &node.Name, &node.Location, &node.IdUser, &node.IdHardwareNode)
	if err != nil {
		return node, err
	}

	return node, nil
}

func (n *NodeRepositoryImpl) GetHardwareNode(ctx context.Context, pool *pgxpool.Pool, hardwareId int64) ([]domain.Node, error) {
	tx, err := pool.Begin(ctx)
	helper.PanicIfError(err)

	var nodes []domain.Node

	script := "SELECT * FROM node WHERE id_hardware = $1"

	rows, err := tx.Query(ctx, script, hardwareId)
	if err != nil {
		return nodes, err
	}
	defer rows.Close()

	for rows.Next() {
		var node domain.Node
		err := rows.Scan(&node.IdNode, &node.Name, &node.Location, &node.IdUser, &node.IdHardwareNode)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nodes, nil
}

func (n *NodeRepositoryImpl) Create(ctx context.Context, pool *pgxpool.Pool, nodePayload domain.Node, currentUser *web.UserRead) (domain.Node, error) {
	tx, err := pool.Begin(ctx)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(ctx, tx)

	// node object
	node := domain.Node {
		Name: nodePayload.Name,
		Location: nodePayload.Location,
		IdUser: currentUser.IdUser,
		IdHardwareNode: nodePayload.IdHardwareNode,
	}

	// sql script
	script := "INSERT INTO node (name, location, id_user, id_hardware) VALUES ($1, $2, $3, $4)"

	// insert node
	_, err = tx.Exec(ctx, script, node.Name, node.Location, node.IdUser, node.IdHardwareNode)
	if err != nil {
		return node, err
	}

	return node, nil
}

func (n *NodeRepositoryImpl) Update(ctx context.Context, pool *pgxpool.Pool, node domain.Node) (domain.Node, error) {
	tx, err := pool.Begin(ctx)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(ctx, tx)

	script := "UPDATE node SET name = $1, location = $2 WHERE id = $3"

	res, err := tx.Exec(ctx, script, node.Name, node.Location, node.IdNode)
	if err != nil {
		return node, err
	}

	if res.RowsAffected() != 1 {
		log.Println("No row affected on update node with id: ", node.IdNode)
		return node, err
	}

	return node, nil
}

func (n *NodeRepositoryImpl) Delete(ctx context.Context, pool *pgxpool.Pool, id int64) error {
	tx, err := pool.Begin(ctx)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(ctx, tx)

	script := "DELETE FROM nodes WHERE id = $1"

	_, err = tx.Exec(ctx, script, id)
	if err != nil {
		return err
	}
	return nil
}