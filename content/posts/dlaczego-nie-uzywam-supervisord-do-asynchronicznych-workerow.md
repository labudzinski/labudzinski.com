---
date: 2026-06-29T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD2CMYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIp1vIP/R2Ln/vdMiI4uTU2PXGm72Qr
  tWNveE/ddYO4GOxLDkVqY6VpcykVM1mCe2b0AuvG/TYPcBkriG+UTQGONOZHGnSo
  a04ZTcVtQG59s3OEr1MVs3JBPgbXRnsbYqzW0sxYS7qbWMLvxz0LfV2aMoZbkD+7
  NcZlEJgSkZdn4ujMcWNtY2EoswHaRM/tyNKpoNlNU0JfVbg+N6sDtXEjCFTQBM1t
  N3bPoRY5Hl2kqycURCdQAln+SsCgMagw1IhrsgWAWYnTmPb7wCbDWKe5yVuWDe4g
  tAwDuUwM2MyLBWJQUyIKdZDQ9R7YepZY+a0/PR7NQ3ooZgry387A0stVA1EGPnZh
  dW2npXm7hfxuIWQw5JbPkaQiA8qhlvuHoajlCaUg1AP9aB52Y9sDpUU56dhJarkL
  DOnHTAEnu1yF8567+WJVdYi4qy/m5HOM3mGoqmcjbGIwKkmG2zLK3xkbgsg8GZcd
  QYVmoR0opdu89aE1kTV3d/qDw+hfFoSedPd8s7oJYWcyMqFm8p5LLLhxw8DgGjC9
  yZreapj4GxfnkHoxK5rKSDT6CVlzvIkGZso3LLMEzBiqThBHQV89Pq7GgShBocP1
  emoeIfXIDM71/aE1sZq9suSrgeIWz5fBp1hDppyzmOU30Cd6mstTnt7wVVnrphil
  V3GeOFL4P7M7D9UZI00g
  =evbA
  -----END PGP SIGNATURE-----
tags:
  - PHP
  - Kubernetes
  - workery
title: Dlaczego nie używam supervisord do asynchronicznych workerów
toc: false
---

W projektach PHP, w których ostatnio pracuję, asynchroniczne workery obsługuję w sposób, który na papierze może wydać się dziwny. Zamiast supervisord lub podobnych narzędzi do zarządzania procesami wewnątrz kontenera - co godzinę po prostu ubijam cały pod. Nowy startuje automatycznie i kontynuuje pracę.

PHP nigdy nie był językiem stworzonym do długotrwałych procesów. Nawet przy czystym kodzie z czasem rośnie zużycie pamięci. Obiekty, cache, biblioteki, połączenia zewnętrzne, wszystko powoli się akumuluje. Po kilkunastu godzinach worker potrafi zużywać wyraźnie więcej RAM-u niż na starcie. Restart poda resetuje stan i utrzymuje zużycie zasobów na przewidywalnym poziomie.

W Kubernetes takie podejście wydaje mi się naturalne. Nie dokładam do obrazu dodatkowego procesu nadzorującego, restart leży po stronie orchestratora, co jest spójne z filozofią K8s. Mogę zrobić graceful shutdown przed zabiciem poda, zalogować stan i wystartować od nowa ze świeżym kontekstem.

Największą wartość widzę w przewidywalności. Dokładnie wiem, kiedy dochodzi do restartu. Unikam sytuacji, w której subtelne wycieki kumulują się przez wiele godzin bez mojej wiedzy. Mniej warstw oznacza mniej rzeczy do debugowania.

Zdaję sobie sprawę, że to nie jest rozwiązanie uniwersalne. Przy bardzo długich zadaniach lub przy ogromnej skali workerów pewnie wybrałbym inne podejście. Ale w projektach, z którymi mam ostatnio do czynienia, ten schemat działa i to właśnie za tę prostotę go trzymam.

Czasem prostota wygrywa z bardziej zaawansowanymi narzędziami. Nie dlatego, że te narzędzia są złe, tylko dlatego, że nie zawsze są potrzebne.
