---
date: 2019-04-25T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD6MUYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpzOkP/iJa1/FFnTHzsmI+CILcqbfj
  JDhN/zfxqUQtAFeeWEcJnV4yp8dtqUh+gsURekIdXxMJk4MDibF6KXdtZQWa4w0S
  9t0G56qaNNEyDr16Ulz0FTjrny++RAOLqj+NkIZMhsqJPApNsToHrNQsRY+ZEFb8
  peWH5PasgAHHjMpfwkYbMUy6zi5DVlnr9eRw5a20KifoHXSeR9U3BeL7eibLRTy4
  DtGqrpmO0Rgt7rZXGQ8lUsiLQ9s2KlNuIeesGKAbC3cqS43st4Q3kWzYyRPytrZU
  yoWGeMIL5RbBhQ5u8XElan9td88PqKgGYhWt8esl+OA+CDwEiPOEdz0m+l4clKf9
  BCGJ5LNWLaB83T2SL7F01mjwzQJqyMrCLvAKiDb/IEkyljTZjBklccl/htKk4odK
  DmCZ5gaFcIaHHVynvHLMHutabA6CYZ9/s4i1hq7guhMBVORz+WA+HqJAfdVSRSPJ
  Rv7jFA0AYpc+Ly570qczaE6rF+UB3zKzCKPdkj0PO90NmCVOU/ZHz6MCL5g5/kIH
  85cMxf9v7S1j9Zr1+iNgg3S1uMh64HHkUu2sqk38EdqdAzbVRLF2xnuurLLIHuiN
  yqvoXWGFQGJKGGNkDWQjATTB54pqx4bM+waLXmnTCTaQqONiQNx/xF3+9XegKf3H
  P3CaTSpS4+Mg3uh9iHIB
  =k/xm
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
