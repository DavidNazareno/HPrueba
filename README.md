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


### Schemas


# Client 
Objeto client para gestionar los ticket y validacion de cada cliente que haga una solictud al servicio
```graphql

type Client {
  id: ID!
  ticket: String!
}

```
# Order 
Objeto Order para construir las ordenes y relacionarla a la tabla technician
```graphql

type Order {
  id: ID!
  technician: Technician!

}

```
# Request 
Objeto Order para construir las ordenes y relacionarla a la tabla technician

```graphql
type Request {
  token: String!
  clients: Client!
  status: Int!
  score: Int
}

```
