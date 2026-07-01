---
date: 2026-05-28T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQIzBAABCAAdFiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpEzlMACgkQXN6UakFg
  wimm0BAAnFjJ6Dx/uFBuFVK3ykr1H/HAqVVQold8fRG/ujWWrK3jJH7F78m83fYY
  RyAzElktxqjbnZlf5sEauXgyeioi1YnIHHar+LC9sGCDREOJQdUWd0uFIaS1gG0+
  q+b0sizJHXkDsI4Djb9ix3VG7pQN0DX9IQtFVS0Dm161jlxWGV3Fp5kK/AafEOGk
  ZBRcYAEB0HpckTo/tJoEZlqL9KGR+Qc3Y4m51mwWB0wcQ8AmcWoZKh91xnH59bVJ
  msWEZwy9jGBYyMiAtOXXn2KihGsQhServyT/mySzLhMtw64sMAcKvBR4Zaha2dzd
  jC3sPP0gcP2rWiLkv6XRZMIEeXO87radPpv1fQY/FThqV5WMcd4iOxyvMwR88zVr
  2Z1NVXm5YJgTO/alaAVzwP9ckELVgOBUb21qnN0Sdvgmo/Kv+NaJO6YeGEQqY+b9
  xJLar7StbgsCGahHpsI/k4HsEu/J82Ws7/tv5oycs8JrmWXoCmOr1bedkZ57v5MA
  G9yLlL1dmQPSla2eryIKPYm/YpCt2vxvx0oY7Z64XRzlgvgTZpqxR+8QoC3v+L28
  98hKNo4He5czcsdLZAdkqDxDgyLtzZDGb25M6kGEm1haFv0yK7WkgjsRDr/R/8bF
  XZ52FGbfVukNRXp5maAYhbEz6Bx87rTYfWJakKuXo37La13QYH4=
  =Rxsy
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
