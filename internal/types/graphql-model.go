package types

// FieldScalar represents the scalar types used in GraphQL.
type FieldScalar string

const (
	ID      FieldScalar = "ID"
	Int     FieldScalar = "Int"
	BigInt  FieldScalar = "BigInt"
	String  FieldScalar = "String"
	Date    FieldScalar = "Date"
	Boolean FieldScalar = "Boolean"
	Bytes   FieldScalar = "Bytes"
	Float   FieldScalar = "Float"
)

// GraphQLJsonObjectType represents a JSON object type in GraphQL.
type GraphQLJsonObjectType struct {
	Name   string                 `json:"name"`
	Fields []GraphQLJsonFieldType `json:"fields"`
}

// GraphQLJsonFieldType represents a field within a JSON object type.
type GraphQLJsonFieldType struct {
	Name          string                 `json:"name"`
	Type          FieldScalar            `json:"type"` // Excluding 'ID' type is handled in usage
	JsonInterface *GraphQLJsonObjectType `json:"jsonInterface,omitempty"`
	Nullable      bool                   `json:"nullable"`
	IsArray       bool                   `json:"isArray"`
}

// GraphQLModelsRelationsEnums represents the structure of models, relations, and enums.
type GraphQLModelsRelationsEnums struct {
	Models    []GraphQLModelsType    `json:"models"`
	Relations []GraphQLRelationsType `json:"relations"`
	Enums     []GraphQLEnumsType     `json:"enums"`
}

// GraphQLEnumsType represents an enum type in GraphQL.
type GraphQLEnumsType struct {
	Name        string   `json:"name"`
	Values      []string `json:"values"`
	Description *string  `json:"description,omitempty"`
}

// GraphQLModelsType represents the structure of a GraphQL model.
type GraphQLModelsType struct {
	Name        string               `json:"name"`
	Fields      []GraphQLEntityField `json:"fields"`
	Indexes     []GraphQLEntityIndex `json:"indexes"`
	FullText    *GraphQLFullTextType `json:"fullText,omitempty"`
	Description *string              `json:"description,omitempty"`
}

// GraphQLEntityField represents a field in a GraphQL entity.
type GraphQLEntityField struct {
	Name          string                 `json:"name"`
	Type          string                 `json:"type"`
	JsonInterface *GraphQLJsonObjectType `json:"jsonInterface,omitempty"`
	IsArray       bool                   `json:"isArray"`
	Nullable      bool                   `json:"nullable"`
	IsEnum        bool                   `json:"isEnum"`
	Description   *string                `json:"description,omitempty"`
}

// IndexType represents different types of indexes.
type IndexType string

const (
	BTREE  IndexType = "btree"
	HASH   IndexType = "hash"
	GIST   IndexType = "gist"
	SPGIST IndexType = "spgist"
	GIN    IndexType = "gin"
	BRIN   IndexType = "brin"
)

// GraphQLEntityIndex represents an index on a GraphQL entity.
type GraphQLEntityIndex struct {
	Fields []string   `json:"fields"`
	Unique *bool      `json:"unique,omitempty"`
	Using  *IndexType `json:"using,omitempty"`
}

// GraphQLFullTextType represents full-text search fields in GraphQL.
type GraphQLFullTextType struct {
	Fields   []string `json:"fields"`
	Language *string  `json:"language,omitempty"`
}

// GraphQLRelationsType represents a relationship between models.
type GraphQLRelationsType struct {
	From       string  `json:"from"`
	Type       string  `json:"type"` // 'hasOne', 'hasMany', 'belongsTo'
	To         string  `json:"to"`
	ForeignKey string  `json:"foreignKey"`
	FieldName  *string `json:"fieldName,omitempty"`
}
