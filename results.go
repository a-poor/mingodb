package main

type InsertOneResult struct {
}

type InsertManyResult struct {
}

type SingleResult struct {
}

type MultiResult struct {
}

type UpdateResult struct {
	UpdateCount int // Number of rows updated
}

type DeleteResult struct {
	DeleteCount int // Number of rows deleted
}
