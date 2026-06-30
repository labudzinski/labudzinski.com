---
date: 2026-05-28T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD7DAYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpqHoP/A1fm/o2Zjvj2hbLp6tP0YDV
  7/aGsiMTa4rmLbmpX6ot5amTMI5OEeZY/nF9/UKKxyeshpvw7/+ftLZad/IbO310
  SxFv/ucWwFXQQyZEtr7a/E56m+3MPslra3uJhlNCuhP7jsNbF2WzbLfU23RzwDPN
  esl7Mt3NbNpcyCm9VPpM/mBN8gBYqkoYEEs7zlEUDa330l0FJHkPd7vIn7fD1dp/
  xmCGhxt7LIDLQqWeOXcEDo+lqpaftMMjvolyXUuUr1hdjfVYQAMoXavtfGwQKO+1
  TZRhoPOHnkGp0i2P1CgMJyXLnkzhvEaaM2m5Vgq1hvP/T52uKXmxCMJRXp1df3xP
  ag9tQIU1s98p8gM9xNnWiOshKfM/sHoNf3TsdaCZ+Y+B3O5FYoeV8cjvyPMEgpSw
  kzN+XTtNPtNQJ6tr1HV5AF9k7qhYuuHJqzLhJokgsncT/hVGxqzo+y4C+i3DtSnH
  OerUzmE5pt16Kx04coM27XEQL6iB6Iw9WE6LC8zo7y+DMjQnV1MABQS3Q1UePOa8
  vARjyep9CSaIhUE+lpCSwEjyZ77MyFXYyxdLkC+CAb6f+nhDn5P3JmNvVE84DEVb
  L6pVfFb3hdulCPjMhXYE1B/lBX533HJYtoiIlAJ68j9DGpIHte4nEu0N4/KI6+Cq
  kVVF6nHWx+dY4ssSeGQ5
  =ER6T
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
