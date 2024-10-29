package models

// TODO
// Create struct for original json card table: FileCard
// For each created table we need a matching struct: Card, Printcard, Langcard, etc
// Make new Folder in dataload project called models - models.go (package models at top) - all structs go in here

type Card struct {
	Name string `json:"name"`
}
