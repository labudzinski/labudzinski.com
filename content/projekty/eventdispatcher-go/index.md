+++
title = "Event Dispatcher for Go"
description = "Biblioteka Go do wzorca obserwatora: listenery z priorytetem, dispatch zdarzeń po nazwie, subskrybenci, pooling GenericEvent i proxy tylko do odczytu."
date = 2026-07-08T12:00:00+02:00
draft = false
docs_url = "https://docs.lemric.com/eventdispatcher-go/"
go_module = "github.com/lemric/eventdispatcher-go"
tags = ["Go", "events", "concurrency"]
toc = false
+++

Event Dispatcher for Go to lekka biblioteka do obsługi zdarzeń w stylu wzorca obserwatora. Każde zdarzenie ma unikalną nazwę, na przykład `user.created` lub `order.placed`, a zarejestrowane listenery otrzymują instancję implementującą interfejs `Event`. Proste zdarzenia mogą być puste, a zdarzenia z możliwością zatrzymania propagacji osadzają `BaseEvent` i wywołują `StopPropagation()`, aby pominąć kolejne handlery.

`EventDispatcher` jest wątkowo bezpieczny. Wewnętrznie korzysta z `sync.Map` oraz blokad na poziomie pojedynczego zdarzenia, a podczas dispatchu robi snapshot listy handlerów i zwalnia blokadę przed ich wywołaniem. Paniki w listenerach są przechwytywane, żeby jeden błąd nie wyłączał całego mechanizmu.

Listener rejestrujesz przez `AddListener` z opcjonalnym priorytetem. Wyższa wartość oznacza wcześniejsze wykonanie. Funkcja zwraca closure do wypisania konkretnego handlera po identyfikatorze. Możesz też usunąć listener wskaźnikiem funkcji przez `RemoveListener`.

Własne zdarzenie definiujesz jako strukturę z osadzonym `BaseEvent` i metodą `EventName()`. Dispatch przyjmuje instancję i opcjonalnie nazwę. Gdy przekażesz pusty string, dispatcher odczyta nazwę z metody zdarzenia.

Dla dynamicznych ładunków służy `GenericEvent` z pulą obiektów. `AcquireGenericEvent` i `ReleaseGenericEvent` ograniczają presję na GC przy dużej liczbie krótkich zdarzeń. Argumenty obsługujesz jak mapę: `GetArgument`, `SetArgument`, `OffsetGet` i pozostałe metody z API dokumentacji.

Wzorzec subskrybenta grupuje wiele handlerów w jednej strukturze. `SubscribedEvents` zwraca listę par nazwa, funkcja i priorytet. `AddSubscriber` rejestruje całość naraz, a `RemoveSubscriber` wypisuje wszystkie powiązane listenery.

`ImmutableEventDispatcher` to proxy tylko do odczytu. Dispatch działa normalnie, ale mutatory takie jak `AddListener` kończą się paniką. Przydaje się, gdy przekazujesz dispatcher dalej bez prawa modyfikacji rejestru.

Pełne API, przykłady i testy opisuje dokumentacja na docs.lemric.com. Instalacja modułu: `go get github.com/lemric/eventdispatcher-go`.
