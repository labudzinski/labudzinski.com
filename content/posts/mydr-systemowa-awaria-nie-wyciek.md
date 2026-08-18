---
date: 2026-08-18T20:13:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQIzBAABCAAdFiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmqEqiEACgkQXN6UakFg
  win80Q/+Pnk81IIbeMOnWZ+R5nAWQMk63zzZa3mT/wBLSHmSC7WOhjCBVfqZaQME
  N+bgIFwX2nLg0uEbZAhgbTqOgJ8l6DxM56UxEYrGvr9idjplK20Gk/VPkUU54fio
  6lU1DLZQVrwU1nWleuHr+mJjHX7fAXJZQwLDYD0HM5xNRUrEamkO0+jxWNcKFBHE
  It59hLYgC/+mtLqaNQ8oEaQwns4FclA2Pj1QXKFpGmkDI05keZKAbI+d3WTdwCvJ
  fI7UzPuq7swlWO9NrAbNBRChK4cfO7Kw3yIPEdMC8/1lamgQoZVyL3cF3gumjMaK
  K89ImffGMtmA9EcwGWIEf5eZF00iGS4kpsyaDGqfKsCYOOjDgxBUmnbwAs6pTQTi
  yqkLRz8ktJKTGhtu9fEk3XT/flAGH6TYZBxrgMT43CP3Ng2S0JXGsFnu8fkBxCQ8
  fO2/B6zrhEIXeC0LHjUPmMka+9gsOyD+7i0D2LoiuKFMRDfgdihYb15Dn15Dk8tR
  YjzWUJMapmRF1HhsZvRJem6p7wspfHR0s9drXUEEKhl3tk8jVL429mXCK6NGs5Hx
  fMIqk2YLMKh2cIMg/be1MNBjuPWzHEY9XK3QOvxzzL8BuEvzxZR5b8IgbB8PHxzG
  lAizRNEqzQi4PAFKeszP2WNifadhpyLryotzNNPiqnkruJ1K8uo=
  =tprJ
  -----END PGP SIGNATURE-----
tags:
  - bezpieczeństwo
  - RODO
  - MyDr
  - PESEL
  - dane osobowe
title: MyDr. Systemowa awaria, nie wyciek
toc: false
---

Kiedy dziś patrzę na to, co wyszło z MyDr, przechodzi mnie dokładnie ten sam dreszcz, co wtedy, gdy otwieram raporty z testów. Tylko że tym razem skala jest zupełnie inna. Prawie 19 milionów rekordów. Ponad 2 terabajty, w tym dane medyczne i osobowe, które z łatwością mogły trafić w ręce osób, które zapewnie wiedzą jak je spieniężyć. To już nie jest tylko zwykły wyciek danych.
To bezprecedensowa akcja, która pokazuje, jak krucha jest w rzeczywistości cała struktura, którą od lat budujemy wokół ochrony danych.

Przyczyny sięgają znacznie głębiej niż pojedyncza luka w kodzie. Z tego, co sprawcy opisali pracownikom zaufanej trzeciej strony, łańcuch zdarzeń wyglądał jak klasyczny atak, a jego prostota była niemal bolesna. XXE w obsłudze certyfikatów PKCS#12, zdalne wykonanie kodu, a następnie klucz API serwisu GitHub, kod źródłowy i infrastruktura w AWS. 
Każdy kto kiedykolwiek miał do czynienia z prawdziwym incydentem bezpieczeństwa, zna ten schemat na pamięć. Pytanie nie brzmi, czy ktokolwiek mógł to przewidzieć. Pytanie brzmi: jak to możliwe, że firma przetwarzająca dane z 12 000 placówek dopuściła do sytuacji, w której taki scenariusz jest choćby w najmniejszym stopniu prawdopodobny. Jeśli klucze szyfrujące były przechowywane w tym samym środowisku co same dane, jeśli Jira lub inne systemy wewnętrzne zawierały informacje, które umożliwiły przeprowadzenie naruszenia, a recepty i dokumentacja medyczna nie były zaszyfrowane tak, jak powinny, to nie jest to już kwestia pecha. Mówimy tu o fundamentalnej wadzie w podejściu do architektury systemu żeby nie powiedzieć o jawnym pogwałceniu fundamentalnych zasad. A teraz pytanie, które niepokoi mnie najbardziej: czy firma MyDr w ogóle posiada certyfikat ISO 27001? Bo jeśli tak, to na jakiej podstawie go utrzymuje, skoro widać wyraźnie, że podstawowe wymagania dotyczące podziału obowiązków i ochrony kluczy najwyraźniej nie zostały spełnione, a fakt, że taki wyciek byłby możliwy nawet przy posiadaniu ISO 27001 jest nieporozumieniem. Czym innym jest akceptowanie ryzyka, a czym innym jest jego ignorowanie.
W normalnej organizacji taki incydent oznaczałby natychmiastowe zwolnienia, audyt zewnętrzny i całkowitą zmianę strategii. W tym przypadku słyszymy oświadczenia o „celowej działalności przestępczej” oraz zapewnienia, że systemy są teraz bezpieczne. To nie wystarczy. To zdecydowanie za mało.

