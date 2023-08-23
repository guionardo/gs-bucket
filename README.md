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

<h1 id="gs-bucket-api">GS-Bucket API v0.0.5</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

This application will run a HTTP server to store files

Email: <a href="mailto:guionardo@gmail.com">Guionardo Furlan</a> Web: <a href="https://github.com/guionardo/gs-bucket">Guionardo Furlan</a>

<h1 id="gs-bucket-api-auth">auth</h1>

## get__auth_

`GET /auth/`

*List all users allowed to publish*

<h3 id="get__auth_-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|api-key|header|string|true|API Key (master key)|

> Example responses

> 200 Response

```json
[
  {
    "key": "string",
    "user": "string"
  }
]
```

<h3 id="get__auth_-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|[server.ErrResponse](#schemaserver.errresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[server.ErrResponse](#schemaserver.errresponse)|

<h3 id="get__auth_-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[domain.AuthResponse](#schemadomain.authresponse)]|false|none|none|
|» key|string|false|none|none|
|» user|string|false|none|none|

<aside class="success">
This operation does not require authentication
</aside>

## post__auth_{user}

`POST /auth/{user}`

*Create a key for a user*

Post a file to a pad, accessible for anyone

<h3 id="post__auth_{user}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|api-key|header|string|true|API Key|
|user|path|string|true|User name|

> Example responses

> 201 Response

```json
{
  "key": "string",
  "user": "string"
}
```

<h3 id="post__auth_{user}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Created|[domain.AuthResponse](#schemadomain.authresponse)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Required user name|[server.ErrResponse](#schemaserver.errresponse)|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|[server.ErrResponse](#schemaserver.errresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[server.ErrResponse](#schemaserver.errresponse)|

<aside class="success">
This operation does not require authentication
</aside>

## delete__auth_{user}

`DELETE /auth/{user}`

*Delete all keys of user*

<h3 id="delete__auth_{user}-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|api-key|header|string|true|API Key|
|user|path|string|true|User name|

> Example responses

> 201 Response

```json
{
  "code": 0,
  "error": "string",
  "status": "string"
}
```

<h3 id="delete__auth_{user}-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Created|[server.ErrResponse](#schemaserver.errresponse)|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|Unauthorized|[server.ErrResponse](#schemaserver.errresponse)|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[server.ErrResponse](#schemaserver.errresponse)|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="gs-bucket-api-pads">pads</h1>

## get__pads

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
    "owner": "string",
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
|» owner|string|false|none|none|
|» seen_count|integer|false|none|none|
|» size|integer|false|none|none|
|» slug|string|false|none|none|
|» valid_until|string|false|none|none|

<aside class="success">
This operation does not require authentication
</aside>

## post__pads

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
|api-key|header|string|true|API Key|
|name|query|string|true|File name|
|slug|query|string|false|Slug or easy name (if not informed, will be used a hash value)|
|ttl|query|string|false|Time to live (i.Ex 300s, 1.5h or 2h45m). Valid time units are: 's', 'm', 'h')|
|delete-after-read|query|boolean|false|If informed, the file will be deleted after first download.|
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
  "owner": "string",
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
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|Internal Server Error|[server.ErrResponse](#schemaserver.errresponse)|
|507|[Insufficient Storage](https://tools.ietf.org/html/rfc2518#section-10.6)|Insufficient Storage|[server.ErrResponse](#schemaserver.errresponse)|

<aside class="success">
This operation does not require authentication
</aside>

## get__pads_{code}

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

<h2 id="tocS_domain.AuthResponse">domain.AuthResponse</h2>
<!-- backwards compatibility -->
<a id="schemadomain.authresponse"></a>
<a id="schema_domain.AuthResponse"></a>
<a id="tocSdomain.authresponse"></a>
<a id="tocsdomain.authresponse"></a>

```json
{
  "key": "string",
  "user": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|key|string|false|none|none|
|user|string|false|none|none|

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
  "owner": "string",
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
|owner|string|false|none|none|
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


_swagger data generated @ Tue Aug 22 2023 21:10:34 GMT+0000 (Coordinated Universal Time)_
