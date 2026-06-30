---
date: 2019-04-25T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD2CMYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpmqcP/Rt9cY0cxVCbxtbFI1fJzS2m
  wuXiaEsiP6eT4HRUBMBi3Ee6d8/fdd47enlV1qn3K7xEBzJJ9ulP5oAKyxhX6Unh
  teFU6W9qtz7keCkq+S1+LMoS8NZdILpwtwf23hJkY3jgIGNhYt0SUA38ZJpNCjR/
  mKa2upDRUkG3U6S5ep2r6R727INEnuWyTmS1gWHHraEFNbs3/2XsOyQzVF8jeiH8
  zq0RrQ9ClztPWQQNZhHZ3x+dXYsboUf7dWKK3WXe+pwD4NVWf5yGFzZUHMmYovME
  w8VItmzOkcI5ZmMeryeMVac0TlC9C4VB/PUa4KNvHbiksa9dwwQeboxr1FqmcmYu
  t5iYT3uBiRptdnHUKKlnFXT2IAom4+oH6hGdz6iQSc2CGLkHjhCycaBWNoGwOpNf
  NWmxFhoC/T2o8Ufo+YXqXL90i1Sose0vJVgmynbJh9U2XYZM0D96aOipefkRCs8f
  o3f4pj4mjz8gn9tHwoiuvdIRwrWZre9z6YNneDtVqFcy/MV5nynYjJveYCF+ehw3
  6jA6TcMZ7lCFJyqrx/GKPuHMxLjKvYuyEUTrdoZTQYHvwKciNvQ1gOz7lBvDv3Tz
  yc+S9IJEBBk3KkrrqMRlUOMciDHqsq8BW77iEO7ytLjt27+4fK9yi5oG3uuVtKMz
  GcaJjnv/mdkkYv42i8T/
  =Br9j
  -----END PGP SIGNATURE-----
tags:
  - backup
  - kopia
  - zapasowa
  - awaria
title: Backup Schrödingera
toc: false
---

Wyobraźmy sobie, że każdy robi regularnie kopię zapasową najważniejszych danych i przechowuje je w bezpiecznym miejscu.
Oczywiście już pierwsze zdanie jest nie prawdziwe, Zwyczajnie nie jestem w stanie uwierzyć, że istnieją osoby, 
które nie posiadają kopii zapasowych swoich najważniejszych danych i świadomie narażają się na ich utratę. 
<!--more--> 
Natomiast wracając do kopii zapasowych, każdy je ma, prawda? A czy kiedykolwiek kopie, które posiadamy były testowane?
Czy mamy wiedzę o ich stanie, czy wiemy co w nich się zawiera, czy mamy pewność, że uda się je odzyskać?

W tym miejscu pojawia się tytułowa nazwa „Backup Schrödingera” - jest to oczywista analogia do [Kota Schrödingera](https://pl.wikipedia.org/wiki/Kot_Schr%C3%B6dingera). 
Backup, który posiadamy jest w stanie nieoznaczonym. Nie wiemy co w nim jest, czy działa itp. 
Taki dylemat logiczny pojawia się zawsze w związku z awarią, zanim nie zaczniemy przywracać kopii zapasowych. 
Jest to mało komfortowe i bardzo stresujące, dlatego za każdym razem gdy wykonujecie kopię zapasową zadbajcie również o jej testy. 

Koszt testowania kopii zapasowej jest niczym w porównaniu niedziałającym backupem. Dlatego polecam każdemu, 
aby do swojej procedury tworzenia kopii dołączył zapis o testach, w tym o testowym przewracaniu backupu.