Zbyt dobrze wiem, co może się teraz wydarzyć, na tyle dobrze, że nikt nie może spać spokojnie. Numery PESEL, imiona, nazwiska, numery telefonów, adresy e-mail, treść notatek z wizyt, informacje o lekach, receptach i schorzeniach. Dane historyczne, głównie do kwietnia 2024 r., ale to nie sprawia, że są one mniej niebezpieczne. Historia leczenia obejmująca kilka lat pozwala na stworzenie bardziej szczegółowego profilu osoby niż ten, który może zapewnić większość banków. Na podstawie takiej bazy danych można ustalić, kto jest w trakcie leczenia psychiatrycznego, kto cierpi na choroby przewlekłe i kto przyjmuje konkretne leki na receptę.
W nieodpowiednich rękach wszystko to staje się gotowym narzędziem. Dla przeciętnego człowieka stanowi to konkretne, namacalne zagrożenie. Ktoś może próbować zaciągnąć kredyt, wykorzystując dane. Ktoś może stworzyć fałszywe dokumenty. Ktoś może przeprowadzić kampanię phishingową w oparciu o wiedzę na temat leków, które dana osoba faktycznie przyjmuje. Albo po prostu sprzeda cały pakiet na zamkniętym forum, gdzie nikt nie pyta o jego pochodzenie. A konsekwencje nie kończą się na pojedynczym oszustwie. Ciągną się one latami, ponieważ raz skradziona tożsamość medyczna nie znika sama z siebie.

W tej historii RODO jawi się jako papierowy tygrys, który po prostu się załamał, gdy nadszedł prawdziwy kryzys. MyDr nie było administratorem danych. Pełniło rolę podmiotu przetwarzającego dane. Dane były przechowywane na jego serwerach, przepływały przez jego usługi i znajdowały się w jego systemach, jednak formalna odpowiedzialność spoczywała na tysiącach klinik i gabinetów lekarskich, z których większość nie posiada nawet własnego działu IT.
Kliniki te nie wiedzą dokładnie, jakie dane wyciekły. Nie wiedzą, ile rekordów dotyczy ich pacjentów. Nie wiedzą, które pola danych zostały wyeksportowane. Brakuje im narzędzi do zweryfikowania tego. Brakuje im wiedzy specjalistycznej niezbędnej do przeprowadzenia właściwego dochodzenia. RODO zakłada, że administrator jest w stanie wypełnić wszystkie obowiązki informacyjne, a w rzeczywistości mamy do czynienia z sytuacją, w której podmiot przetwarzający traci kontrolę nad ogromnym zbiorem danych, a administratorzy dowiadują się o wszystkim z mediów i komunikatów rządowych.
Prawo, które miało chronić obywateli, pokazało w praktyce, że zostało napisane z myślą o idealnych warunkach laboratoryjnych, a nie o rzeczywistym świecie, w którym pojedynczy dostawca oprogramowania gromadzi dane dotyczące prawie połowy populacji.

Zastrzeżenie numeru PESEL to kolejna wielka iluzja, którą sprzedaje się ludziom jako rozwiązanie gwarantujące bezpieczeństwo. Włączasz tę opcję w aplikacji mObywatel i masz wrażenie, że faktycznie coś zrobiłeś. Tymczasem zablokowanie PESEL nie zmienia faktu, że nadal istnieje, nadal jest niepowtarzalny i nadal może być wykorzystywany w sytuacjach, których system nie blokuje automatycznie. Nadal możliwe jest próby otwierania kont, tworzenia dokumentów, ubiegania się o świadczenia lub przeprowadzania ataków socjotechnicznych.
Architektura PESEL jest zasadniczo wadliwa. Numer przypisywany raz w życiu, niezmienny i powiązany z ogromną liczbą procesów administracyjnych i komercyjnych. Gdy numery ponad połowy obywateli są narażone na niebezpieczeństwo, rząd nie dysponuje mechanizmem umożliwiającym masową wymianę tych numerów. Nie ma procedury, która pozwoliłaby obywatelom uzyskać nowy identyfikator i unieważnić stary w sposób, który rzeczywiście odciąłby stare wektory ataku. Zamiast tego otrzymujemy zapewnienia, że to ma nas chronić, a to tylko fikcja. 

