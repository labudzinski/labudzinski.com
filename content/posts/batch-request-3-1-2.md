---
date: 2026-05-28T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQIzBAABCAAdFiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD1pMACgkQXN6UakFg
  wimuSw//RbXLSfieGwVRBWM7Glyqz8uOvLqWWH6IWrhPcNgkuftwNDfXjYJbcIaD
  RjWKpboT+rX0lNlOglpMUN8NeRxszhkQpVkmbPYDyThVdPTO/AI4bXa+2uWRPgqd
  EBtzlhnOCnozS2KjMPMoIvoz+jAWTmODNjWPvLjc8JBGS1J7Yh4qmIvWvE9qex9c
  zhehdpBmB+3CZ1AywtnVHkf1I13xPUFI0FLXdGrtgSlKEdHROyzWMdtIWj/7A3GE
  B1nSWOXEyaQ5EkDbkbnCbt69z5RosJMlD/QXu20624lD9cqPjwwSFnf8eRshs9MO
  41d46/K82fPARHlT1t2LKCU4/+/w5XwGy7t56S96xx7GXXnumxJZ7JJlKdRSf5fx
  Nvni674BTR7GcM91xCC/rV3Qsqm47RqbBIIfZiqu9tYO4wRnATD8ProvvcswUMPX
  95vsAXs6nELtjotKlkkC8r5+6WFZMWjRinzULCQXng90Mv9NvZBaWT4nb42ccJOg
  E41mIsX4RlDNbcHoDiJEUeboX29TsMMfcToEAAU4eX8vED0DSI5RwsqS2/12itMD
  ySRuEpxo4pVKfsjy2UpM+LPyqd8tWboScfJnfqdqkPFL7Ra2qr2EYOFbe2adnfZr
  bEtSysL0231YcAvWGsuK21cdEewVjHM3P5neKzMY2LqI/ryCOZA=
  =y4gI
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
