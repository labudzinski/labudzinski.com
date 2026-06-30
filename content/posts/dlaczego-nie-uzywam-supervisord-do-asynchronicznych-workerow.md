---
date: 2026-06-29T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD6eIYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpea0P/A3Gbl46cejwLfQZA1XzgfFU
  4JCRZWbXPlzqtblCql645lzazAG8BenkMhDHziFMPoQimtXwDS1fO3FYdFdnc4nN
  8e1pyuh8nMsJKT60qOhFW/PbF0T016Y3dakZ/KJ6WChN//obtyAlb8gmuEKCrNc2
  Qr/upIKWSGZDdMV2SkdV8lW4YwgN5AgEnPUsaUaN84WimKq/TGGS6tDZmoHtAW5U
  IvnMBb5r4ArvRsEQVaxfuJ4FcH+qsU1jFsfC0yvt/C6fw8vhd5kmQapqNONbNuMk
  6l+eAnQ81z1RThJlOyrLCVHQMHsA5ohC7HsZFqzm8cklx+u6ATBL/gQWQgHVkoS6
  J97htrO6nXwvcwAxxRATgL+by89phJc7arvgql4Qj+lXbEEnG7GjJhdwPbuS1xKk
  IqxkqvFUY/nEV6UISdb+cpAjrvolJSlbmZPBnhKdIiXAkfmIYmQGD/6eRMe+4LsS
  nrunhqnu/i09ffikeZfBt9lQSSopXQtFO4ooBhC6SftsIfrDs0PBqsv56xxQmQxk
  JwrK2peT5nrvACsx/anqSXMWG1193VoWyAHuqVqwG24mXhZAI0bWTks634KQe0s0
  tBBrBnWUVnw9T4Gev1MskTvjZSTOatULZc9Q2T7En823XmyL0pEdikx3RoM6yym8
  q8H7lSVsrEzeOsWzlxh9
  =5CLD
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
