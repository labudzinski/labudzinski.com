---
date: 2026-08-26T15:10:00+02:00
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJoBAABCABSFiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmqO5gMbFIAAAAAABAAO
  bWFudTIsMi41KzEuMTIsMCwzGBxkb21pbmlrQGxhYnVkemluc2tpLmNvbQAKCRBc
  3pRqQWDCKYccEACGFxElraxeESp/iHMtAreuhZt7HIwB8iyR39PIAlaJ+d5ORqVG
  /6ZRQkJx4MIZhzPFEnQWgWdglzgKSLBpCFCqIsQXZIZk8Sqx0JjbCc1AdJWaYXQv
  hP4eKkvTOKVNXhhX1gspDVjc+1mlusIlyOnAimLZHDuo7nwbfNQEDkYL0fFxREf8
  EOmx+PGqklVA9zahUgKS3WlrdE4xLBwrC4im9CowyS6g0U4uqUc6kIen/nSivAS8
  yxNanIPyC9A4nfzNwzKqnuGPXJiEEz+BnazCh644L3lnKp7d/pWQZb9IQEkSQaKH
  Vp+cEQ8YUXMIDf5l0pDPsrIdpz5KfVgxA4SEjWjwnG3Cyzwavdcj/JoqO7CELFZh
  11yV6CwDmCV2ZkIQ5NvcM98WkxCjM4RdGneS8nULRoBwu3V086b1uB0gbJP8xR7N
  vnds9Iy1eNcZiEUyL19uvL8El4fQQSPfyT0e5KbGI0xXoM1Hig0pJNkAtWYWwaA6
  YkZA58h9Q9GAvrQhN0SRoT8JgEFfDUklBcfow2Ky4fPj7U2EHxhdevXiScLGlxxq
  vUkA1qVWMklzuOu56I/2QqSPVprt/2DH/pQt5SHH69zCvc/cerQmJefuLtY+p8qT
  FxzHVQpsGS/qT9LoIc7W689MmhB0C/s5Box8F6/4pYImXfyFLvggE8V3wQ==
  =C8Px
  -----END PGP SIGNATURE-----
tags:
  - AI
  - programowanie
  - bezpieczeństwo
  - OWASP
  - vibe coding
title: Od script kiddie do prompt kiddie. Technologia się zmieniła, pajac został ten sam
toc: false
---

Kiedyś przynajmniej trzeba było umieć rozpakować ZIP-a.

Pamiętam czasy, w których Internet był znacznie mniej przyjaznym miejscem, dokumentacja przypominała coś, co pisało się dla ludzi potrafiących czytać dokumentację, a zdobycie wiedzy technicznej wymagało czegoś więcej niż wpisania pytania w okno czatu i poczekania kilkunastu sekund. W tamtym świecie funkcjonował pewien szczególny gatunek użytkownika komputera, którego szybko i całkiem trafnie nazwano *script kiddie*.

Script kiddie nie musiał wiedzieć, dlaczego coś działało. Wystarczyło, że wiedział, jak to uruchomić.

Ściągał gotowy skrypt, exploit albo cały zestaw narzędzi, czytał instrukcję — o ile w ogóle ją czytał — wpisywał adres celu i naciskał Enter. Jeżeli po drugiej stronie coś się przewróciło, nagle wyrastał nam haker. Nie rozumiał protokołu, nie znał podatności, nie wiedział, dlaczego exploit działa ani dlaczego czasem nie działa, ale miał wynik, a wynik w jego własnym przekonaniu stanowił wystarczający dowód kompetencji.

Minęło ćwierć wieku, komputery są nieporównywalnie szybsze, języki programowania dojrzalsze, narzędzia lepsze, dokumentacja łatwiej dostępna niż kiedykolwiek wcześniej, a my dokonaliśmy rzeczy naprawdę niezwykłej.

Wynaleźliśmy script kiddie od nowa.

Tyle że teraz nazywam go **prompt kiddie**.

## Nowa technologia, stary mechanizm

Prompt kiddie nie pobiera już `exploit.pl` z podejrzanego serwera FTP. Nie musi.

