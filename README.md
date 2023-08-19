# gs-bucket

[![github](https://img.shields.io/badge/github-gs--bucket-blue)](https://github.com/guionardo/gs-bucket)
![go-version](https://img.shields.io/github/go-mod/go-version/guionardo/gs-bucket)
[![SonarCloud analysis](https://github.com/guionardo/gs-bucket/actions/workflows/sonarcloud.yml/badge.svg)](https://github.com/guionardo/gs-bucket/actions/workflows/sonarcloud.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=guionardo_gs-bucket&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=guionardo_gs-bucket)

[![deploy](https://img.shields.io/badge/Deploy-fly.io-orange)](https://gs-bucket.fly.dev/)

![this-version](https://img.shields.io/badge/This%20Version-$THISVERSION$-blue)
![release](https://img.shields.io/github/v/release/guionardo/gs-bucket?display_name=tag&style=flat-square)

Temporary web file transfer

This application will run a HTTP server to store files

## Swagger documentation


<!-- Generator: Widdershins v4.0.1 -->

<h1 id="gs-bucket-api">GS-Bucket API v0.4</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

This application will run a HTTP server to store files

Email: <a href="mailto:guionardo@gmail.com">Guionardo Furlan</a> Web: <a href="https://github.com/guionardo/gs-bucket">Guionardo Furlan</a>

<h1 id="gs-bucket-api-pads">pads</h1>

## get__pads

> Code samples

```shell
# You can also use wget
curl -X GET /pads \
  -H 'Accept: application/json'

```

```http
GET /pads HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('/pads',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json'
}

result = RestClient.get '/pads',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/pads', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/pads', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/pads");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/pads", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /pads`

*List pads*

> Example responses

> 200 Response

```json
[
  {
    "content_type": "string",
    "creation_date": "string",
    "delete_after_read": true,
    "last_seen": "string",
    "name": "string",
    "seen_count": 0,
    "size": 0,
    "slug": "string",
    "valid_until": "string"
  }
]
```

<h3 id="get__pads-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|[server.ErrResponse](#schemaserver.errresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[server.ErrResponse](#schemaserver.errresponse)|

<h3 id="get__pads-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[domain.File](#schemadomain.file)]|false|none|none|
|» content_type|string|false|none|none|
|» creation_date|string|false|none|none|
|» delete_after_read|boolean|false|none|none|
|» last_seen|string|false|none|none|
|» name|string|false|none|none|
|» seen_count|integer|false|none|none|
|» size|integer|false|none|none|
|» slug|string|false|none|none|
|» valid_until|string|false|none|none|

<aside class="success">
This operation does not require authentication
</aside>

## post__pads

> Code samples

```shell
# You can also use wget
curl -X POST /pads?name=string \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST /pads?name=string HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = 'string';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('/pads?name=string',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.post '/pads',
  params: {
  'name' => 'string'
}, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('/pads', params={
  'name': 'string'
}, headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/pads', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/pads?name=string");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/pads", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /pads`

*Create a pad*

Post a file to a pad, accessible for anyone

> Body parameter

```json
"string"
```

<h3 id="post__pads-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|name|query|string|true|File name|
|ttl|query|string|false|Time to live|
|delete-after-read|query|boolean|false|If informed, the file will be deleted after first download|
|body|body|string|true|Content|

> Example responses

> 201 Response

```json
{
  "content_type": "string",
  "creation_date": "string",
  "delete_after_read": true,
  "last_seen": "string",
  "name": "string",
  "seen_count": 0,
  "size": 0,
  "slug": "string",
  "valid_until": "string"
}
```

<h3 id="post__pads-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Created|[domain.File](#schemadomain.file)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[server.ErrResponse](#schemaserver.errresponse)|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|[server.ErrResponse](#schemaserver.errresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[server.ErrResponse](#schemaserver.errresponse)|

<aside class="success">
This operation does not require authentication
</aside>

## get__pads_{code}

> Code samples

```shell
# You can also use wget
curl -X GET /pads/{code} \
  -H 'Accept: application/json'

```

```http
GET /pads/{code} HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('/pads/{code}',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json'
}

result = RestClient.get '/pads/{code}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('/pads/{code}', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/pads/{code}', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/pads/{code}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/pads/{code}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /pads/{code}`

*Download a pad*

<h3 id="get__pads_{code}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|code|path|string|true|File code|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="get__pads_{code}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|string|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|[server.ErrResponse](#schemaserver.errresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[server.ErrResponse](#schemaserver.errresponse)|

<aside class="success">
This operation does not require authentication
</aside>

## delete__pads_{code}

> Code samples

```shell
# You can also use wget
curl -X DELETE /pads/{code} \
  -H 'Accept: application/json'

```

```http
DELETE /pads/{code} HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('/pads/{code}',
{
  method: 'DELETE',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json'
}

result = RestClient.delete '/pads/{code}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.delete('/pads/{code}', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('DELETE','/pads/{code}', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/pads/{code}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("DELETE");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("DELETE", "/pads/{code}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`DELETE /pads/{code}`

*Delete a pad*

<h3 id="delete__pads_{code}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|code|path|string|true|File code|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="delete__pads_{code}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|string|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|Not Found|[server.ErrResponse](#schemaserver.errresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[server.ErrResponse](#schemaserver.errresponse)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_domain.File">domain.File</h2>
<!-- backwards compatibility -->
<a id="schemadomain.file"></a>
<a id="schema_domain.File"></a>
<a id="tocSdomain.file"></a>
<a id="tocsdomain.file"></a>

```json
{
  "content_type": "string",
  "creation_date": "string",
  "delete_after_read": true,
  "last_seen": "string",
  "name": "string",
  "seen_count": 0,
  "size": 0,
  "slug": "string",
  "valid_until": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|content_type|string|false|none|none|
|creation_date|string|false|none|none|
|delete_after_read|boolean|false|none|none|
|last_seen|string|false|none|none|
|name|string|false|none|none|
|seen_count|integer|false|none|none|
|size|integer|false|none|none|
|slug|string|false|none|none|
|valid_until|string|false|none|none|

<h2 id="tocS_server.ErrResponse">server.ErrResponse</h2>
<!-- backwards compatibility -->
<a id="schemaserver.errresponse"></a>
<a id="schema_server.ErrResponse"></a>
<a id="tocSserver.errresponse"></a>
<a id="tocsserver.errresponse"></a>

```json
{
  "code": 0,
  "error": "string",
  "status": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|code|integer|false|none|application-specific error code|
|error|string|false|none|application-level error message, for debugging|
|status|string|false|none|user-level status message|


_swagger data generated @ Sat Aug 19 2023 01:19:43 GMT+0000 (Coordinated Universal Time)_
