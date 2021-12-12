package main

type Database struct {
	Path string
}

// Open creates a new database connection at the path specified.
// If the path does not exist, it will be created.
func Open(path string) (*Database, error) {
	return &Database{Path: path}, nil
}

// Close closes the database connection and cleans up any resources.
// Will block until all pending operations have completed.
func (db *Database) Close() {
}

func (db *Database) Collection() *Collection {
	return &Collection{}
}

// Collection represents a collection of MingoDB documents.
type Collection struct {
	db   *Database
	name string
}

// Name returns the name of the collection.
func (c *Collection) Name() string {
	return c.name
}

// Database returns a pointer to the collection's parent database.
func (c *Collection) Database() *Database {
	return c.db
}

// Drop deletes the collection.
func (c *Collection) Drop() error {
	return nil
}

// InsertOne inserts a single document into the collection.
func (c *Collection) InsertOne(doc interface{}) (*InsertOneResult, error) {
	return nil, nil
}

// InsertMany inserts multiple documents into the collection.
func (c *Collection) InsertMany(docs []interface{}) (*InsertManyResult, error) {
	return nil, nil
}

// Find
func (c *Collection) Find(filter interface{}) (*MultiResult, error) {
	return nil, nil
}

// FindOne
func (c *Collection) FindOne(filter interface{}, result interface{}) (*SingleResult, error) {
	return nil, nil
}

// CountDocuments returns the number of documents that match the filter.
func (c *Collection) CountDocuments(filter interface{}) (int, error) {
	return 0, nil
}

// UpdateOne
func (c *Collection) UpdateOne(doc interface{}) (*UpdateResult, error) {
	return nil, nil
}

// UpdateMany
func (c *Collection) UpdateMany(docs []interface{}) (*UpdateResult, error) {
	return nil, nil
}

// DeleteOne deletes a single document into the collection based on the filter
func (c *Collection) DeleteOne(filter interface{}) (*DeleteResult, error) {
	return nil, nil
}

// DeleteMany inserts multiple documents into the collection.
func (c *Collection) DeleteMany(filter []interface{}) (*DeleteResult, error) {
	return nil, nil
}

// // Aggregate
// func (c *Collection) Aggregate(pipeline interface{}) (*MultiResult, error) {
// 	return nil, nil
// }
