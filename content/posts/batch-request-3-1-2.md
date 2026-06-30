---
date: 2026-05-28T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD2CMYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIp/vkQAKLHRpmJxZUuF6dt83QhPKYa
  7rZQyj8a5tZaF5JjDUcIgKFLAkvbcO8hjc8D4vMw7WVy+OALT1lDLKBkE4BUlgDR
  COwhGBx5NbAqay8RlrauseTWCbbrHQFOM54+kYkUCFDIVq1koVxhEGg57xlYWTDp
  tAtlGlKEdb3j3a86WA4rxLnyVLjtivJZhuhDanRhppBvCieGfF6udwzr2DeaVWHH
  SdFiLflud+JZuUmfrvy6dqEMXQrcAds08uJfA1DSYjVbreiP4S3BUmiwMRD/XPpU
  m3jkC0qfTcqKJRZHqHCm/JXiKY2kSiFdANYSD/dYKjI28UzfirZmY6QU8fm41u3C
  pgtZq4dkhmf7YMvUls3Z3VgF9SRotQjOHvhhTtLQwNWnZHYVedMSAmXPH6AIDeSD
  99mpoK22WsXIbXeSXMHSMjsEFZ1Q32KFkQBWocsbD8328dnqWfpNObKa/WJZnnLo
  +AjPgBtLrU2lch6cWXANsAR/0n+9PW6yi0GTm1SgilT6VGaJBCYKsCVsJt8cJQI0
  hiCrKE0QYVr8rG9I6Wo2HGsX7OUR8NIlsvWdnBMOMuG8qFgvSyjUbompzh1DC8Uj
  8v4nalWa3vORB17bA2BjbdK2H8zNPPzGh8rZxVosT3OyrpdsWDW6guSCpgdiOOSB
  fxBcw2Jd5UF61D9D6R04
  =bHgm
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