W tej sprawie rząd zachowuje się jak organizacja, która za wszelką cenę odmawia wzięcia odpowiedzialności za system, który sama stworzyła. PESEL to identyfikator nadany przez państwo. Za jego bezpieczeństwo odpowiada państwo. Gdy wyciek danych osiąga skalę, o której tu mowa, państwo powinno dysponować gotowym, dobrze przećwiczonym planem działania. Tymczasowe blokady. Przyspieszona wymiana numerów dla osób najbardziej narażonych. Prawdziwe wsparcie dla ofiar, a zamiast tego mamy serię oświadczeń, zapewnień, że dane nie krążą publicznie, oraz prośby, by ludzie sami się chronili. To po prostu przerzucanie odpowiedzialności za swoją niekompetencję na obywateli, którzy nie mają narzędzi do samodzielnej oceny ryzyka. W normalnych firmach, gdy dochodzi do wycieku danych klientów, bierze ona na siebie pełną odpowiedzialność. Informuje swoich klientów i zapewnia wsparcie. W tym przypadku rząd chowa się za formalnym podziałem ról i twierdzi, że MyDr jest prywatną firmą.

Wreszcie pojawia się kwestia samej firmy i jej systemów. Organizacja zarządzająca danymi medycznymi milionów ludzi powinna mieć bezpieczeństwo zakodowane w swoim DNA. Jeśli wrażliwe dane nie były odpowiednio zaszyfrowane, jeśli klucze były dostępne w tym samym środowisku, jeśli architektura pozwalała na taki ciąg ataków, to mamy do czynienia z zaniedbaniem wykraczającym daleko poza zwykłą pomyłkę. Certyfikaty ISO nie są magią. Są tak dobre, jak procesy, które za nimi stoją. Jeśli po incydencie firma nie jest w stanie szybko i dokładnie określić, jakie dane wyciekły, oznacza to, że rejestrowanie i monitorowanie nie funkcjonowały na poziomie wymaganym w tego typu działalności. Tymczasem władze chowają się za oświadczeniami, które brzmią, jakby zostały zaczerpnięte prosto z podręcznika kryzysowego PR-u. „Zajmujemy się tą sprawą”. „Monitorujemy dark web”. „Dane nie są na sprzedaż”. To wszystko może nawet być prawdą, jednak dla osoby, której numer PESEL i historia medyczna mogły już trafić w niepowołane ręce, te oświadczenia nie mają absolutnie żadnego znaczenia.
Potrzebujemy konkretów, jasnego harmonogramu działań i prawdziwej pomocy, a nie tylko kolejnego oświadczenia, którego głównym celem jest uspokojenie mediów i inwestorów.

To, co się wydarzyło, nie jest odosobnionym przypadkiem. To symptom, który pokazuje, że skupienie wrażliwych danych u jednego dostawcy, w połączeniu z niejasnością odpowiedzialności wynikającą z RODO, stwarza idealne warunki do katastrofy. Pokazuje, że system identyfikacji obywateli oparty na unikalnym numerze nie jest w stanie sprostać rzeczywistości, w której dochodzi do naruszeń bezpieczeństwa danych na dużą skalę.
Pokazuje, że państwo wciąż nie nauczyło się brać odpowiedzialności za infrastrukturę, którą samo narzuciło. Dla przeciętnego obywatela oznacza to, że musi żyć z ryzykiem, którego nie jest w stanie samodzielnie zneutralizować. Może zabezpieczyć swój numer PESEL, może sprawdzać swoje konta, może uważać na podejrzane wiadomości. Nie może jednak zmienić faktu, że jego dane medyczne mogły już zostać skopiowane, a system mający je chronić zawiódł na każdym możliwym poziomie i zawsze musi już patrzeć za siebie, czy aby przypadkiem ktoś po latach nie będzie chciał użyć tych danych do oszustwa.

W firmach technologicznych wyciągamy wnioski z takich historii. Zmieniamy architekturę, wdrażamy model bezpieczeństwa „zero-trust”, rozdzielamy klucze, przeprowadzamy audyty dostawców. W Polsce, po incydencie z MyDr, prawdopodobnie zobaczymy więcej komunikatów prasowych, być może kontrolę przeprowadzoną przez Urząd Ochrony Danych Osobowych (UODO) oraz powrót do normalnego trybu działania. I tyak aż do kolejnego wycieku danych. Ponieważ problemy systemowe rzadko rozwiązują się same, a przepisy na papierze nigdy nie chronią przed ludźmi, którzy wiedzą, jak znaleźć ten jeden, jedyny słaby punkt.

I patrząc na firmy zajmujące się szkoleniami z bezpieczeństwa, z RODO czy innych tematów na pograniczu technologii i prawa możemy odnieść mylne wrażenie, że wszyscy są dobrzy. Już zapomnieliśmy, że polska scena underground, która była jedną z silniejszych na świecie, nie zniknęła, tylko przeniosła się w bardziej ustronne miejsce. 
