# h_prueba
### Instalacion
```bash
git clone https://github.com/DavidNazareno/HPrueba.git
```
### Uso
<p>Para ejecutar acceder desde la terminal y escribir </p> 

# Nota:
Cuando se hace una mutacion GenerateRequest retonar url, esa url abre un html statico sencillo donde le muestra la informacion de la solicitud y su tecnico.

### Ejecutar

```bash 
go run ./main.go
```

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

### Database Diagram
Digrama de base de datos, con relacions corresondientes para la funcionalidad de la app


### Querys

- getRequest() recibe como parametro el token generado, y retorna el Objecto Request desde la tabla requests en la base de datos.
- getTechOrders() recibe como parametro el id del tecnico y retorna la cantidad de  ordenes que tiene el tecnico en ese momento.

```graphql

 getRequest(token: String!): Request!
 getTechOrders(id: String!): String!
```

### Mutation
```graphql

  AddClient: Client!
  GenerateRequest(ticket: String!): String!

```
- AddClient() recibe como parametro el cliente y es para crear el ticket y almacenarlo en la base de datos 

- GenerateRequest() recibe como parametro el input NewRequest creado en AddCliente para validar el cliente y crear la solicitud junto al token y asingarlo aleatoriamente a un tecnico, retornando el link para el seguimiento.

```graphql

  input NewRequest {
  token: String!
  client: Int!
  status: Int!
  score: Int!
}

```

### Schemas


#### Client 
Objeto client para gestionar los ticket y validacion de cada cliente que haga una solictud al servicio
```graphql

type Client {
  id: ID!
  ticket: String!
}

```
#### Order 
Objeto Order para construir las ordenes y relacionarla a la tabla technician
```graphql

type Order {
  id: ID!
  technician: Technician!

}

```
#### Request 
Objeto Request para construir las solicitudes de los clientes, cuando se valida el ticket, con relacion a a las tablas Client y Status

```graphql
type Request {
  token: String!
  clients: Client!
  status: Int!
  score: Int
}

```
#### Technician 
Objeto Technician para hacer hacer uso de los tecnicos almacenados en la base de datos.
```graphql
type Technician {
  id: ID!
  name: String!
}

```
