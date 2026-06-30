---
date: 2019-04-25T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD2RoYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIperEP/3NfH3jHOYi61ZTq63GTECIf
  eQix+pLqHMS54GmQen5TYgIHH2q9NWn/JpXkBn117sv7bmTcQh5oqAJC02b8z3z/
  Wwn3FQ21rLmnvOKT3XGNRHBmF18Rg7mO5zUNu1kshEkE217Y7Sg8nNCDoCOf7rzc
  XSj4XfjAwO4bzo/q28JuiOPdY9an2dfEFWQddtYR/rbswhHc3E+Z3f9Jr2thG7SS
  hFcmMt7PfHLyAiOO2bcgt3Q/rPPyrzeGONBR02LBkzPsH3KeBquvux1x+XWwlnoo
  tL/RbNmjsLKRx5V+R4kcsJY4auUCkrbeeiA5dlFRCqZr/2RCepu598oyFzeobpo7
  4gDCQZpO9eRnHBONzalFbVNoVjjG9mqnfXkXPbzP956frnB+cdQy1GFCpAIaEU/o
  A5BkZjYzT7JoG4ZtH8/bm2vJOYwpk1/U2CluthD5Dnb4vzOkH/ihBAbKa3bdqAG3
  ENodCM5G5ppKsNwdhr3/GYp4NMwdTviG9ebVYt3Ppiptvf3S/Ze/YSr/RpQPFgmI
  q63nzhrVtTl5h4Rhk0CHtQiDK8Z1gNV+V2YJUeEgsehBo4rKsAXfzQUlE0I2sVQK
  59CFnrXQWqVXolBCN9+/FQ3xnQ84jlCI6e4WpW9vb4VwH/VvuI0pHjlOiboaWFpy
  YnrrXWXN1j2PwEhFiLnW
  =/s/J
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
