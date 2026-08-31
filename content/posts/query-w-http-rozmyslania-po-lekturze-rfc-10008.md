---
date: 2026-07-07T13:00:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQIzBAABCAAdFiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpM3P8ACgkQXN6UakFg
  wikZdRAArhLSd40KGMblAYh058cVqsEZdzD1toLk7uCJ5N8J5+ul8KTMX/GE6kHo
  JcGZTaYnwRqjJKVqaiaTCvzWTlQcnd8sps+YikDGOhvSzP+jkIg4xwFLsJRM0VHH
  WMpnIURmtP/FSTlUubZTI1o88c/JZXPcRLKongYXRot9yxQnGKt9JVDJkXdwS0Zt
  1qzdimoREhobvSFILdokHrHgn4kiPkh3WSgQWbVgkfq5TwtgbFaXOnAq+0hLTLX1
  FulmX4jXC8zrH49mUd+MU0DvS5UDMQg41AGz+HpkKBg4NC1gsmoUp3Gc3Mn+xpMs
  vJBDupZeFcXi4Q4FIj9FcZjiWjDeN42ZDSCv/xWbz7M14yW51y8iGAwlCTzgY51P
  QCZKHsRfhhc+qvF1Kl77xSJmbdPYvaDlivwWyJfqoD5Iuqbml8cyCaaFAnFuAUmM
  7wQKh7g6c1vHGezfmCdaYAFTNAK25I5TXwgMJljyZOfR0LE1GI8tu+bp93rWNWds
  BNvV0GNiJ5jRqQ8rmRKeqWEwAJJQXHL4pYI/wyxRjwUzvzmQvgpoTEnGeiqzR5lZ
  jS7LGrO2aoD3fC48G5dKIdBZfllG/vlH4ThdPNtWbrS8XVafr6I0sWTrISQGC9w2
  H4Qu8v5fqfRpGti++7096TZ9cBfwQZklW9P8qOYrjKe3xxoD/JQ=
  =i8ww
  -----END PGP SIGNATURE-----
tags:
  - HTTP
  - API
  - REST
title: QUERY w HTTP. Rozmyślania po lekturze RFC 10008
toc: false
description: "RFC 10008 i metoda QUERY w HTTP. Idempotentny GET z ciałem POST wygląda wygodnie, ale psuje model zasobów."
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

Oczywiście istnieją konteksty, w których QUERY jest wartościowe. Bardzo złożone raporty, zaawansowane wyszukiwarki korporacyjne i systemy wymagające wysokiej dynamiczności zapytań mogą skorzystać na tym rozwiązaniu. W takich przypadkach metoda oferuje czystsze semantycznie podejście niż nadużywanie POST, natomiast nie wymusza optymalizacji jak przy GET.

W typowych aplikacjach internetowych i interfejsach REST jednak odbieram to jako krok w niewłaściwym kierunku.

Po lekturze RFC 10008 utwierdziłem się w przekonaniu, że najlepsze interfejsy, z jakimi pracowałem, nie wymagały takiej metody. Od początku projektowano je z poszanowaniem ograniczeń i z dbałością o konsekwencje. Nie chodziło o to, jak przesłać jak najwięcej danych, lecz o to, jak przesłać ich jak najmniej, ale wystarczająco dużo, aby spełnić rzeczywiste potrzeby.

QUERY jest wygodne. Pytanie tylko, czy zawsze powinniśmy wybierać rozwiązania, które są najwygodniejsze.
