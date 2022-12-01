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


## Endpoints

### [Home (this page)](/)

### [Stored files](/store)

<pre><code>
[
  {
    "name": "teste3.json",
    "created": "2022-11-30T17:30:27-03:00",
    "valid_until": "2022-12-01T17:30:27-03:00",
    "code": "63890ee3_dGVzdGUzLmpzb24=",
    "mime_type": "application/json",
    "size": 36,
    "url": "http://gs-bucket.fly.dev/63890ee3_dGVzdGUzLmpzb24="
  }
]
</code></pre>

### Sending a file

<pre>
<code>
POST filename.ext 
</code></pre>

Optional parameters (used with headers or queries)

* Setting time-to-live (after this time, the file will be deleted): *ttl*

Example: send file with 30 minutes of life. You can use the time.Duration golang pattern (15s, 1h30m, 50m, etc).

<pre><code>
$ curl -X POST https://gs-bucket.fly.dev/postgres.sql?ttl=30m -d @postgres.sql
{"success":true,"message":"File uploaded successfully","fileName":"postgres.sql","hashFileName":"63894bec_cG9zdGdyZXMuc3Fs", "url": "http://gs-bucket.fly.dev/63894bec_cG9zdGdyZXMuc3Fs}
</code></pre>

* Setting content type: Content-Type. If not informed, will try to detect the mime type of data.

<pre><code>
$ curl -X POST https://gs-bucket.fly.dev/postgres.sql  -H "Content-Type: application/json" -d @postgres.sql
{"success":true,"message":"File uploaded successfully","fileName":"postgres.sql","hashFileName":"63894bec_cG9zdGdyZXMuc3Fs", "url": "http://gs-bucket.fly.dev/63894bec_cG9zdGdyZXMuc3Fs}
</code></pre>

<pre><code>
$ curl -X POST https://gs-bucket.fly.dev/postgres.sql -d @postgres.sql
{"success":true,"message":"File uploaded successfully","fileName":"postgres.sql","hashFileName":"63894bec_cG9zdGdyZXMuc3Fs", "url": "http://gs-bucket.fly.dev/63894bec_cG9zdGdyZXMuc3Fs"}
</code></pre>

### Getting a file

Use the URL data from POST result 

<pre><code>
$ curl http://gs-bucket.fly.dev/63894bec_cG9zdGdyZXMuc3Fs
{"FILE":"CONTENT"}
</code></pre>

### Deleting a file

Use the same URL to get the file, just using the DELETE METHOD
<pre><code>
$ curl -X DELETE http://gs-bucket.fly.dev/63890ee3_dGVzdGUzLmpzb24=
{"success":true,"message":"File deleted successfully","fileName":"teste3.json","hashFileName":"63890ee3_dGVzdGUzLmpzb24=","validUntil":"0001-01-01T00:00:00Z"}
</code></pre>


