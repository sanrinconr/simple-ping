## Simple ping
An API that respond health checks

## How run in local
First download the code with:


You must generate a self-signed cert with
```bash
git clone https://github.com/sanrinconr/simple-ping.git
cd simple-ping
```

Generate a self-signed certificate with
```bash
openssl req -x509 -newkey rsa:2048 -keyout server.key -out server.crt -days 10 -nodes -subj "/C=XX/ST=StateName/L=CityName/O=CompanyName/OU=CompanySectionName/CN=CommonNameOrHostname"
```

That generate two files (certificate and private key) that will be used in the next step

Define the next environment variables:
```bash
export CERT_FILE_LOCATION=./server.crt
export KEY_FILE_LOCATION=./server.key
export PORT=1443
```

Finally run the API with

```bash
go run main.go
```

You can check https://127.0.0.1:1443/ping
