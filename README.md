# h_prueba
### Instalacion
```bash
git clone https://github.com/DavidNazareno/HPrueba.git
```
### Uso

<p>Para ejecutar acceder desde la terminal y escribir </p> ```bash go run ./main.go ```

 <hr/>

### Crear Ticket
```graphql
mutation{
  AddClient{
    id
    ticket
  }
}
```
### Generar Solicitud
```graphql
mutation {
  GenerateRequest(ticket: "Agregar Ticket para validar cliente y crear solicitud")
}

```


### Schemas

```graphql

type Client {
  id: ID!
  ticket: String!
}

type Query {
  getRequest(token: String!): Request!
  getTechOrders(token: String!): [Request]
}

type Mutation {
  AddClient: Client!
  addRequest(ticket: NewRequest!): String!
  GenerateRequest(ticket: String!): String!
}

input NewRequest {
  token: String!
  client: Int!
  status: Int!
  score: Int!
}

type Order {
  id: ID!
  technician: Technician!
  token: String!
}

type Request {
  token: String!
  clients: Client!
  status: Int!
  score: Int
}

type Technician {
  id: ID!
  name: String!
}
```
