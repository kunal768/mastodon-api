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
  ALLOWED_ORIGIN=your_client_url_with_port
```

7. To get value for `USER_ID` use below curl with your <b> Mastodon Social Acccount name </b> and copy the `id` value from `JSON` response
```zsh
curl --location 'https://mastodon.social/api/v1/accounts/lookup?acct=<accountname>'
```
8. Once the .env file is set up in the `./cmd/master` run the below command opening shell in the same directory :
```bash
go run main.go
```

   

>[!IMPORTANT]  
>to avoid CORS error update `ALLOWED_ORIGIN` variable in server .env with IP from where below client will be hosted


<br />
<br />

## Client Setup

To set up and run the Next.js client:

1. Open another shell and run below command : 
```bash
npm install && npm run dev
```

<br />
<br />

Congratulations you're good to go ! Made with ❤️  
