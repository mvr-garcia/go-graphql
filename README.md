# Go GraphQL API

A GraphQL API built with Go following clean architecture principles, using gqlgen for GraphQL code generation. This project was developed for educational purposes as part of the Full Cycle Go Expert course.

## 🏗️ Architecture

This project follows **Ports and Adapters (Hexagonal Architecture)** pattern, ensuring separation of concerns and high testability:

- **Domain Layer**: Contains business entities and rules
- **Application Layer**: Contains use cases and configuration
- **Infrastructure Layer**: Contains database adapters and external integrations
- **UI Layer**: Contains GraphQL resolvers and schema definitions

## 🛠️ Technologies Used

- **Go**: Programming language
- **gqlgen**: GraphQL code generation library
- **Cobra**: CLI command framework
- **Viper**: Configuration management
- **SQLite**: Database (via modernc.org/sqlite)
- **GraphQL**: Query language for APIs
- **Ports and Adapters**: Architectural pattern

## 📁 Project Structure

```
go-graphql/
├── cmd/                    # CLI commands
│   ├── graphql.go         # GraphQL server command
│   └── root.go            # Root command configuration
├── internal/
│   ├── app/               # Application layer
│   │   └── config.go      # Configuration management
│   ├── domain/            # Domain layer
│   │   ├── entity.go      # Business entities
│   │   ├── errors.go      # Domain errors
│   │   └── port.go        # Interface definitions
│   ├── infra/             # Infrastructure layer
│   │   ├── category_adapter.go  # Category repository implementation
│   │   ├── course_adapter.go    # Course repository implementation
│   │   └── db.go               # Database connection
│   └── ui/                # User interface layer
│       └── graph/         # GraphQL layer
│           ├── generated.go           # Generated GraphQL code
│           ├── resolver.go            # Resolver implementations
│           ├── schema.graphqls        # GraphQL schema
│           ├── schema.resolvers.go    # Custom resolvers
│           └── model/                 # GraphQL models
│               ├── category.go
│               ├── course.go
│               └── models_gen.go      # Generated models
├── main.go                # Application entry point
├── gqlgen.yml            # gqlgen configuration
├── go.mod                # Go modules
└── env.example           # Environment variables example
```

## 🚀 Getting Started

### Prerequisites

- Go 1.24 or higher
- gqlgen CLI tool

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/mvr-garcia/go-graphql.git
   cd go-graphql
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Install gqlgen CLI** (if not already installed)
   ```bash
   go install github.com/99designs/gqlgen@latest
   ```

## 🔧 GraphQL Setup Process

This project was set up using gqlgen with the following steps:

### Step 1: Initialize gqlgen
```bash
gqlgen init
```

### Step 2: Move graph directory to internal/ui/
```bash
mkdir -p internal/ui
mv graph internal/ui/
```

### Step 3: Update schema.graphqls
Edit `internal/ui/graph/schema.graphqls` to define your GraphQL schema:

```graphql
type Category {
  id: ID!
  name: String!
  description: String
  courses: [Course!]!
}

type Course {
  id: ID!
  name: String!
  description: String
  category: Category!
}

input NewCategory {
  name: String!
  description: String
}

input NewCourse {
  name: String!
  description: String
  categoryId: ID!
}

type Query {
  categories: [Category!]!
  courses: [Course!]!
}

type Mutation {
  createCategory(input: NewCategory!): Category!
  createCourse(input: NewCourse!): Course!
}
```

### Step 4: Update gqlgen.yml configuration
Edit `gqlgen.yml` to reflect the new package paths:

```yaml
schema:
  - internal/ui/graph/*.graphqls

exec:
  package: internal/ui/graph
  layout: single-file
  filename: generated.go

model:
  filename: internal/ui/graph/model/models_gen.go
  package: internal/ui/graph/model

resolver:
  package: internal/ui/graph
  layout: follow-schema
  dir: internal/ui/graph
  filename_template: "{name}.resolvers.go"
```

### Step 5: Generate GraphQL code
```bash
gqlgen generate
```

## ▶️ Running the Application

### Start the GraphQL server
```bash
go run main.go graphql-api
```

The server will start on `http://localhost:8080` by default.

### Access GraphQL Playground
Open your browser and navigate to `http://localhost:8080` to access the GraphQL Playground where you can test queries and mutations.

## 📝 Example Queries

### Create a Category
```graphql
mutation {
  createCategory(input: {
    name: "Programming"
    description: "Programming courses"
  }) {
    id
    name
    description
  }
}
```

### Create a Course
```graphql
mutation {
  createCourse(input: {
    name: "Go Fundamentals"
    description: "Learn Go programming language"
    categoryId: "1"
  }) {
    id
    name
    description
    category {
      name
    }
  }
}
```

### Query Categories and Courses
```graphql
query {
  categories {
    id
    name
    description
    courses {
      id
      name
      description
    }
  }
}
```

## 🎯 Features

- **GraphQL API**: Complete GraphQL implementation with queries and mutations
- **Clean Architecture**: Follows Ports and Adapters pattern
- **Code Generation**: Uses gqlgen for automatic GraphQL code generation
- **CLI Interface**: Cobra-based command line interface
- **Configuration Management**: Viper-based configuration
- **Database Integration**: SQLite database with repository pattern
- **Category Management**: CRUD operations for categories
- **Course Management**: CRUD operations for courses with category relationships

## 🔧 Configuration

Create a `.env` file based on `env.example`:

```env
DB_DRIVER=sqlite
DB_DSN=./app.db
PORT=8080
```

## 📄 License

This project is part of an educational challenge from the Go Expert course by Full Cycle.

## 🤝 Contributing

This is an educational project. Feel free to fork and experiment with the code!

## 📚 Learning Resources

- [gqlgen Documentation](https://gqlgen.com/)
- [GraphQL Specification](https://graphql.org/)
- [Go Documentation](https://golang.org/doc/)
- [Cobra CLI](https://cobra.dev/)
- [Viper Configuration](https://github.com/spf13/viper)