Otwiera narzędzie wykorzystujące sztuczną inteligencję i pisze mniej więcej: „zbuduj mi system logowania klasy produkcyjnej, bezpieczny, skalowalny i zgodny z dobrymi praktykami”.

Chwilę później ma kilkaset linii kodu.

Uruchamia.

Działa.

No i, kurwa, jest programistą.

Nie wie wprawdzie, dlaczego wybrano taki sposób przechowywania stanu, nie potrafi wyjaśnić modelu autoryzacji, nie zauważy, że kontrola dostępu została wykonana w niewłaściwej warstwie, nie zastanowi się nad współbieżnością, transakcjami, obsługą błędów, granicami zaufania, zachowaniem systemu podczas częściowej awarii ani nad tym, co właściwie stanie się po wdrożeniu trzeciej instancji aplikacji, ale przecież endpoint odpowiada kodem 200, testy są zielone, a agent napisał na końcu, że implementacja jest „production ready”.

Czego chcieć więcej?

Problem polega na tym, że jest to dokładnie ten sam mechanizm, który obserwowaliśmy dwadzieścia kilka lat temu.

Script kiddie mylił **umiejętność uruchomienia narzędzia z umiejętnością wykonania pracy, którą to narzędzie automatyzowało**.

Prompt kiddie myli **umiejętność opisania oczekiwanego rezultatu z umiejętnością zaprojektowania i stworzenia systemu, który ten rezultat zapewnia**.

Różnica jest jednak zasadnicza, ponieważ stary script kiddie mógł za pomocą gotowego narzędzia zrobić głupią rzecz, natomiast współczesny prompt kiddie może w jedno popołudnie wyprodukować kilkadziesiąt tysięcy linii kodu, których nie rozumie, a następnie, zachęcony tym, że wszystko się kompiluje, wdrożyć je tam, gdzie konsekwencje jego niewiedzy poniosą już inni.

## Kod działa. I właśnie to jest najbardziej niebezpieczne

Największym problemem kodu generowanego przez sztuczną inteligencję nie jest bowiem to, że nie działa.

Gdyby nie działał, sytuacja byłaby wręcz komfortowa.

Program się wywala, człowiek widzi błąd, ktoś zaczyna szukać przyczyny i istnieje przynajmniej szansa, że zanim to gówno trafi na produkcję, ktoś zorientuje się, że coś jest nie tak.

Znacznie ciekawszy jest kod, który działa.

Przyjmuje żądania, zapisuje rekordy, wysyła wiadomości, wystawia interfejs programistyczny, przechodzi testy wygenerowane przez tego samego agenta, który chwilę wcześniej wygenerował implementację, a na demonstracji wygląda znakomicie.

Tylko nikt nie zadał pytania, co stanie się przy dwóch równoczesnych żądaniach, ponieważ prompt kiddie nie wie, że powinien je zadać.

Nikt nie sprawdził granic autoryzacji, bo system przecież „ma JWT”.

Nikt nie zastanowił się nad możliwością powtórzenia operacji, integralnością danych, wyścigami, eskalacją uprawnień, ujawnieniem sekretów, zależnościami, zachowaniem po utracie połączenia czy możliwością obejścia walidacji, ponieważ człowiek odpowiedzialny za kod nie tyle nie zna odpowiedzi, ile przede wszystkim **nie zna pytań**.

I to jest fundamentalna różnica pomiędzy brakiem wiedzy a brakiem doświadczenia.

Człowiek posiadający doświadczenie bardzo często nie zna odpowiedzi, ale zazwyczaj wie, gdzie znajduje się obszar niewiedzy, który należy zbadać. Prompt kiddie dostaje natomiast działający rezultat, więc nawet nie wie, że czegoś nie wie.

To nie jest wyłącznie złośliwa obserwacja. OWASP wprost opisuje już ryzyka związane z kodowaniem wspomaganym przez sztuczną inteligencję: podatne lub zmyślone zależności, niebezpieczne konfiguracje, błędy kontroli dostępu, wycieki kontekstu, pośrednie wstrzykiwanie poleceń do agentów oraz zagrożenia wynikające z tego, że agent może wykonywać polecenia systemowe, instalować pakiety i modyfikować proces budowania oraz wdrażania.

