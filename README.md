# Ledger

This is a toy ledger application that allow to:
- Create new accounts
- Authorize money
- Present it from the account

## How to run

0. Switch direcory to the project root
1. Create a new data file for Tigerbeetle cluster:
```bash
docker run -v $(pwd)/data:/data ghcr.io/tigerbeetledb/tigerbeetle \
    format --cluster=0 --replica-count=1 --replica=0 /data/0_0.tigerbeetle
```
2. Start the Tigerbeetle locally
```bash
docker compose up
```
3. Install and start Temporalit cluster
```bash
temporalite start --namespace default
```
4. Start the Encore application
```bash
encore run
```
5. Now you can create a new account via GQL query or via internal Encore dashboard

5.1. To create a new account do the following GQL query
   ```gql
   mutation CreateAccount {
        createAccount(input: {id: "2", initialAmount: 12}) {
            status
        }
    }
   ```
   To make it you can install https://www.postman.com/ or any other GQL client. The query URL is `http://localhost:4000/graphql`.

5.2. To use the encore dashboard go to http://localhost:9400/pave-bank-odai/api#ledger.CreateNewAccount and make an API call. Don't use account id 1, it's reserved as a bank account.
Run the application
   
6. Via GQL or via Encore panel now you can Authorize and Present money from a newly created account, check `authorize`, `present` mutations and http://localhost:9400/pave-bank-odai/api#ledger.Authorize, http://localhost:9400/pave-bank-odai/api#ledger.Present.
7. Also you can get account balance via 
```GQL
query Accounts {
    accounts(input: {accountId: "2"}) {
        id
        credits
    }
}
```

or http://localhost:9400/pave-bank-odai/api#ledger.GetAccount

## GraphQL

### To generate new GQL entities use
```bash
 go run github.com/99designs/gqlgen generate
```

### View the GraphQL Playground
Open http://localhost:4000/graphql/playground in your browser.

