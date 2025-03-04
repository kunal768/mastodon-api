CMPE272 : Mastodon API Exercise

## Running the Server locally 

```bash
git clone https://github.com/kunal768/mastodon-api.git
```

The Go server file is located at `cmd/master/main.go`. To run the server:

1. Ensure you have Go installed on your system.
2. Navigate to the `cmd/master` directory.
4. Set up your mastodon social account and in developer settings create a project.
5. Give read, write and push permissions to the project
6. Create a .env in `cmd/master` file with the following Mastodon account values:
  ```env
  MASTODON_SERVER=your_server_url
  CLIENT_KEY=your_client_key
  CLIENT_SECRET=your_client_secret
  ACCESS_TOKEN=your_access_token
  USER_ID=your_user_id
```

7. To get value for `USER_ID` use below curl with your <b> Mastodon Social Acccount name </b> and copy the `id` value from `JSON` response
```zsh
curl --location 'https://mastodon.social/api/v1/accounts/lookup?acct=<accountname>'
```
8. Once the .env file is set up in the `./cmd/master` Run the command: `go run main.go`

## Client Setup

To set up and run the Next.js client:

1. Open another shell and run below command:

   ```zsh
   npm install && npm run dev
   ```

Congratulations you're good to go !

Made with ❤️  
