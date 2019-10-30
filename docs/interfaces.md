# Interfaces

Interfaces serve as parent objects from which other objects can inherit.  
An interface has its own list of named fields that are shared by implementing objects.

## Article

Represents an arXiv article

### Fields

#### id ([String!](scalars.md#string))

The unique id of the arXiv article

#### title ([String!](scalars.md#string))

The title of the arXiv article

#### summary ([String!](scalars.md#string))

The summary of the arXiv article

#### published ([DateTime!](scalars.md#string))

The date and time when arXiv article was published

#### updated ([DateTime!](scalars.md#string))

The date and time when arXiv article was last updated

#### primaryCategory ([Category!](interfaces.md#Category))

The primary known category for the arXiv article

#### category ([[Category!]!](interfaces.md#Category))

The list of categories to which arXiv article belongs

#### author ([[Author!]!](interfaces.md#Author))

The list of authors of the arXiv article

#### link ([[Link!]!](interfaces.md#Link))

The list of types of links for the arXiv article

## Category

Represents the category of arXiv article

### Fields

#### scheme ([String!](scalars.md#string))

The link to the arXiv atom scheme

#### term ([String!](scalars.md#string))

The term represents the field and the category of the arXiv

## Author

Represents an author of arXiv article

### Fields

#### name ([String!](scalars.md#string))

The name of the author of arXiv article

## Link

Represents the link related information of arXiv article

### Fields

#### href ([String!](scalars.md#string))

The hyperlink to document of arXiv article

#### rel ([String!](scalars.md#string))

The relation string of the link

#### title ([String!](scalars.md#string))

The type of the document available via link
