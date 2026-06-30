---
date: 2026-05-28T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD6MUYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpk1cQAKX6q+06zEo2AkvsUULB3XC8
  pkbJub+9ojFyEQhByeYqJjH8uUf/4Lc4gpE682392uytuO90cH5q43mq+Xlu+Hy6
  rq212g8HsnrJ6auLMQikeYyETiSPJZckLOxDqwLzS7AsIJPNtPVE1t3u1Gehb9oI
  T79X9LGcw/z/8c3n7KaOetIhiuQc1ymw+oq/NhgfSceuKGGHtO2B4wEhnb9bVqoV
  qrF99V2pabsVHzeaYzKbXviK+gEgcK1LqpOny89hjRXuJ83ZYqNkzdHoWIH3scNl
  kb3t7Og7mT6dGy5HlgCpRMeqiTdhaaDOyQ05d1EU4+uYBz783Puc4mU0GRBpJ9Xb
  qwlA75eQY6bKN/MTsu08nxzyiFoZZULYurJUgfWnzPZH1dGPrvxOABIFjkni+KfW
  5Qc2Kh0+j0XVhXkBDfCggFpiAz/PkF5VuarKWc1BUlB7234ioR4x26IYxhxTyvOP
  rT+NOvE+9aE1W1BQv+yh5Pa5YCcDKuax7/5IsLQKxje/5+M3i3+eZlPqKRh08SNy
  99rEypD9JcP0zS4aaNSq4ipgEV33qT/BdJCgsWAKSfBlbMdtxGw/kuSihiAcbCJc
  fdjUfPT6tU6xaDMe4kPK2btv/SwSUJgf7QCzIDsIKC5ZpG99KP3WSLiIHhoJJ2vY
  2L3vLFwaBnCOrUOzlp7W
  =30bH
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
