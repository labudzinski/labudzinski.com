---
date: 2026-06-29T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpEzusYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpwzkP/1vP9iFEGahnywltIkBGVUFI
  D1mAA+5stFFvkPU/hACS/z4/qCIKRS4Z+np68VQkKE2gpf3svqtFbv1ZMkwtjxR/
  +Ak+3THlE9Lm+P3eIZbMtv4+6wem/v44VX4acQ46oK746FEk0KPm+A1uZmPBZBdD
  I5oFpOiHrz98N66xTlsMEzZsoFIO33aQbzYRup7iN89g5LDuv4xk4WwZFULqrhm7
  VAAFO74SNYsIeW0AAPePMfHMGLc/QumuhvaWMzDYlM9OA4zPaQhL44q/oH2T55oZ
  b0yhnzTKXL1yrwoVABsbGHfxxcQVCVXp+sEH6a6WAZIX79jw7BgU4kjYVH6SM1JM
  lcdkekKe59UqFCS44NePrX6z99Sy4c2R0JxgBarzGZvCQXXK4XNjfaTdHIcvO4kL
  NxCb2YjwVNY1GKRauh4/OQR9s315vgbcheD+MlfBXBPuDl0vigO4+C8EUN1MbFHP
  QEDfi1IWcqDlFAmGU0jYxZ4vCVfJpeU690TNA8oCggNDvdF/YFxSLMKjdBA8VUxf
  dB0I/jQVOB3Ff9wsow32tvOWMk+HJUdjbkChgzrPfaMT7xvgPif41t+nXtmmwp/q
  GCYMuPaQp4tjR4hhnPN3N1Phhw/RWr2e1etTDdh2t6HlRujQbj31tcBZYiSozPyR
  TRFflGV8wXQ9ObTCoCy6
  =xZ1o
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
