# Ledger

This Encore example shows how you can build a GraphQL server with Encore.

## Run the application

Start the database
```bash
docker compose up
```

Run the application

```bash
$ encore run
```

## GraphQL

To generate new GQL entities use
```bash
 go run github.com/99designs/gqlgen generate
```

## View the GraphQL Playground
Open http://localhost:4000/graphql/playground in your browser.

## Using the API

Execute the below queries using the GraphQL Playground (or method of your choice).

#### Shorten a URL

```graphql
mutation {
  shorten(url: "https://encore.dev") {
    id
    url
  }
}
```

#### Listing all shortened URLs

```graphql
query {
  urls {
    id
    url
  }
}
```

#### Getting a URL from a shortened ID

```graphql
query {
  get(id: "some-id") {  # Use an actual ID you have
    id
    url
  }
}
```
