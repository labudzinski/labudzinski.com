---
date: 2026-06-29T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD6MUYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpRB4P/RD13r2ozZyeGAglEux/b7G7
  x4QGTwO7L6vEyNctzo+AB1QYsEsOxCIK/YdzBQr8YjKz+zWh+Osj+RblyJ86nU8C
  jD8vpqPWShQ/9fvnlqiAZcfChVE2i03VxATIRE6wY9enPDYqv4fn+1aChXJ0fUKN
  7Gh8Q8rDZfMlmT4ZlDCxw37WjRrcD62jPd09kxCLQf+59fcD9RAOfGG3mO70VLct
  Y6070lbPL4ugPYQ7yJ40Y8v8JS6Z0C2jyXOoetqowHP21eMNvwcFkYhy7KotFW/1
  MW/m9CMIzvb9z4aYxqTwg0UkU/F7Y7ZvD4PkzVV8rMy9k+loBFIDDaQiROlVWaaA
  jLJsXiMR3JqUH19ynkvUP/+CSVSGUvdD+KtmyHCl0UC/BwrQn7hMnaZyKOdcmx+K
  8ePWIELOq7+rG4dtOege4A/4XSeAxanjZX+gbPhRUPeW5ubBSFw3cHM1LBqnmWvr
  LgoXWHJLD7ZXAJ8LMp0APWwVHauoqHREmUR/e/4cGoxtLznjIIjANtTlukvVGNW2
  QWKwvnUYBfCzXl3Kd9ufTmEf75/SkG8YnVmGfyMPMJ2biPSEt7beVD/FR3qel0FG
  3jj+KrFKtHdFm0uwUGrKfCnOnnnV2rivlCvV2xfqNfPip0/a40Jn0MDov33Jm+PK
  FuhNWOLnFb5YZli079kE
  =IgTz
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
