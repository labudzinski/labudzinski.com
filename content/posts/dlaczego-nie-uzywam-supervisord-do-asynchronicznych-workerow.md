---
date: 2026-06-29T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD2RoYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIp31YQAJyX5DEenOihEe1aTXYrTxu2
  ut3xBhrf/jBvxNZ4qqjG/bq7LU1AESPUoETOAIZW3MPe0GZVmiyFa5ndvap/tLUQ
  4Qe/uAI31mu2XOSqpMHjR+uzQ3SKqvdOSoZ1ciuwkF1BBXoIwsNsXGyLI/05mRZE
  8SWlxlldlUJSMYjkUeBCJI0qK9DEyhMbCJ+YxhJWuvmwRvvxyIyzZsKMaq+TZzEK
  vx60p1AwABNZcILcMj33y7z6IxV+waaXIp7OvZgOlfmaJXxzU/ir9T88R4plxyWw
  f6JOFDIUTcnE0MDA9+3gIryBbQooC8Mv+kVaBDekNKyPwswqKot6OpsXPiIyP9/t
  XMcEINgG4WEbMB5BcMgwuiJHFObOA0SWQSzvHTnZ9do9WJUn1APx0VBZgeFKgAiQ
  t1efX+OL2hIg4hoazN30FTt6424lfupVDtQ4NXMpDP7VZGSwaVbbcLp5MXAqUPBn
  5P9hg6joLeDkArCpUBOF73rivcGoGFfmyMhqFSUPNNVTbZnHUOao2uHYDLGyrBGT
  LmHjyEnf0C/lGbNE/4lQ5WtJ3HcjZV3DCeJfqAhgq0SmxYCQqDGPntumTWR5qihM
  203hiMV1FW68ys39K5Ktpb5fD9iaqmblcyMi+XQ7XwXofBZyHkA/qERaTpb9X7bd
  L8QZJ3G7fQThmJFw7JZO
  =B0ds
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