Czyli, mówiąc językiem mniej konferencyjnym: **to, że model napisał kod wyglądający profesjonalnie, nie oznacza jeszcze, że ten kod nie jest gównem**.

## „Ale mam testy”

Oczywiście.

Napisała je ta sama sztuczna inteligencja.

To jeden z moich ulubionych elementów współczesnego teatru jakości.

Agent generuje implementację na podstawie własnego rozumienia polecenia, następnie generuje testy na podstawie własnego rozumienia wygenerowanej przez siebie implementacji, uruchamia je, wszystkie przechodzą, po czym triumfalnie informuje użytkownika, że rozwiązanie jest poprawne.

Prompt kiddie patrzy na `148 passed, 0 failed` i osiąga stan technicznej nirwany.

Problem w tym, że sto procent zaliczonych testów mówi wyłącznie, iż program zachował się zgodnie z tym, czego sprawdzenie przewidziano w testach. Jeżeli ten sam mechanizm stworzył błędne założenie, implementację tego założenia i test potwierdzający to założenie, otrzymaliśmy nie trzy niezależne warstwy kontroli, lecz ten sam błąd ubrany w trzy różne pliki.

OWASP zwraca na to uwagę bardzo jednoznacznie: zaliczenie wszystkich testów nie stanowi dowodu bezpieczeństwa, jeżeli same testy potwierdzają błędne zachowanie, a kod o znaczeniu krytycznym dla bezpieczeństwa nie powinien być bez niezależnej weryfikacji tworzony i testowany przez tego samego agenta.

Ale jest zielono.

Można wdrażać.

## Od kopiowania ze Stack Overflow do produkcji przemysłowej

Nie zamierzam przy tym udawać, że problem narodził się wraz ze sztuczną inteligencją.

Programiści od zawsze kopiowali kod.

Z książek, grup dyskusyjnych, IRC-a, blogów, Stack Overflow, dokumentacji, cudzych projektów i wszystkiego, co akurat podsunęła wyszukiwarka. Samo wykorzystanie kodu stworzonego przez kogoś innego nigdy nie było i nadal nie jest problemem, ponieważ cała informatyka stoi na abstrakcji, ponownym użyciu oraz korzystaniu z pracy wykonanej wcześniej przez innych.

Problem zaczyna się wtedy, kiedy **przestajesz rozumieć to, za co bierzesz odpowiedzialność**.

Kiedyś skopiowanie dwudziestu linii podejrzanego kodu wymagało przynajmniej znalezienia tych dwudziestu linii, dopasowania ich do własnego programu, poprawienia nazw, rozwiązania błędów kompilacji i, przy odrobinie szczęścia, przeczytania komentarza człowieka, który ostrzegał, że rozwiązanie jest paskudnym obejściem problemu.

Dzisiaj prompt kiddie może wygenerować kompletną aplikację wraz z kontenerami, bazą danych, procesem wdrażania, uwierzytelnianiem, testami i konfiguracją chmury, po czym zmienić pół architektury poleceniem:

„Napraw wszystkie problemy i zrób to zgodnie z najlepszymi praktykami”.

Piękne.

Największym osiągnięciem sztucznej inteligencji w programowaniu może się więc okazać nie automatyzacja pisania kodu, lecz **automatyzacja produkcji długu technicznego w skali wcześniej ekonomicznie nieosiągalnej**.

Człowiek produkował gówno z prędkością ograniczoną szybkością pisania na klawiaturze.

Maszyna tego ograniczenia nie ma.

## Vibe coding, czyli nie patrz pod maskę

Samo określenie *vibe coding* powstało jako nazwa sposobu pracy, w którym człowiek opisuje modelowi oczekiwane zachowanie, sprawdza rezultat i prosi o kolejne zmiany, nie zajmując się szczególnie kodem znajdującym się pod spodem. Martin Fowler zwraca uwagę, że takie podejście może mieć sens przy oprogramowaniu jednorazowym lub przeznaczonym dla ograniczonego grona odbiorców, natomiast problemy zaczynają się przy wymaganiach dotyczących poprawności, bezpieczeństwa i utrzymywalności.

