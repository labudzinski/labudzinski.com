---
title: "Dlaczego nie używam supervisord do asynchronicznych workerów"
date: 2026-06-29T10:00:00+02:00
draft: false
toc: false
tags:
  - PHP
  - Kubernetes
  - workery
---

W projektach PHP, w których ostatnio pracuję, asynchroniczne workery obsługuję w sposób, który na papierze może wydać się dziwny. Zamiast supervisord lub podobnych narzędzi do zarządzania procesami wewnątrz kontenera - co godzinę po prostu ubijam cały pod. Nowy startuje automatycznie i kontynuuje pracę.

PHP nigdy nie był językiem stworzonym do długotrwałych procesów. Nawet przy czystym kodzie z czasem rośnie zużycie pamięci. Obiekty, cache, biblioteki, połączenia zewnętrzne, wszystko powoli się akumuluje. Po kilkunastu godzinach worker potrafi zużywać wyraźnie więcej RAM-u niż na starcie. Restart poda resetuje stan i utrzymuje zużycie zasobów na przewidywalnym poziomie.

W Kubernetes takie podejście wydaje mi się naturalne. Nie dokładam do obrazu dodatkowego procesu nadzorującego, restart leży po stronie orchestratora, co jest spójne z filozofią K8s. Mogę zrobić graceful shutdown przed zabiciem poda, zalogować stan i wystartować od nowa ze świeżym kontekstem.

Największą wartość widzę w przewidywalności. Dokładnie wiem, kiedy dochodzi do restartu. Unikam sytuacji, w której subtelne wycieki kumulują się przez wiele godzin bez mojej wiedzy. Mniej warstw oznacza mniej rzeczy do debugowania.

Zdaję sobie sprawę, że to nie jest rozwiązanie uniwersalne. Przy bardzo długich zadaniach lub przy ogromnej skali workerów pewnie wybrałbym inne podejście. Ale w projektach, z którymi mam ostatnio do czynienia, ten schemat działa i to właśnie za tę prostotę go trzymam.

Czasem prostota wygrywa z bardziej zaawansowanymi narzędziami. Nie dlatego, że te narzędzia są złe, tylko dlatego, że nie zawsze są potrzebne.
