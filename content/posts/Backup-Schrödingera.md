---
date: 2019-04-25T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpEzuoYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpF2kP/joyRRn3U8AIXNM4w0+F/Eae
  bYYwL31hO6HCZSeJkEul8jIg93KVKcHFmuiudO2J9GApLV1OhfbJT0sKW6Qnsmqh
  Ej55oVAg8B8xjtKvv+RdPA0ocA+uiT05LcZU12ZyAi6Y1t72ZOkwILVkzSNu9Zj/
  JHRaXC10vzg19vq8hI0Gxs9EoPM0b8mFC9JGVJ796K/hOdlbK4ZiPJZnhR0N64l/
  Ccamb2+AG5GIEWAeGxsrUcuI0Uk92DAga3MxQrc1FSi7fRnUqV5vbG3jisUOMlbz
  qnmCS3GoHqRuEkSHpK8bEcAYPx0upvtx3BkmYJVJHk/dKtsj56wX5emUmeIRBBTH
  6JbtRoXy+6WP5AhzVbPxlQC5N4hHDOYq90SpcsrXLT67wnVPtW9RNL1hXrb20hOj
  uQI8CJwjhjPmnQNsUXE8eAlm/50wvHB1CrN3q5yLgNagTIQNIVs0W79IPBngmzKR
  nSSu9dEhrIzBOaJgXp8yRBX5NeDyN8dAwUec5VDAB9H5JXkCIYp4Bk6s6fu4b2Z1
  6bQ+j4TrI0JiTqPTmV1OwoKjOBiLPTfH4WP57UYnBZ24Q68Aeb6mS/yVObzviW/r
  McAujmUbhhw/J1GyUTX+I40O/n8/t4m9K1VkbuYUXXNLE8NVL3K2RJh/i5Ekyd0F
  MPnW+jBOJ2s8w3jfKezQ
  =AdPB
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
