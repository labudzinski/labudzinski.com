---
date: 2026-06-29T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQIzBAABCAAdFiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD1pMACgkQXN6UakFg
  winVtRAAo1/B/7KWV6OvPcw2tfQ1v6WYNeJf0Wsf33LnwutyvnBslqZAx75u6VJR
  CSQeD1uWWSwGGyTuq5VDCVupUp2jtzwEYGFiuBJFpwlUIKhrkzUsP+XsRBKreoMZ
  PhWKjPfANdsWs2zUXSTAK74+FuBj1p3MuPJwlJKYa65dtRerVjHgO8iHVCyVt4HH
  oMYvlMgPtk+45lU62wf91tKzCKxlvutCtG9sFMHn8ft2KL2LLTY4Dt79Hzgnsa2w
  W1ntTpXnHsNBkx0cVLf8ACFOHmTvHTtaZw+9HNma593buAk1YZBgfx0fFMATFyDb
  8cqXn/rkIYt/53uaOteecd+4YjOQllU4/VjT3vMx0anZ67BnNXJjgF+8z13RDPA8
  MlTNprl4BJOLGMsnjq0TOmFmXXU8dgaG2Ts7WpmWfV6LL5H+uHgF0GAQac6ZG8QY
  V9vowY2APzJp1u2PubFe+wUcPqC5YNBoINdaqt9vGYRy7OQ8b7v8cQUR5nE1xpSw
  D+akYhWAjMbst6bYytMAz5RTbKZTapaJs5D4bhU2hZrWkOA0WQTRf/uQts2Ae5ol
  dVYxzB4pSuCnKTRmjP87YD8yOpH695Szdwb3jXKNpUC87njJfC8JAzIA7Sf5n82x
  7/vpI7M2VlAJvek+lFdBT//jECkQ8K7C30pejiEbbbjCoJk3FV0=
  =t/I3
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
