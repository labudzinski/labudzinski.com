---
date: 2019-04-25T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD6eIYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIp2iUP/2W4ZHcGek4+F5wYMNhY3DNF
  6J/kbdBOLi/EIUgxH7UJRC3lGIIWAB8JqMIJuabmZ6ScWABLrPgtJEavPNPlUvc6
  Gd/HA1embMpj+7P7iawORtJzp+Fgx7JF0r6HMfYZ4M+hxwyU1uqAYLfmV/2SZbGc
  ade+6Zx9L8OGYG9WgmsJyssfqRnPBp+UR561meVNKbL1Jk3RrdKQzrlksfEoPrVl
  AJlvYLxQW6YovRxEAe6fXTAV0j14gOLj/07/JkmLSWowMU1X98ZNo/5jgwqqdOqm
  aCNRpjSXznBXVFzDZBI+lL3ZAVpxuP+Ry8xLoGpezihx0YVaU37XnjCdQRFn8Cb7
  jSY06BlGL5oVCv46P0bdDGykvAHDWJbuoaUI4yxXBsiAH1HDk9Xl4ovk7a1QjGO6
  zvtAIHp4nDtCFgZ5h9J+gBS8+UUTtkPR+/7n+1y4OwuA0oMc/sczFvnafaHgXDcn
  Moq+b9aj0BOII2f3l9Cs2n5UL9IKxTn+kfoks+UzeZhPd7w8ny/woL1DxqzCoJ0I
  Bzn7pIUxCPLAtuOpqVlskF2J76DCqHX3YDmoZ1P/534/qEsXEhFYNXxcIdIjnP6h
  JNJC5xdIL8goh3/5kabKew/hYMeGe1ChRkLkqUs9fbBrBbCRUWE3mNHQtDGR27/2
  giA9wb4rapf3+odwXKHR
  =7QU0
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
