---
date: 2026-06-29T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD7DAYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpE3UP/1QGMQjNqgayVpMXvzhBR1ik
  DlJNu8QDASsCJvjxVG3GF5RpC+q7016Z+e/PVIBXc8cLN8wDsJxEe2JMPwqLBqjR
  Fw80h9hqBo/XlDMoaD1ZS9Fh0bpx5yPzGToNaZ3/nhW0KHEPUdgCz9F0ktTaOszh
  Pou3iDeBOAZ/M2KmT7g+rxhZeiKvI/wQNcGXmPohBcqPfQhX0oRTlH0do5tNRQHg
  HCjhX7vLgln+ht+EG4LoOTS/qWe6vR3OMRXBVCaJ6e52TG0RDAzj/D7R2Qdlo18b
  V6GlrEDdnOa041wuuzR6bwJtsR6C4GuD3aoET8XphuxhILU9DAm0dfnXXWpRw5Q2
  HQ4WhG5iuoOmjqRXHaXBr5PEAUcEXVPLDpYIzokLmICrxZMxrK9vHcmJNMS433wI
  0M9z5ZOFZeWK4hjFpPjb38Xs3+dEN8Vg2vK2PF1zV2I+TazW/+/tfH9yuJlX8iCb
  no1CoAhM5iDRGEfiOaXdxHBK6jnBfkh6blLNy2MxB49B+TYZ4X2xFzpcNhW8wcfd
  QkMZ19uVjkcO5eBJISLG0rx50RdH9Mellu5WDiJxtLwDBs1+u8j3vUZp7gDCrMa/
  u040pqX/O1OT2KYD1A22qzyGtmcUBPGBX21qerWQ8P4YP0exSAZOz2UwXPgFAruh
  LnGvf+6jAnIf0e2GkyRh
  =Knpj
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
