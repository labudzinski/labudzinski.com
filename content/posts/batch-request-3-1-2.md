---
date: 2026-05-28T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpEzusYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIprXwP/14sqzxJG8EZ4OxgYSASWW2x
  aOELSdPajNKueBgtEV6UkImI7YXfvZax0kIfsza8laW8+hX07RzEL41iZM1QobcQ
  MQyxIgNImHdhtHhibCAlG8RhbfI36sngfGhnlafgNTQDwOapD/HOtWjk512rafZ0
  jaAdoDrUDsXm5YSwsr28k+Sik7D+X2gTLaot+hb5TMmUpHBq6XHt6D7DbBwKnFcd
  ho6jRR+WUYN4KY9URxhYZpGyk3+FzarWsR/RjXUnWwIngbHfTFKc79Johg05BmAy
  aeIdeCZfC+wGdmk4g5rYqzYfRELvRhzJOKptE9fVPETQh3mgih26iZ7TZyzTSFUV
  QHslTxS1L7t4FhJGD/55F6pgFk0aOYQa7GuCHbZEGH5WiDOee2KqiOu1Yx6WyvJ2
  yHMNyPFit5ApSP+cMjrKnRKKRldBDCUitVrdL3yXybZouCQNxH0DKo4J23Hr7vSh
  r7Dc8PlMJ0Q24fEbvIoaoZFlIsTxyEwsLvmxxx1zC8pQChsA9NwnT5XzZFtLU4z8
  u18dDfyNA8lG4ZSWVeeFtrUBu4Qgdg26Aw8Q++FAB+TqH8P7iFM/FtsRY4CJc48j
  JAdjnm2kCXBD4HOq0ZfJxTHquMiX5BuQQBE2pvZXpANDQEmMrCNuI/f/RNCAR2Pe
  SUCza9Oe322ODyqM6bHO
  =ntTv
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
description: "batch-request 3.1.2: jedna ramka HTTP, wiele wywołań API w PHP, operacje niezależne i przetwarzane równolegle."
---

I just released version 3.1.2 of [lemric/batch-request](https://github.com/lemric/batch-request), a PHP library that lets you send a single HTTP request containing multiple API calls, with independent operations processed in parallel.

Here is what is new and improved in this release:

**RFC 7807 error handling:** Failed sub-requests now return `Content-Type: application/problem+json`, making it easy for HTTP clients to distinguish errors from successful responses. The error envelope stays backward compatible while supporting all standard RFC 7807 fields.

**Mixed content type support:** A single batch can now freely mix sub-requests returning different media types, including JSON, HTML, XML, SVG, PDF, PNG, binary streams, and 204 No Content. Binary payloads are automatically base64-encoded and marked with a `body_encoding` field so clients know how to decode them.

**Symfony Profiler integration:** The package now ships a dedicated BatchRequestBundle for Symfony 6.4 through 8.x. It adds a Batch Request panel to the WebProfiler toolbar showing per-transaction details: method, URI, HTTP status, duration, memory delta, request and response headers, and the full response body. The traceable executor is also safe for long-running workers like FrankenPHP, Swoole, and RoadRunner.

**Flexible configuration:** A new `lemric_batch_request.yaml` config allows you to set max batch size, max concurrency, content length limits, header whitelists, rate limiter binding, and profiler toggle, all without touching service definitions.

**Laravel support improvements:** The Laravel bridge received updates aligned with the new facade architecture and service provider structure.

Requires PHP 8.2 and supports Symfony 6.4/7.x/8.x and Laravel 10/11.
