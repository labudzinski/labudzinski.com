---
date: 2026-06-29T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD29YYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIp8CkQAIGOgou4o/GJI4jaLBsdPLWj
  T+UwzAsKSPiFWbGKZ7c3CgqV6z+snZuQhEJLtdt51oKz3u943FeM6wA/2khn/Q+6
  dpsx+NDKUAFop6UTzWxdZuyvLpm2yYv4+aXkw3OnxDPmyr1jzut4YlaWc+78MivI
  NC7svU8QoQ4NfNQ6BbxxmmLPEM3wOpPcVyRYdV53hfME/6NxmlbaK2SPcKVINDL+
  J+dZpAXATU6v0UgkPm11ar7UTiRAGVs0n7VC0VM28GmBHn8ZED0YIjxGmxel2lsE
  m7neSxeqMtZ0bs0VbzvhW4/XbECk/WbYsqCEn3mUDq4auRbXbYvAfm1mXVZG6gj1
  voKhwkXcHRdVAkmqZiDu9ckafKa5UoYSYH7Vld8FKPt4Q3JE1gjetf0vdMwdb1h5
  2K65E4Mh5Zf65wWegXiEuQZtRrI0WG4e1jrwLt9nSaneQWchKUJ96cEE1XB4Muvg
  BzCl/zlHYuZpOTsu6mFMFrMpMg04K4Zd3rJkaWfehLI0d2NElo259jBxm8w6fwbu
  bmobrpW/KHhb1Czbqfk175D5kABBQR6etyPvTbC4JhR4MWQgXQ4eysCGWvAJi1mk
  kmF4vVgfVpAvwQuq34GOxoJX1nJqT/2SzRoQOvSAJE3rhuuOPCCzoHwHR1wr+ZNQ
  YaParGOQp6WNvkyzuaIL
  =SBXa
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