I tutaj dochodzimy do rzeczy istotnej: **vibe coding nie jest tym samym co używanie sztucznej inteligencji do programowania**.

Doświadczony programista może używać agenta do generowania kodu i nadal pozostawać programistą.

Może delegować tworzenie powtarzalnych fragmentów, refaktoryzację, analizę repozytorium, przygotowanie testów, dokumentacji czy pierwszej wersji implementacji, ponieważ potrafi następnie przeczytać rezultat, zakwestionować decyzje modelu, znaleźć naruszenie założeń architektonicznych, zauważyć niebezpieczną zależność i powiedzieć: nie, to rozwiązanie jest wprawdzie sprytne, ale za pół roku będziemy przez nie płakać.

Prompt kiddie tego nie potrafi.

Jego przewaga kończy się dokładnie tam, gdzie kończy się zdolność modelu do samodzielnego doprowadzenia zadania do pozornie działającego rezultatu.

Jeżeli agent utknie, prompt kiddie nie zaczyna debugować systemu.

Prompt kiddie zaczyna **debugować rozmowę z agentem**.

„Napraw to”.

„Nadal nie działa”.

„Przeanalizuj dokładnie”.

„Jesteś ekspertem z dwudziestoletnim doświadczeniem”.

„Spróbuj innego podejścia”.

„Teraz naprawdę znajdź przyczynę”.

Po kilkunastu takich iteracjach kod przypomina stanowisko archeologiczne, w którym każda kolejna warstwa powstała po to, żeby zamaskować problemy warstwy poprzedniej, ale użytkownik jest zachwycony, ponieważ po restarcie kontenera formularz wreszcie się zapisał.

## Najbardziej imponująca jest pewność siebie

To właśnie ona łączy script kiddie sprzed ćwierć wieku z prompt kiddie roku 2026.

Nie brak wiedzy.

**Brak świadomości własnego braku wiedzy.**

Script kiddie po udanym użyciu exploita uważał się za hakera.

Prompt kiddie po wygenerowaniu aplikacji uważa się za programistę, architekta, specjalistę od bezpieczeństwa i administratora infrastruktury jednocześnie, ponieważ przecież aplikacja działa, agent zrobił diagram, a w pliku README znajduje się nawet rozdział „Security”.

I podobnie jak kiedyś najbardziej irytujący nie był dzieciak uruchamiający cudzy skrypt, lecz dzieciak przekonany, że właśnie pokonał administratorów NASA, tak dzisiaj nie przeszkadza mi człowiek, który nie potrafi programować i za pomocą sztucznej inteligencji tworzy sobie narzędzie rozwiązujące jego własny problem.

Wręcz przeciwnie.

To jest znakomite zastosowanie tej technologii.

Problem zaczyna się wtedy, kiedy ten człowiek po trzech tygodniach generowania kodu zaczyna tłumaczyć innym, że klasyczne programowanie umarło, architektura nie jest już potrzebna, doświadczenie programistyczne straciło znaczenie, a on właśnie w weekend stworzył konkurencję dla systemu rozwijanego przez kilkudziesięciu inżynierów.

Nie stworzył.

**Wygenerował coś, czego jeszcze nikt kompetentny nie zdążył dokładnie obejrzeć.**

To nie jest to samo.

## Kod przynajmniej nie jest ich

Jest w tym wszystkim jeszcze jedna piękna ironia.

Dawny script kiddie używał cudzego kodu i doskonale o tym wiedział.

Dzisiejszy prompt kiddie używa kodu wygenerowanego przez model, nie rozumie go, nie potrafiłby go samodzielnie odtworzyć, często nie potrafi go nawet sensownie zdebugować, ale mówi o „swojej aplikacji”, „swojej architekturze” i „swoim rozwiązaniu” z pewnością człowieka, który właśnie wrócił z piwnicy po samodzielnym skonstruowaniu procesora.

A kiedy coś pierdolnie?

„AI zrobiło błąd”.

Nie.

Ty zrobiłeś błąd.

Sztuczna inteligencja wygenerowała tekst.

Ty zdecydowałeś, że ten tekst stanie się programem.

Ty zaakceptowałeś zmianę.

