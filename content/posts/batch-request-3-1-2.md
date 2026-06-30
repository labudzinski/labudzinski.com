---
date: 2026-05-28T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD29YYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpiFEP/2s3+Yje3D00SUHL51A67qnD
  g7xI/L54aI2NehLFJ1IJn8ZmozwSfqIDxNg4WuqbxJpWwpdSgDxD4ABQYQLJy6Tn
  DrGXRVh4dPgJSw1u6ZJkXEAQWlFKO36bcRWelwJNvGsg/1vE5TfDF5V8NhHEOI0I
  6Cr4CoUUWggLUCYq1rLnXlx3fjzgxz6kWGR0UP75ELgUKnjbwv55jvMfeNBB79wx
  AW1zrqYujnvd7eyi7bmEolbXhLhNev+0JZW/PeCAZn3t4Xyptbzp1UZQQS0WzOo9
  xynJY8QDdE3fzlCOZ1Lqw9pJjp59ROsUTggkPjzpmSshlBm7zRj2K20dwg4nBu1P
  gQwVMEhI8MmvR2ho/XCVhvL2OXkWZ+Bs3gVtAXgrkyGo3Aqjmfkg2oeIxmwmRq95
  uwdiI+t5yscUGkCGb9e6tiD1csFulsaqeAKto0ybKjQF4WBRkTXaVLjQzS5+9K22
  ID4p9YAi90ObHvkAjS9pwZnjF29AUrg588ZxIC3DcBU82Z1VaYLwQJ6q5rAV/4IW
  zWlcAYLiWAv+/hmoj/h+F1fpKbbVpWZh069gS3G/v5rjBKjrLkPE59BWOaTdrdF/
  pnpyMhS9cr9wjhrhbE+/MDrSHcpa500Q5aGP2f4ujpnWiD2E5K5WAdzjQskXlgip
  wbb3qdYYllaPDroAi9Gr
  =SRYg
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
