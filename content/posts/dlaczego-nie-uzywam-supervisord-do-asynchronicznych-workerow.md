---
date: 2026-06-29T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQIzBAABCAAdFiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpEzlMACgkQXN6UakFg
  wikVRhAAoDvaDI1EpDGxyYWzWsZOTmHxTcm1NE4kOgBI/9J13uYO/1iRauoKxMoa
  sYrQ+baTBVF6km9mzxoIIjVLkXbfutzlBPuMm3++z9Cm8kVztelW0B6e7iStc2dX
  5ZpB35zHVTScwavjjrSuV3BoBkTL8CQ5d8eO241Jbu5AvF2Tc54vel3ToKZDodl3
  eqAvQC2XedYU+AIKtCPz9Shef7aIsCl40oPyrDjS+cv57zsi+i120acyHZJ5EAO7
  z2+5lQ7tepXWg5Ao59z2meQ/8NI9JXPTJbtpgEE9Kjz8Psbkn+8omWejYLBH0egd
  IYWjxvcEgr0HA7u3xQQCtx3juAxoRGH4s+KB9kJLmMvJOlRWZTI3ryv5oQHppdN6
  RYGOq9M7f++PF6XWNIPWtT+fOzU2tZ80o9vRGiIrvoqCD84XewTTnfx2f0ebelrL
  Al8qYvNnUpode3AoU9I3gEyMSf+eYOxkV1rCJ24Kqy6+VSbNXqUQYlpi3ijA8ZY/
  XhUa6paDEZE+yw6CcO+DjWioiYwkmRizrxVoRYmuXMK+xHW2/wZAd0BtOjXWH/AP
  2zDqmbLEFv0i9jBs8pf0lMP1615dtoWD2EadZsUV+SEAis+Zu7MJOLePdw0Y7hUi
  XPedhupNhM02dOsLMM5lBggs9DHsP2z7QOhstY72kytgqBXDhmE=
  =fxBS
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