Ty ją zatwierdziłeś.

Ty ją wdrożyłeś.

Jeżeli nie rozumiesz kodu wystarczająco dobrze, żeby wziąć za niego odpowiedzialność, problemem nie jest to, że napisała go sztuczna inteligencja. Problemem jest to, że **dopuściłeś do produkcji kod, którego nie jesteś w stanie ocenić**.

OWASP formułuje tę zasadę bez miejsca na wygodne wymówki: kod wygenerowany przez sztuczną inteligencję powinien mieć człowieka odpowiedzialnego za jego poprawność, bezpieczeństwo i utrzymanie, a stwierdzenie „AI to napisała” tej odpowiedzialności nie przenosi na model.

## Prompt kiddie nie jest problemem sztucznej inteligencji

I tu wypada powiedzieć coś, co może zepsuć cały ten piękny festiwal narzekania.

Sztuczna inteligencja w programowaniu jest cholernie dobrym narzędziem.

Prawdopodobnie jednym z najlepszych, jakie dostaliśmy od czasu upowszechnienia Internetu, systemów kontroli wersji i nowoczesnych środowisk programistycznych.

Pozwala szybciej analizować obcy kod, sprawdzać hipotezy, tworzyć prototypy, wykonywać nudne przekształcenia, przygotowywać testy, szukać błędów, porównywać rozwiązania i delegować ogromną ilość mechanicznej roboty, która wcześniej pochłaniała czas mogący zostać wykorzystany na myślenie.

Tyle że właśnie tutaj znajduje się słowo kluczowe.

**Myślenie.**

Sztuczna inteligencja może drastycznie ograniczyć ilość kodu, który programista musi własnoręcznie napisać.

Nie ogranicza proporcjonalnie ilości kodu, który programista powinien **rozumieć**.

Może wręcz powodować, że trzeba rozumieć więcej, ponieważ jeden człowiek jest obecnie w stanie wygenerować w ciągu dnia tyle zmian, ile kiedyś pisał przez tydzień, a każda z tych zmian może zawierać decyzję, której konsekwencje ujawnią się dopiero pod obciążeniem, podczas awarii albo wtedy, gdy ktoś zacznie aktywnie szukać sposobu na obejście zabezpieczeń.

Dlatego nie boję się, że sztuczna inteligencja zastąpi dobrych programistów.

Znacznie bardziej interesuje mnie to, że pozwala słabym programistom produkować kod z prędkością dobrych.

Bo prędkość generowania kodu właśnie przestała być problemem.

Prędkość jego **rozumienia** nadal nim jest.

## Historia zatoczyła koło

Dwadzieścia kilka lat temu mieliśmy dzieciaka z paczką exploitów.

Nie wiedział, jak działa przepełnienie bufora, ale miał program, który je wykorzystywał.

Nie znał protokołu, ale miał skaner.

Nie rozumiał ataku, ale miał przycisk.

Nazywaliśmy go script kiddie albo pack kiddie i nikt poważny nie mylił obsługi narzędzia ze znajomością dziedziny.

Dzisiaj mamy człowieka z agentem sztucznej inteligencji.

Nie rozumie modelu współbieżności, ale ma mikroserwisy.

Nie zna kryptografii, ale ma `EncryptionService`.

Nie rozumie uwierzytelniania, ale ma OAuth.

Nie zna Kubernetes, ale ma manifesty.

Nie zna architektury, ale ma diagram w Mermaidzie.

Nie potrafi znaleźć błędu, ale może napisać:

„Znajdź i napraw wszystkie błędy”.

I tak oto po ćwierćwieczu gigantycznego postępu technologicznego wróciliśmy dokładnie do punktu wyjścia.

Zmieniły się narzędzia.

Zmieniła się skala.

Zmieniła się prędkość.

Pajac został ten sam.

Kiedyś był **script kiddie**.

Dzisiaj jest **prompt kiddie**.

Tylko teraz potrafi wygenerować całe repozytorium, zanim ktokolwiek zdąży mu powiedzieć, że nie ma pojęcia, co robi.

A kod?

Kod może być gówniany, pełen błędów i niebezpieczny.

Ale przynajmniej nie jest jego.
