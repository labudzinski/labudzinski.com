---
date: 2019-04-25T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQIzBAABCAAdFiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD1pMACgkQXN6UakFg
  wimJmA//dBdBTFwu+g3yNBn34tz0cPUq0lAqQJx41t+IhQZ5hAUqUSNcSop6KtyM
  JztcIslzWrqJDCii0wvDVWi1J3FL46x04amB9Z0NwGk4bh4tMlFmM4fyRnp6zXRr
  fKrtHTgOXrYmWq6qNc6Ijy5L7MjrV41J6FIzNPc0Q7lYhqkme7ePl5SDE0/AUHdS
  HT9L04NnShpCDlg3N9piNAPHRPi2lm8vY/QwG5OVXhXAeaStX2s/ntMJZFnmRvci
  dW5nAJjInxvzjOL4Oueg6uS8tP4As7ssq/xKpgS5p7MWDvyvTlY8n2VrhaBTnKEP
  b6xml7oJmOvqfVsBF4SD8xfk6Jwp1rGf/2Q/i+dZa3lzIonkiO3pHLLMMFFKq/5K
  tiseZVTd8TyucINzU9cPdHIr/9VBHJrHKTCZFwdiSfbOn21ayGmZSAkelonMWQTB
  5OpWT752O/wQA8Morl5DudPFeUnOvpuRyI43yDfqT/2EXvM7+6o5mtKzsyIDOLDB
  CGxKpaL/bmDhhfWsYY3evw8FK+WjfKdzTF9EqgZPEqco3W2i/ZBsKyZ2hTET2MjS
  xoR69i8DD3N9ka97bix7SbxfTtGCdKzcl0JZnGypK+2d0BYeQ6GqxXgBHs3VeyU+
  kwevKHTNnGv8vI1nJ7UJ0hG1xRvsuNdDpYPoO19IgPUTFZ/58Ak=
  =eThJ
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
