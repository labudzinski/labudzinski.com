---
date: 2019-04-25T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD29YYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpjOUP+weTx8S2FDY8NqtwpVGuxL+e
  8SbyHXZKiG8O0aJjJ1vzRsfn3PpVt9DGYy3watCnFMH8xElxKrESIPnWZGCJX+/i
  BQ1mSZZrzahkpbBoF6hdPBy+fhDQUz/tIn6VAVIAYFibAYAErC893p2UR2NX4Clt
  YgY68PWv/O6ez1giVPVYGmZJXvccOGqJ4smvtiz9eYXO+eeNKMFgS4BVhyWi6PwX
  6OdfcTTud0/NMRsJL/GbZshS3qnSRxDjQbAsaLMm/7yOWT4m2xRz0CHKiQuFHuKT
  JA34abr/3tLxJxDvhToKr48HeE5o0dCXo7UWVMLDFcTv0yuBE1thyNn2e6oohO4n
  6oY6LjzT3yjdxbYm+uNj1H0XbzUVE6vdTXYgod3cN71clvI0CJrSqG8eYc8b44xf
  /kqnw23Ra0DFK18K1YjeOD5Qg+71+WfGJ/BOurLFXybsjKofT7cytb3wxXOkdcF7
  1AIGACaKxcMfDWkZ4nJm5IFe5tGJkWVX45eZdIO4KlYS8Xs0SzPb5fSSEH5ECLVV
  n9XmtwzF/hDY1Q7NnbRbr6GZ+6+/8qDir1S2yNL3VVsYyXvcNUYhere0BPJ/K7av
  fSoIcwNFrltjl3dVbvObNRqWcgsAkjo+yE95wsWq8po2oXTr2V6HVE9QghrqqTEY
  OFRqptvFoqdBMg5sVMD2
  =JnEI
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
