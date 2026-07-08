+++
title = "BatchRequest"
description = "Biblioteka PHP do wielu wywołań API w jednym żądaniu HTTP dla Symfony i Laravel. Operacje równoległe, scalona odpowiedź JSON, obsługa błędów RFC 7807."
date = 2026-07-08T00:00:00+02:00
draft = false
github = "https://github.com/Lemric/BatchRequest"
composer = "lemric/batch-request"
license = "GPL-3.0"
tags = ["PHP", "Symfony", "Laravel", "API"]
toc = false
+++

BatchRequest to biblioteka PHP do obsługi wielu wywołań API w jednym żądaniu HTTP. Powstała dla aplikacji Symfony oraz Laravel, w których liczba połączeń po stronie klienta wpływa na wydajność i przewidywalność działania.

Zamiast wysyłać serię osobnych żądań klient przekazuje jedną partię wsadową na punkt końcowy, zwykle metodą POST na adres `/batch`. Każda operacja w kopercie JSON zawiera metodę HTTP, adres względny oraz opcjonalne ciało. Operacje niezależne wykonują się równolegle, a zależne sekwencyjnie. Serwer zwraca tablicę odpowiedzi w tej samej kolejności, co ułatwia powiązanie wyniku z konkretnym wywołaniem i decyzję o ponowieniu nieudanych operacji.

Biblioteka współpracuje z `symfony/rate-limiter` i traktuje każde wywołanie w partii jako osobny ruch przy liczeniu limitów API. Błąd pojedynczej operacji nie zatrzymuje pozostałych. Nieudane odpowiedzi podrzędne mają postać dokumentów RFC 7807 z typem `application/problem+json`. Obsługiwane są też mieszane typy treści, w tym JSON, HTML, dane binarne w base64 oraz odpowiedzi bez ciała. W jednym wywołaniu można przesłać także załączniki binarne w formacie multipart.

W Symfony dostępna jest integracja z profilerem pokazująca liczbę żądań podrzędnych, czas wykonania, zużycie pamięci i szczegóły transakcji. Pakiet jest zgodny z Symfony od wersji 6.4 do 8.x. W Laravel wystarczy zarejestrować dostawcę usług i ustawić limit wielkości partii w konfiguracji.

BatchRequest publikowany jest na licencji GPL 3.0. Instalacja: `composer require lemric/batch-request`. Kod, dokumentacja i przykłady znajdują się w repozytorium GitHub.
