+++
title = "DataLeakTracker"
description = "Publiczny rejestr wycieków i ujawnienia danych: karty incydentów ze źródłami, weryfikacją i API metadanych."
date = 2026-10-05T19:10:00+02:00
draft = false
homepage = "https://dataleaktracker.org/pl/"
github = "https://github.com/labudzinski/dataleaktracker.org"
tags = ["OSINT", "bezpieczeństwo", "wycieki"]
toc = false
+++

DataLeakTracker to niezależny, publiczny rejestr wycieków i incydentów ujawnienia danych. Komunikaty organizacji, publiczne rejestry i publikacje medialne trafiają do jednej karty incydentu. Korzystają z niego badacze, dziennikarze, specjaliści cyberbezpieczeństwa i analitycy OSINT.

Karta ma trwały identyfikator, odnośniki do źródeł i historię weryfikacji. Widać z niej, które organizacje zostały dotknięte, jaki zakres danych opisują źródła i jak przebiegała chronologia. Wpisy rozróżniają potwierdzone fakty, szacunki i twierdzenia, których nikt jeszcze nie sprawdził. Status weryfikacji dotyczy opisanego zakresu informacji i nie rozstrzyga odpowiedzialności prawnej.

Rejestr da się filtrować po kraju, sektorze, statusie weryfikacji i poziomie ryzyka. Są też profile organizacji, statystyki i oś czasu. Interfejs jest po polsku, niemiecku, angielsku i czesku. Źródło albo korektę zgłasza się formularzem. Kontakt jest opcjonalny.

Serwis nie publikuje wykradzionych baz, haseł ani danych uwierzytelniających. Działa jako publiczna beta.

Metadane opublikowanych incydentów są dostępne przez publiczne API i eksport CSV. Można ich użyć jako kontekstu w analizie zagrożeń albo zaimportować do własnego SIEM. To nie jest strumień alertów na żywo. Trzeba samemu zmapować pola.
