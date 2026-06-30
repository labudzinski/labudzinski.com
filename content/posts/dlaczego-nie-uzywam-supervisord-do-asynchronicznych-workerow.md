---
date: 2026-06-29T10:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD6wcYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIppYkP/2TWHvFk1LFdi8ojCo5lluDF
  bvfrbbe1tGHxGhP5XSWtIkMKzYnmPaPGGZw/IN1/+8T0W58utNGEPAw3t+VBteVA
  sYiQPBFLCQda3ZQj6s3UCn6c6TTbwXUvv/aTHfgvmzDbs6t4BLWfUJdStVpnxfFA
  nogxOZxawgc8SVc21p/ruUiJRp2ZtjQKjS3nYKk5WqcRHKTzsauHjN3QWKiLWLYn
  o/INXr1Vpe8ChF7ZDHxogoxkJ4cO8u+0uWquQT01wdIpGAaKzKwwMj1lFUtNSDcM
  onkvBWLNNocjP09hvWJKLJkMlWgS+BDTi/4Vlc0wD7kX3aKoKRUXUFhD7JBVNGqM
  FUdYRV3efklRzy+OA7X7iN/uh2A5z++lfMf53LB9cLJ1VLukJzkq1qF2qQZpKbKm
  egjo5wqKiJQX/gagsDiSGxVhqooQ3NuDejsj+fqzjitQ8WXJfd/uU8pu0TChWQ9o
  E+DG0/CAglSHDBBHd4OQbf8QaXUQfJ3nPMBmNl93v782gGaq9zhY+KpJ8C7+fBp0
  uoiUbvKmkEHQuOMCYkUjC7rDIhnoRlWLbI61dL5TEjGPw52vs5UaMhq48H15Uaer
  o22nnJ0jrYdUD6AbBcEmQAiBRwrj6pLXGAng5srnooeRBeOCMXPPpYZVwXMMVWS9
  O9hRxP617rDG1r5LF+sq
  =YdMr
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
