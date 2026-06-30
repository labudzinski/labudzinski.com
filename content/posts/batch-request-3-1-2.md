---
date: 2026-05-28T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD2RoYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpuWkP/RJQBhJy0qANIVuLTkgdiHmZ
  XzqyjuPZUk9plX6eX4OSIGjN5ZowE/XxtJSGLUdFUHpoabx/aBElXpl7IklOhrPJ
  KTglkMnPPsKHPaP4Jj+1aK0eFjyziLY+kWlkVFcH2IsDcLLkFIsk/adhZxagBp9M
  6FxNCOCMMNkVE+Ol2ZYLfN4dKzll8TyVY8lpoMK1aNrzNsYra1HYgIQCcOxBFKuK
  q6r8ZJtmIiroQskf5JOo8dvfI8ZNH5gdIkuqexB1i6vfeiQpPD7voShxH8XLIMo5
  QjFKGa9Evh4EJz38xaxTIst9H6vut/jG8cUuxEsGTsBK6RqPeySOC+dyyiZ6zWF9
  CVx4MhYHC/Ck2wj5+pRpggrKLtXTeIyoByDzL5efUq8YGYTXTLu6099eyFvQpy1k
  KYd9y8UppjFi/38v2V12ki8cnM8choBq6ONKoxIGK9gklbTS7me59qSodCn9VsJ6
  +7WhBD9mUm8DfhSBKGh3xE6y2nC/pociA/z6T7YhdAo2MTQy2yzCmDK5+R1X1s/s
  DG7e4vaoYKenXyUO6qAmsl9fgWftdMziPpUDK0YFsmawc0blbvAnxdTx9aiIuUtn
  v51BPeOcOEn0aQ0R6AtVl2JbgguXDWi5dLBRYIi3Tf3TiVHCRE6yxvA+ce4k4Jg6
  ZWhTNAAQ53l8v1qGoxYo
  =5qWu
  -----END PGP SIGNATURE-----
tags:
  - PHP
  - Symfony
  - Laravel
  - OpenSource
  - API
  - BatchRequests
title: batch-request 3.1.2
toc: false
---

I just released version 3.1.2 of [lemric/batch-request](https://github.com/lemric/batch-request), a PHP library that lets you send a single HTTP request containing multiple API calls, with independent operations processed in parallel.

Here is what is new and improved in this release:

**RFC 7807 error handling:** Failed sub-requests now return `Content-Type: application/problem+json`, making it easy for HTTP clients to distinguish errors from successful responses. The error envelope stays backward compatible while supporting all standard RFC 7807 fields.

**Mixed content type support:** A single batch can now freely mix sub-requests returning different media types, including JSON, HTML, XML, SVG, PDF, PNG, binary streams, and 204 No Content. Binary payloads are automatically base64-encoded and marked with a `body_encoding` field so clients know how to decode them.

**Symfony Profiler integration:** The package now ships a dedicated BatchRequestBundle for Symfony 6.4 through 8.x. It adds a Batch Request panel to the WebProfiler toolbar showing per-transaction details: method, URI, HTTP status, duration, memory delta, request and response headers, and the full response body. The traceable executor is also safe for long-running workers like FrankenPHP, Swoole, and RoadRunner.

**Flexible configuration:** A new `lemric_batch_request.yaml` config allows you to set max batch size, max concurrency, content length limits, header whitelists, rate limiter binding, and profiler toggle, all without touching service definitions.

**Laravel support improvements:** The Laravel bridge received updates aligned with the new facade architecture and service provider structure.

Requires PHP 8.2 and supports Symfony 6.4/7.x/8.x and Laravel 10/11.
