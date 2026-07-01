---
date: 2019-04-25T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD7DAYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpAcgP/0bETVM2yE468nhxMAla68dV
  aNNrgaNT4sUSTAhUADfXyjDyWsNbHpL94/oQWC0C/NybIeZiwqVram5n5fxEnFC5
  WwCvGbYEBeOyXBWAiI4K0TEXdyuj2VN9QSv1VXEog0K6A7WgNfq4alu23TuWAXKn
  v1b+ICP306kY0RaKs4WLCcXDSzVsauwX5BxQQCTACw9sCX0c7O9CWLRhfa0lerf9
  td8dB6/SACICugfUuwqaH4aU4rmDZv65ehjK58HFoa1o14sJTQtxzBG8x+fSipgD
  tdFBg4g29QZrnW2UK/Utz42gzUZ6qOV4T9AQ1IapKV7WeK6xlLq/UrRRpyJgnTL3
  HpVKKhsIwsyu9wv5JouhvyAjNcvlVLzGjkCMZJYZj7nh6UXkJNN92JDvSUbVyz3G
  O7oHSlazmdjwg3nQ/onwwyypXFy1VBs4MKfjcL9FwhsrF8/MxaHZOUE9/Q9xck5G
  b0kruXcAf/cQdybO0+j31q4cer+/sqG8JCIdT+THK2YFM4jJqYifGiTEkzvD4hAU
  cklyiy/MZz1EzW8BmFKDeWF2yPs96ExTY0hR2mAMV2Q2ALAh4GsRYiXNfEzkCFee
  NcnXHSVQiBTMFqaDJOKpa/dv81sB3IkoVDSRiCx2YjXQetathRyMK8ho+9KWy3A+
  GYRbXmll9oADe5OIJkZI
  =c31E
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
