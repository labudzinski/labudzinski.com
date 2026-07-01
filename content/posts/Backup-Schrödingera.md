---
date: 2019-04-25T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQIzBAABCAAdFiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpEzlIACgkQXN6UakFg
  wimqvQ/+J2Twh0m6ImRSDy2eKjYXSCNhCUtIBpV6ZnmGMrewuSNaaXqIRaVq8uG3
  V8hIOQOfOBzfgIcJI6LZBHoIQnBrilOr7i/x8sA8iYwSqAwc/xZzAhSULwTijjlL
  rQ+fQZhhlZMgWNNgqrKbWn9d14GL1jOl4q6wTD3thAH78xVC5vHyC1KYfU982J4F
  +79qzdCeuA1kRgCDkR3Ll7b62Au2XKu3TtAWVbP9IQ8HLHZcVpQJ3skqwbD4lIkS
  qJ2fV3UU5FUp33vM6W7jvUd+6434ESRNs9MzxKkpiZtggMiyWmdt+HwxkYlAQF6h
  fdF/FOOdXpCBpK8ynNO11/Zk4IVD0mrIy3V5xxV22radKP11kBCfM+TV1tkYZrfV
  Vn/Zr+3oQTFZvT1iOkzCEZP772UOWjzfzZrTttKyXt+lTUwbJQxr5OC+BJCBRDbi
  w7i0HQZ2oaVDYvRz38q7ukTLeueEHuMfeZR7+Ni69eUH5jqETvTd4wAH0lNoAW5d
  df9hcaf/0U1gY1CdQcmRzm87KmSA1kZZrBrZEYZjcZpzuAah6qQr+QmdYZ9Tuggo
  VXsJJHwk6evVNzs7drrHY9RGGYpurTF1rwIXk99EQiiTuh/5rcATD7uNCq1U/6hG
  BZiik5VnNXBb6FllIalM0pIDe3a0abvXJwsARe2SgAgxW2W/V2U=
  =Iybc
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
