---
date: 2026-07-07T13:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQIzBAABCAAdFiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpM2+sACgkQXN6UakFg
  wiluvg//cmRjwLsw3eSinvDoxGfV4GURTz2meUzhBE8hVQKWND5VvuY/gnXFClpu
  zVAfts7G0lJy/Grz9DdFOT0Cn+dol0lwrLmSwWPLGv9xSILhkkJAkIzB9FUq6Rws
  k/cDI47O4NCIbpLWgyuaR8iLcsMZrfdnH7OsMHgoOJHAUZlCJsC9RzMz269C+aOx
  hP5WUZ/ZWiUZOx9gPvnWOuWeyU5lFb2iQXJ5QB/Vlhu1b5XKL6nGwLSGm/ptl/hY
  MbEuNw5+o66x4TksA/SXAOd0WGsEEdUuMjZdpic2TtHS990x8oQLaIw60CExrHxY
  dcjEeZCQV8l4y5KC2DhEHGSOubMZqLcaHkh8nEcaa66Q8KTdPjRe9ghLRvdK3jN3
  M2fSWIu6P/jL50GI0Wo/f4iCvx5ltJmSE3YrYH2+nihNkwZlTDEflOUN23hK9VxQ
  bPKrztOzJFDeB709rS21cAeC0JwuIlYaok9DAku3Njh0EnsvJPlOK+IYueVD9Wr3
  j7V+8KL9KMPBB5ahpJ/5obepn/MzQK38XH2smEIdXF1zdk3exRQxFVG+Tx4g+r0U
  dxrz2gMG9o//4jvfcNqBM4NYQk1ARbaU7Gfs3wqvSskr3RYsuMgcwcG6o/FB8joQ
  aRQHygL1ICjMOuGZMHc64EMrdnXLAxkvwilg/I3z9XjVXqFtZhs=
  =bxFO
  -----END PGP SIGNATURE-----
tags:
  - HTTP
  - API
  - REST
title: QUERY w HTTP. Rozmyślania po lekturze RFC 10008
toc: false
---

Dziś mój znajomy przysłał mi link do [RFC 10008](https://www.rfc-editor.org/rfc/rfc10008.html), specyfikacji zatytułowanej *The HTTP QUERY Method*. Przeczytałem dokument i od kilku godzin wracam do niego myślami. Na pierwszy rzut oka rozwiązanie wydaje się rozsądne i eleganckie. Wprowadza metodę, która łączy zalety GET, czyli bezpieczeństwo, idempotencję i możliwość buforowania, z możliwością przesłania większej ilości danych w treści żądania, podobnie jak w przypadku POST.

Im dłużej analizuję to podejście, tym silniejsze mam przekonanie, że nie jest to rozwiązanie typowe dla inżyniera dbającego o architekturę systemu. Raczej wygląda na wygodne obejście stworzone z myślą o programiście, który chce przesłać tysiące parametrów bez głębszej refleksji nad konsekwencjami.

Problem jest dobrze znany każdemu, kto projektował interfejsy z zaawansowanym wyszukiwaniem. Filtry stają się coraz bardziej rozbudowane, ciąg zapytania rośnie ponad miarę, pośrednicy sieci tną adresy, logi tracą czytelność, a mechanizmy buforowania przestają działać efektywnie. Do tej pory wielu programistów radziło sobie, stosując metodę POST na punktach końcowych odpowiedzialnych za wyszukiwanie. Rozwiązanie technicznie działa, ale semantycznie jest nieczyste, ponieważ POST nie gwarantuje idempotencji ani bezpieczeństwa operacji.

Metoda QUERY ma to zmienić. Przykładowe żądanie wygląda następująco:

```http
QUERY /search HTTP/1.1
Host: api.example.com
Content-Type: application/json

{
  "filters": { "... bardzo rozbudowany obiekt ..." },
  "include": ["id", "name", "status"],
  "sort": ["-created_at"],
  "limit": 500
}
```

Technicznie jest to czyste. Semantycznie również wydaje się poprawne. Mimo to coś w tym rozwiązaniu budzi mój niepokój.

Największe zastrzeżenie dotyczy tego, że standard zamiast wymuszać optymalizację, legalizuje jej brak. Zamiast skłaniać do pytania, dlaczego filtr zawiera setki pól i głębokie zagnieżdżenia, daje narzędzie, które pozwala wrzucić wszystko do treści żądania i oznaczyć je jako QUERY. To podejście charakterystyczne dla osoby piszącej kod doraźnie, a nie dla inżyniera projektującego system.

Inżynier w takiej sytuacji zacząłby od analizy przyczyn. Czy da się uprościć kontrakt interfejsu? Czy warto wprowadzić dedykowane widoki danych? Czy GraphQL lepiej obsłuży potrzebną elastyczność? Czy po stronie serwera można zoptymalizować indeksy i widoki materializowane?

QUERY pozwala uniknąć tych pytań. Wystarczy przesłać więcej danych. I jeszcze więcej.

Wady tego podejścia są zauważalne. Ukrywa złożoność zapytania, ponieważ szczegóły znikają z adresu URL, co utrudnia debugowanie i analizę logów. Zachęca do dalszego rozrastania się interfejsów. Wprowadza też wyzwania dla infrastruktury, ponieważ serwery proxy, mechanizmy ochrony aplikacji, sieci dystrybucji treści i narzędzia monitoringu nie zawsze są w pełni przygotowane na nową metodę. W efekcie maskuje problemy projektowe zamiast je rozwiązywać.

Oczywiście istnieją konteksty, w których QUERY jest wartościowe. Bardzo złożone raporty, zaawansowane wyszukiwarki korporacyjne i systemy wymagające wysokiej dynamiczności zapytań mogą skorzystać na tym rozwiązaniu. W takich przypadkach metoda oferuje czystsze semantycznie podejście niż nadużywanie POST.

W typowych aplikacjach internetowych i interfejsach REST jednak odbieram to jako krok w niewłaściwym kierunku.

Po lekturze RFC 10008 utwierdziłem się w przekonaniu, że najlepsze interfejsy, z jakimi pracowałem, nie wymagały takiej metody. Od początku projektowano je z poszanowaniem ograniczeń i z dbałością o konsekwencje. Nie chodziło o to, jak przesłać jak najwięcej danych, lecz o to, jak przesłać ich jak najmniej, ale wystarczająco dużo, aby spełnić rzeczywiste potrzeby.

QUERY jest wygodne. Pytanie tylko, czy zawsze powinniśmy wybierać rozwiązania, które są najwygodniejsze.
