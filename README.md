# grpcApi- Unary API
<p>1 request-1 response</p>
<h2>Build a simple API grpc: AddUser</h2>
<h2>Requirement: </h2>
    <p>- docker</p>
    <p>- golang (1.18.3)</p>
<h2>I.Create a Mysql database </h2>
 <p>- cd grpcApi/server/sql</p>
 <p>- docker-compose up -d </p>
<h2>II. Run Server </h2>
 <p>- cd grpcApi/server</p>
 <p>- go run main.go </p>
 <h2>II. Run Client (AddUser request) </h2>
 <p>- cd grpcApi/client</p>
 <p>- go run client.go </p>
