---
date: 2019-04-25T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD6wcYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIp8BQQAIcz7bweqD3wU7/++m74UxsK
  mYC9ZoSVmby5kEZE+4HJ2SJco0MqSRyuzueDhdaDoWiuEj4cOeZF//B7KV/NRPO+
  mmegWBVmS0XFl/leyzQWlJd3IGOqoKt2atKhVSRiiSjWqZWZQ5/rXi1UOF8GW4Ox
  Kns+I0V578mMQ1ZYvaLjnfN+ZyPCvDr47Oba2K0/1bpkSTj+/8uu06ixVbLViUbs
  vB7aXjgeg0gruu3vTzlhs5B1c65yRPpLSSqXS1dVmTGOhSLOXsAIXpiWIihk1yY0
  ydVRM+gKCCEqeBjyGDgI9WERgm1GPyQGZJ0a04Y3k4sO9aDd4dukQoaCnMdnywIu
  gLp0dKxilc2S3Z84Cm952TNEP7wPCxr458lvQ6xdnu9I4uMRSfJ4H4NA9nH8ezh3
  mZpUcg6DSnj1w6Lq++/j9vFcnS6KGaX+rle2P+M78n6/Hd/IUihn0YtI1M0xmtZ5
  UlD0cRfnuBwI1gP7Wv1ulyQac0OEZhGO1k7gqU6dLmQ7jVHDRhuBjCZjV5Z9/MSh
  KR3lWLJdYnalVEU5LoosYp6MFfwJ2diQbdJWw1FCjKBfcSFYtNvxN7toneV4ZJjA
  PjECAh4PZEy9I9SknUNm8gKUeGhQwMhhJaWzMKgID6h3zQwRZMr+E0qlAbKVE575
  DSi+qQgjyeBRHo4kNrN+
  =7Ovu
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
