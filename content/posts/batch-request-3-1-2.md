---
date: 2026-05-28T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD6wcYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpjj0P/A2LzbEWtyuJiOL0Gl76dAB9
  deH8yOGJu1LoTHFRkwvGgbHj0GWrI6nvogMVGPGANS633Ll4x8uauBTrgAvOrihv
  LtJE6E8/xxG5IcgBtKrjrH/cO6yRTuD2Too0E4HXD73xQ2/engZQVA7BfvP4sB6q
  tVNTWbLNvQi95lUN/r9FYb70bTH5FCEa6qqTzsRxsXiHkOyhr9P2Tty37TB1yxLy
  NpdrTysKuC6q0M72fBI6Op6U8qVE8hRqg4d61pf3w4vMePo7DHwv1QfUA6yfPDoX
  ztkZh7T49bKkeu2SHYMKe9wIpKDOgiJgUXqHvC5mfC4rbHRn6FihAiegZRx7kgAN
  REgud9NO5PMjn+NlAsuh2RmFHrnKfgNzMS3s6rkW510k8oKm+Ndme4NoVI2D89Tm
  8De3rpseCa40EDFDUHfQ1kt/zza7LTD5DUk3zhvQMLO0Ad3YCN4j20UyI/lDAd+B
  4/ENQKEOQbPTsRy4birGPrLOz1Df8vQFigTXSxkzuTyFybERx9b9qIvr5lVOVICv
  2Pqtl2Am5kFj3vSiZuFCe9ZxfuX6vDtgo3xyQkPa+mg24EGyirizVNi2WDywWtKL
  HA8JHwlsxIYDntGsx8JIviF3RiXM9so8HNUnhHYpzm6T2IgIVIhSUXrZEMDYFAfx
  +B1Pou/+kntmsn6VjBuS
  =7SzF
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
